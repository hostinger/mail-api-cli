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

// Request returns the API client for a command run. Without a static token it
// signs in through OAuth, opening the browser when needed.
func Request() client.ClientWithResponsesInterface {
	return newClient(true)
}

// CompletionRequest returns an API client for shell completion: it uses the
// static token or an already stored OAuth sign-in, and never starts an
// interactive login.
func CompletionRequest() client.ClientWithResponsesInterface {
	return newClient(false)
}

// BaseURL returns the configured api_url, or the API's default base URL.
func BaseURL() string {
	if apiURL := viper.GetString("api_url"); apiURL != "" {
		return apiURL
	}
	return "https://api.mail.hostinger.com"
}

// UserAgent is the User-Agent header value of every API request.
func UserAgent() string {
	return fmt.Sprintf("hostinger_mail-cli/%s", utils.CLIVersion.String(false))
}

func newClient(interactive bool) client.ClientWithResponsesInterface {
	apiURL, opts := clientOptions(interactive, interactive && viper.GetBool("dry_run"))
	c, err := client.NewClientWithResponses(apiURL, opts...)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

// clientOptions are shared by the typed client and Raw. dryRun prints each
// request instead of sending it.
func clientOptions(interactive, dryRun bool) (string, []client.ClientOption) {
	apiURL := BaseURL()

	opts := []client.ClientOption{
		client.WithRequestEditorFn(addUserAgent),
	}

	var doer client.HttpRequestDoer = &http.Client{}
	if viper.GetBool("verbose") {
		doer = &verboseDoer{base: doer}
	}

	if dryRun {
		// Print the request instead of sending it; the request itself needs no
		// credentials (lookups through Raw may, see Raw).
		opts = append(opts, client.WithHTTPClient(&dryRunDoer{}))
	} else if staticToken := viper.GetString("api_token"); staticToken != "" {
		// Static token wins: use it unchanged. OAuth and 401 re-auth are skipped
		// because there is no refresh source for a user-managed token.
		bearer, err := securityprovider.NewSecurityProviderBearerToken(staticToken)
		if err != nil {
			panic(err)
		}
		opts = append(opts, client.WithRequestEditorFn(bearer.Intercept), client.WithHTTPClient(doer))
	} else {
		opts = append(opts, client.WithHTTPClient(noTokenDoer{}))
	}

	return apiURL, opts
}

func addUserAgent(ctx context.Context, req *http.Request) error {
	req.Header.Set("User-Agent", UserAgent())
	return nil
}
