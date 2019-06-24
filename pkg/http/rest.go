package http

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	log "github.com/golang/glog"
)

const (
	defaultTimeout = 3 * time.Second
)

// RestHelper provides an easy interface to send requests to a REST endpoint
type RestHelper struct {
	baseURL             string
	authorizationHeader string
	timeout             time.Duration
}

// NewRestHelper returns an instance of RestHelper
func NewRestHelper(baseURL string, authorizationHeader string) *RestHelper {
	return &RestHelper{
		baseURL:             baseURL,
		authorizationHeader: authorizationHeader,
		timeout:             defaultTimeout,
	}
}

func NewRestHelperWithTimeout(baseURL string, authorizationHeader string,
	timeout time.Duration) *RestHelper {
	return &RestHelper{
		baseURL:             baseURL,
		authorizationHeader: authorizationHeader,
		timeout:             timeout,
	}
}

// SendRequestWithRetry generates an HTTP request to the REST endpoint with retries
func (h *RestHelper) SendRequestWithRetry(endpointName string, method string, params *url.Values,
	payload interface{}, maxAttempts int, baseWaitMs int) ([]byte, error) {
	url := fmt.Sprintf("%v/%v", h.baseURL, endpointName)
	return h.SendRequestToURLWithRetry(url, method, params, payload, maxAttempts, baseWaitMs)
}

// SendRequestToURLWithRetry generates a HTTP request to a URL with retries
func (h *RestHelper) SendRequestToURLWithRetry(url string, method string, params *url.Values,
	payload interface{}, maxAttempts int, baseWaitMs int) ([]byte, error) {

	attempt := 1
	var err error
	var bys []byte

	for {
		bys, err = h.SendRequestToURL(url, method, params, payload)
		if err != nil {
			log.Infof(
				"err with request, sleep/attempt again, waiting %v ms...",
				baseWaitMs*attempt,
			)

			// Take a break and retry
			time.Sleep(time.Duration(baseWaitMs) * time.Duration(attempt) * time.Millisecond)
			if attempt >= maxAttempts {
				return nil, err
			}

			attempt++
			continue
		}
		break
	}
	return bys, nil
}

// SendRequest generates an HTTP request to the REST endpoint
func (h *RestHelper) SendRequest(endpointName string, method string, params *url.Values,
	payload interface{}) ([]byte, error) {
	url := fmt.Sprintf("%v/%v", h.baseURL, endpointName)
	return h.SendRequestToURL(url, method, params, payload)
}

// SendRequestToURL generates an HTTP request to the given URL.
func (h *RestHelper) SendRequestToURL(url string, method string, params *url.Values,
	payload interface{}) ([]byte, error) {

	client := &http.Client{}
	// Make sure a timeout is set or this will hang forever when making
	// request that fails to return
	if h.timeout == 0 {
		h.timeout = defaultTimeout
	}
	client.Timeout = h.timeout

	var req *http.Request
	var err error

	if method == http.MethodPost {
		req, err = h.buildPostPutRequest(method, url, payload)
	} else {
		req, err = h.buildGetDeleteRequest(method, url, params)
	}

	if err != nil {
		return nil, err
	}

	// Add the authorization header if given
	if h.authorizationHeader != "" {
		req.Header.Add("Authorization", h.authorizationHeader)
	}

	// Make the request
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close() // nolint: errcheck
	rspBodyData, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode != 200 && rsp.StatusCode != 201 {
		return nil, fmt.Errorf("Request failed: %v, %v", rsp.StatusCode, string(rspBodyData))
	}

	return rspBodyData, nil
}

func (h *RestHelper) buildPostPutRequest(method string, url string,
	payload interface{}) (*http.Request, error) {
	var reqBody *bytes.Buffer
	// If there was a payload struct to marshal into payload string
	if payload != nil {
		payloadData, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBufferString(string(payloadData))

	} else {
		reqBody = bytes.NewBufferString("")
	}

	// Build a new request
	req, err := http.NewRequest(
		method,
		url,
		reqBody,
	)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (h *RestHelper) buildGetDeleteRequest(method string, url string,
	params *url.Values) (*http.Request, error) {
	req, err := http.NewRequest(
		method,
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}
	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	return req, nil
}

// BuildBasicAuthHeader builds an auth header needed for Basic Authentication
func BuildBasicAuthHeader(username string, password string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", username, password)))

	return "Basic " + encoded
}
