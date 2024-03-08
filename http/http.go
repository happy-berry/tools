package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Param struct {
	Url      string
	Body     interface{}
	Auth     string
	TimeOut  int
	Header   map[string]string
	Callback func(b []byte) error
}

func Post(p Param) error {
	jsonStr, _ := json.Marshal(p.Body)
	fmt.Printf("[Http Util] post : %s\nbody: %s\nauth:%s\n ", p.Url, jsonStr, p.Auth)
	req, err := http.NewRequest("POST", p.Url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("[Http Util] post new request err : ", err)
		return err
	}
	for k, v := range p.Header {
		req.Header.Add(k, v)
	}
	if p.Auth != "" {
		req.Header.Add("Authorization", p.Auth)
	}
	defer req.Body.Close()
	client := &http.Client{Timeout: time.Duration(p.TimeOut) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[Http Util] post client err : ", err)
		return err
	}
	result, _ := io.ReadAll(resp.Body)
	fmt.Println("[Http Util] post resp: ", string(result))
	if resp.StatusCode != 200 {
		return errors.New(string(result))
	}
	defer resp.Body.Close()
	return p.Callback(result)
}
