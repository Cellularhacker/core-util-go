package core

import (
	"crypto/tls"
	"errors"
	"net/http"
	"time"
)

var httpClient *HttpClient

// HttpClient : The struct for using HTTP(S) Requests.
type HttpClient struct {
	client     *http.Client
	maxRetries int
	retryDelay time.Duration
}

// HttpErrors : that could be happened during~end of the HTTP(S) Requests in Unexpected errors or Custom errors.
var HttpErrors = []error{
	errors.New("out of retries to try fetch HTTP(S) request"),
}

func initHttpClient() {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 1000
	t.MaxIdleConnsPerHost = 1000
	t.IdleConnTimeout = 1200 * time.Second
	t.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	httpClient = &HttpClient{
		client: &http.Client{
			Timeout:   1200 * time.Second,
			Transport: t,
		},
		maxRetries: 5,
		retryDelay: 1 * time.Second,
	}
}
func initResponse() (*http.Response, error) {
	return nil, HttpErrors[0]
}
func GetHttpClient() *HttpClient {
	return httpClient
}
func (hc *HttpClient) SetMaxRetries(maxRetries int) *HttpClient {
	hc.maxRetries = maxRetries

	return hc
}
func (hc *HttpClient) SetRetryDelay(retryDelay time.Duration) *HttpClient {
	hc.retryDelay = retryDelay

	return hc
}
func (hc *HttpClient) GetMaxRetriesCount() int {
	return hc.maxRetries
}
func (hc *HttpClient) GetRetryDelay() time.Duration {
	return hc.retryDelay
}
func (hc *HttpClient) Do(req *http.Request) (*http.Response, error) {
	// MARK: Initialize return values.
	resp, err := initResponse()

	// MARK: Start n count of HTTP(S) Request(s)...
	for i := 0; i < hc.maxRetries; i++ {
		resp, err = hc.client.Do(req)
		if err == nil {
			return resp, nil
		}

		<-time.NewTimer(hc.retryDelay).C
	}

	return resp, err
}

/*
 * Utilities : Setting up User-Agent related headers
 */

// SetUserAgentMacOS : Based on Macbook Pro Chrome
func (hc *HttpClient) SetUserAgentMacOS(req *http.Request) *http.Request {
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 Edg/128.0.0.0")
	req.Header.Set("sec-ch-ua", `"Chromium";v="128", "Not;A=Brand";v="24", "Microsoft Edge";v="128"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?0")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Dnt", "1")
	req.Header.Set("Priority", "u=0, i")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Accept-Language", "ko,en;q=0.9,en-US;q=0.8,ja;q=0.7")
	req.Header.Set("Accept-Encoding", "deflate, br, zstd")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")

	return req
}
func (hc *HttpClient) SetUserAgentiPhone(req *http.Request) *http.Request {
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Mobile/15E148 Safari/604.1 Edg/128.0.0.0")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Priority", "u=0, i")
	req.Header.Set("Dnt", "1")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Accept-Language", "ko,en;q=0.9,en-US;q=0.8,ja;q=0.7")
	req.Header.Set("Accept-Encoding", "deflate, br, zstd")

	return req
}
func (hc *HttpClient) SetUserAgentAndroid(req *http.Request) *http.Request {
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Mobile Safari/537.36 Edg/128.0.0.0")
	req.Header.Set("sec-ch-ua", `"Chromium";v="128", "Not;A=Brand";v="24", "Microsoft Edge";v="128"`)
	req.Header.Set("sec-ch-ua-mobile", "?1")
	req.Header.Set("sec-ch-ua-platform", "Android")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Dnt", "1")
	req.Header.Set("Priority", "u=0, i")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Accept-Language", "ko,en;q=0.9,en-US;q=0.8,ja;q=0.7")
	req.Header.Set("Accept-Encoding", "deflate, br, zstd")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")

	return req
}
