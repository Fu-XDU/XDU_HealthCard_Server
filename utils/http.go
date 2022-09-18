package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	TcpFailError    = "tcp连接失败"
	ApplicationJson = "application/json"
)

func Post(data interface{}, url string) (body []byte, err error) {
	jsonStr, _ := json.Marshal(data)
	resp, err := http.Post(url,
		ApplicationJson,
		bytes.NewBuffer(jsonStr))
	if err != nil {
		return
	}
	if resp == nil {
		return nil, errors.New(TcpFailError)
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	return
}

func Get(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp == nil {
		return nil, errors.New(TcpFailError)
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	return
}
