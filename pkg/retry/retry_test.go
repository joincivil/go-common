package retry_test

import (
	"fmt"
	"testing"

	"github.com/joincivil/go-common/pkg/retry"
)

func TestRetryWrapperString(t *testing.T) {
	testFunc := func() (interface{}, error) {
		return "test", nil
	}

	result, err := retry.Wrapper(testFunc, 5, 300)
	if err != nil {
		t.Fatalf("Should not have received an error: err: %v", err)
	}

	resultStr, ok := result.(string)
	if !ok {
		t.Fatalf("Should have correctly inferred as a string")
	}
	if resultStr != "test" {
		t.Fatalf("Should have been the correct string value")
	}
}

func TestRetryWrapperStringMaxAttempts(t *testing.T) {
	testFunc := func() (interface{}, error) {
		return nil, fmt.Errorf("Error")
	}

	result, err := retry.Wrapper(testFunc, 5, 300)
	if err == nil {
		t.Fatalf("Should have received an error: err: %v", err)
	}

	if result != nil {
		t.Fatalf("Should have returned a nil value")
	}
}

func TestRetryWrapperStringSomeAttempts(t *testing.T) {
	attempt := 0
	testFunc := func() (interface{}, error) {
		if attempt == 2 {
			return 1000, nil
		}
		attempt++
		return nil, fmt.Errorf("Error")
	}

	result, err := retry.Wrapper(testFunc, 5, 300)
	if err != nil {
		t.Fatalf("Should not have received an error: err: %v", err)
	}
	resultInt, ok := result.(int)
	if !ok {
		t.Fatalf("Should have correctly inferred as an int")
	}
	if resultInt != 1000 {
		t.Fatalf("Should have been the correct int value")
	}
}
