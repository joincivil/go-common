package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	chttp "github.com/joincivil/go-common/pkg/http"
)

func TestSendRequestWithRetry(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK")) // nolint: errcheck
	}))
	defer ts.Close()

	rh := chttp.NewRestHelper(ts.URL, "")
	bys, err := rh.SendRequestWithRetry("/", http.MethodGet, nil, nil, 3, 1000)
	if err != nil {
		t.Errorf("Should not have received error: err: %v", err)
	}

	if string(bys) != "OK" {
		t.Errorf("Should have received an OK")
	}
}

func TestSendRequestWithRetryErrors(t *testing.T) {
	count := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR")) // nolint: errcheck
		count++
	}))
	defer ts.Close()

	rh := chttp.NewRestHelper(ts.URL, "")
	_, err := rh.SendRequestWithRetry("/", http.MethodGet, nil, nil, 3, 100)
	if err == nil {
		t.Errorf("Should have received error")
	}

	if count != 3 {
		t.Errorf("Should have seen 3 attempts: %v", count)
	}
}

func TestSendRequestWithTimeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK")) // nolint: errcheck
	}))
	defer ts.Close()

	rh := chttp.NewRestHelperWithTimeout(ts.URL, "", 3*time.Second)
	_, err := rh.SendRequest("/", http.MethodGet, nil, nil)
	if err != nil {
		t.Errorf("Should not have received timeout or error")
	}
}

func TestSendRequestWithTimeoutTimedOut(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		w.Write([]byte("OK")) // nolint: errcheck
	}))
	defer ts.Close()

	rh := chttp.NewRestHelperWithTimeout(ts.URL, "", 3*time.Second)
	_, err := rh.SendRequest("/", http.MethodGet, nil, nil)
	if err == nil {
		t.Errorf("Should have received timeout")
	}
}

func TestSendPostRequestToURL(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK")) // nolint: errcheck
	}))
	defer ts.Close()

	rh := chttp.NewRestHelper(ts.URL, "")
	bys, err := rh.SendRequest("/", http.MethodPost, nil, nil)
	if err != nil {
		t.Errorf("Should not have received error: err: %v", err)
	}

	if string(bys) != "OK" {
		t.Errorf("Should have received an OK")
	}

	bys, err = rh.SendRequest("/", http.MethodPost, nil, "testpayload")
	if err != nil {
		t.Errorf("Should not have received error: err: %v", err)
	}

	if string(bys) != "OK" {
		t.Errorf("Should have received an OK")
	}
}
