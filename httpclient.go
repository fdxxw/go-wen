package wen

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

var HttpClient *http.Client

func HttpGet(url string, params map[string]string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if params != nil {
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	return HttpClient.Do(req)
}

func HttpPost(url string, body io.Reader, params map[string]string, headers map[string]string) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if params != nil {
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	return HttpClient.Do(req)
}

func HttpPostJson(url string, body interface{}, params map[string]string, headers map[string]string) (*http.Response, error) {

	var data []byte

	if body != nil {
		var err error
		data, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")

	q := req.URL.Query()
	if params != nil {
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	return HttpClient.Do(req)
}
