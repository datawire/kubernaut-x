package util

import (
	"crypto/tls"
	"net/http"
	"time"
)

// NewHTTPClient returns an HTTP client configured with the passed options.
func NewHTTPClient(timeout time.Duration, skipVerify bool) *http.Client {
	client := &http.Client{
		Timeout: timeout,
	}

	if !skipVerify {
		// Disable HTTPS certificate validation for this client because we are likely using self-signed certificates
		// during tests.
		client.Transport = &http.Transport{
			/* #nosec */
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	return client
}
