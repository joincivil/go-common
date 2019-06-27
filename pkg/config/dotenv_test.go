package config_test

import (
	"os"
	"testing"

	"github.com/joincivil/go-common/pkg/config"
)

func TestPopulateFromDotEnv(t *testing.T) {
	err := os.Setenv("TESTING_ENV", "testing")
	if err != nil {
		t.Fatalf("Should have succeeded in setting env env var: %v", err)
	}

	err = config.PopulateFromDotEnv("TESTING_ENV")
	if err != nil {
		t.Errorf("Should not have gotten error %v", err)
	}
	testVar := os.Getenv("THIS_IS_A_TEST_VAR")
	if testVar == "" {
		t.Errorf("Should have put the test into the environment")
	}

	if testVar != "hello world" {
		t.Errorf("Should have put the correct env var value in environment")
	}
}

func TestPopulateFromDotEnvBad(t *testing.T) {
	err := os.Setenv("TESTING_ENV", "bad")
	if err != nil {
		t.Fatalf("Should have succeeded in setting env env var: %v", err)
	}

	err = config.PopulateFromDotEnv("TESTING_ENV")
	if err == nil {
		t.Errorf("Should have gotten error")
	}
}
