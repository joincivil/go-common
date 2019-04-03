// +build integration

package errors_test

import (
	"errors"
	"net/http"
	"os"
	"testing"
	"time"

	reporting "github.com/joincivil/go-common/pkg/errors"
)

const (
	testSentryDsnEnvVarName   = "SENTRY_DSN"
	testGoogleCredsEnvVarName = "GOOGLE_APPLICATION_CREDENTIALS"
)

func TestMetaErrorReportingErr(t *testing.T) {
	config := &reporting.MetaErrorReporterConfig{
		StackDriverProjectID:      "civil-media",
		StackDriverServiceName:    "test-service",
		StackDriverServiceVersion: "0.1",
	}
	sentryDSN := os.Getenv(testSentryDsnEnvVarName)
	googleCreds := os.Getenv(testGoogleCredsEnvVarName)
	if googleCreds == "" && sentryDSN == "" {
		t.Logf("No creds set, skipping error reporting tests")
		return
	}
	if sentryDSN != "" {
		config.SentryDSN = sentryDSN
		config.SentryDebug = true
		config.SentryEnv = "test"
		config.SentryLoggerName = "test_logger"
		config.SentryRelease = "0.1"
		config.SentrySampleRate = 1.0
	}
	reporter, rerr := reporting.NewMetaErrorReporter(config)
	if rerr != nil {
		t.Errorf("Error creating meta reporting: %v", rerr)
	}
	err := errors.New("Testing error, ignore")
	reporter.Error(err, nil)

	// Ensure all the events get flushed by the libs
	time.Sleep(3 * time.Second)
}

func TestMetaErrorReportingErrWithMeta(t *testing.T) {
	config := &reporting.MetaErrorReporterConfig{
		StackDriverProjectID:      "civil-media",
		StackDriverServiceName:    "test-service",
		StackDriverServiceVersion: "0.1",
	}
	sentryDSN := os.Getenv(testSentryDsnEnvVarName)
	googleCreds := os.Getenv(testGoogleCredsEnvVarName)
	if googleCreds == "" && sentryDSN == "" {
		t.Logf("No creds set, skipping error reporting tests")
		return
	}
	if sentryDSN != "" {
		config.SentryDSN = sentryDSN
		config.SentryDebug = true
		config.SentryEnv = "test"
		config.SentryLoggerName = "test_logger"
		config.SentryRelease = "0.1"
		config.SentrySampleRate = 1.0
	}
	reporter, rerr := reporting.NewMetaErrorReporter(config)
	if rerr != nil {
		t.Errorf("Error creating meta reporting: %v", rerr)
	}
	err := errors.New("Testing error, ignore")

	userID := "tobiasfunke"
	testReq, _ := http.NewRequest("GET", "http://civil.co", nil)
	meta := &reporting.ErrorMeta{
		UserID: &userID,
		Tags: map[string]string{
			"testkey": "testval",
		},
		RelatedRequest: testReq,
	}
	reporter.Error(err, meta)

	// Ensure all the events get flushed by the libs
	time.Sleep(3 * time.Second)
}
