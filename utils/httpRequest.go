package utils

import (
	"bytes"
	"net/http"
)

//SendHTTPRequest ...
func SendHTTPRequest(header http.Header, body []byte, url, method string) (*http.Response, error) {

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header = header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp, nil
}
