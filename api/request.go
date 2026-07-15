package api

import (
	"context"
	"fmt"
	"github.com/hostinger/mail-api-cli/client"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func Request() client.ClientWithResponsesInterface {
	apiURL := viper.GetString("api_url")
	if apiURL == "" {
		apiURL = "https://api.mail.hostinger.com"
	}

	opts := []client.ClientOption{
		client.WithRequestEditorFn(addUserAgent),
	}

	if staticToken := viper.GetString("api_token"); staticToken != "" {
		// Static token wins: use it unchanged. OAuth and 401 re-auth are skipped
		// because there is no refresh source for a user-managed token.
		bearer, err := securityprovider.NewSecurityProviderBearerToken(staticToken)
		if err != nil {
			panic(err)
		}
		opts = append(opts, client.WithRequestEditorFn(bearer.Intercept))
	} else {
		log.Fatal("no API token configured")
	}

	c, err := client.NewClientWithResponses(apiURL, opts...)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func addUserAgent(ctx context.Context, req *http.Request) error {
	req.Header.Set("User-Agent", fmt.Sprintf("hostinger_mail-cli/%s", utils.CLIVersion.String(false)))
	return nil
}
