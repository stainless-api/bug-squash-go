package testutil

import (
	"net/http"
	"os"
	"strconv"
	"testing"
)

func CheckTestServer(t *testing.T, url string) bool {
	if _, err := http.Get(url); err != nil {
		str := os.Getenv("SKIP_PRISM_TESTS")
		skip, err := strconv.ParseBool(str)
		if err != nil && str != "" {
			panic("invalid value for flag SKIP_PRISM_TESTS \"" + str + "\", should be parsable as a bool")
		}
		if skip {
			t.Skip("Skipping test that requires a mock Prism server due to SKIP_PRISM_TESTS=true")
			return false
		}
		panic("The test will not run without a mock Prism server running against your OpenAPI spec: see README.md > Development. You can set the environment variable SKIP_PRISM_TESTS to true to skip running any tests that require the Prism server.")
	}
	return true
}

func CheckIntegrationServer(t *testing.T, url string) bool {
	if _, err := http.Get(url); err != nil {
		str := os.Getenv("SKIP_INTEGRATION_TESTS")
		skip, err := strconv.ParseBool(str)
		if err != nil {
			panic("invalid value for flag SKIP_INTEGRATION_TESTS \"" + str + "\", should be parsable as a bool")
		}
		if skip {
			t.Skip("Skipping test that requires a mock integration server due to SKIP_INTEGRATION_TESTS=true")
			return false
		}
		t.Error("The test will not run without a mock integration server running locally: see README.md > Development. You can set the environment variable SKIP_INTEGRATION_TESTS to true to skip running any tests that require the integration server.")
		return false
	}
	return true
}
