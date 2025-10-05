package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func sendMessage(cfg *Config, message []byte) (*http.Response, error) {
	payload := []byte(`<?xml version="1.0" encoding="utf-8" ?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:m="http://schemas.microsoft.com/exchange/services/2006/messages" xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Header>
    <t:RequestServerVersion Version="Exchange2007_SP1" />
  </soap:Header>
  <soap:Body>
`)
	payload = append(payload, message...)
	payload = append(payload, "\n  </soap:Body>\n</soap:Envelope>"...)

	req, err := http.NewRequest("POST", cfg.URL, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	req.SetBasicAuth(
		fmt.Sprintf("%s\\%s", cfg.Domain, cfg.Username),
		cfg.Password,
	)
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("Accept", "text/xml")
	req.Header.Set("User-Agent", DefaultUserAgent)

	client := new(http.Client)
	client.Timeout = cfg.Timeout
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return client.Do(req)
}
