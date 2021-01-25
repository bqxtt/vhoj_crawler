package http_utils

import (
	"github.com/bqxtt/vhoj_crawler/pkg/factory"
	"io/ioutil"
	"net/http"
)

func Download(method string, url string) (string, error) {
	httpClient := factory.NewHttpClient()
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}
	resp, err := httpClient.Client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	return string(body), nil
}
