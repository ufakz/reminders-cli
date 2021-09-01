package client

import (
	"net/http"
	"time"
)

type HTTPClient struct {
	client     *http.Client
	BackendURI string
}

func NewHTTPClient(uri string) HTTPClient {
	return HTTPClient{
		BackendURI: uri,
		client:     &http.Client{},
	}
}

func (c HTTPClient) Create(title, message string, duration time.Duration) ([]byte, error) {
	res := []byte(`response for create reminder`)
	return res, nil
}

func (c HTTPClient) Edit(id, title, message string, duration time.Duration) ([]byte, error) {
	res := []byte(`response for edit reminder`)
	return res, nil
}

func (c HTTPClient) Fetch(ids []string) ([]byte, error) {
	res := []byte(`response for fetch reminder`)
	return res, nil
}

func (c HTTPClient) Delete(ids []string) error {
	return nil
}

func (c HTTPClient) Healthy(host string) bool {
	return true
}
