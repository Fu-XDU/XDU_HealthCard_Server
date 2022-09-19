package utils

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	TcpFailError    = "tcp连接失败"
	ApplicationJson = "application/json"
)

func Post(data []byte, url string, header map[string]string) (resp *http.Response, body []byte, err error) {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", ApplicationJson)
	for key := range header {
		req.Header.Set(key, header[key])
	}
	resp, err = (&http.Client{}).Do(req)
	if err != nil {
		return
	}
	if resp == nil {
		return nil, nil, errors.New(TcpFailError)
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	return
}

func Get(url string) (resp *http.Response, body []byte, err error) {
	resp, err = http.Get(url)
	if err != nil {
		return
	}
	if resp == nil {
		return nil, nil, errors.New(TcpFailError)
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	return
}
