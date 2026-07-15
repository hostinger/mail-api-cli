package utils

import (
	"strconv"
	"time"
)

func StringPtrOrNil(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func IntPtrOrNil(i int) *int {
	if i == -1 {
		return nil
	}
	return &i
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func StringArrayToIntArray(strArr []string) []int {
	intArr := make([]int, 0, len(strArr))

	for _, str := range strArr {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		intArr = append(intArr, num)
	}

	return intArr
}

func StringToTime(dateStr string) time.Time {
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		panic(err)
	}

	return t
}
