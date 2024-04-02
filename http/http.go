package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Method string

const (
	Get  Method = "GET"
	Post Method = "POST"
)

type Param struct {
	Method   Method
	Url      string
	Body     interface{}
	Auth     string
	TimeOut  int
	Header   map[string]string
	Callback func(b []byte) error
}

func Request(p Param) error {
	jsonStr, err := json.Marshal(p.Body)
	if err != nil {
		fmt.Println("[Http Util] Marshal body err: ", err)
		return err
	}

	fmt.Printf("[Http Util] method: %s, url: %s, body: %s, auth: %s\n", p.Method, p.Url, jsonStr, p.Auth)

	req, err := http.NewRequest(string(p.Method), p.Url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("[Http Util] New request err: ", err)
		return err
	}

	for k, v := range p.Header {
		req.Header.Add(k, v)
	}

	if p.Auth != "" {
		req.Header.Add("Authorization", p.Auth)
	}

	client := &http.Client{
		Timeout: time.Duration(p.TimeOut) * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[Http Util] Client err: ", err)
		return err
	}

	defer resp.Body.Close()
	result, _ := io.ReadAll(resp.Body)
	fmt.Println("[Http Util] post resp: ", string(result))
	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	return p.Callback(result)
}
