package factory

import (
	"net/http"
	"sync"
)

type httpClient struct {
	Client *http.Client
}

var client *httpClient
var once sync.Once

func NewHttpClient() *httpClient {
	once.Do(func() {
		client = &httpClient{Client: &http.Client{}}
	})
	return client
}
