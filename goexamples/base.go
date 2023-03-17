package goexamples

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type payload struct {
	Source    string         `json:"source"`
	Url       string         `json:"url,omitempty"`
	Query     string         `json:"query,omitempty"`
	Domain    string         `json:"domain,omitempty"`
	StartPage int            `json:"start_page,omitempty"`
	Context   []contextEntry `json:"context,omitempty"`
	Parse     bool           `json:"parse"`
}

type contextEntry struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

const (
	realtimeAPIURL string = "https://realtime.oxylabs.io/v1/queries"
	username       string = "YOUR_USERNAME"
	password       string = "YOUR_PASSWORD"
)

func createRequestWithAuthAndPayload(p payload) (*http.Request, error) {
	payloadJson, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", realtimeAPIURL, bytes.NewBuffer(payloadJson))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, password)

	return req, nil
}

func doRealtimeRequest(p payload) ([]byte, error) {
	c := &http.Client{}

	req, err := createRequestWithAuthAndPayload(p)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		respBody, err := io.ReadAll(resp.Body)
		if err == nil {
			return nil, errors.New(fmt.Sprintf("received non-200 response code: %d (body: %s)", resp.StatusCode, respBody))
		}
		return nil, errors.New(fmt.Sprintf("received non-200 response code: %d ", resp.StatusCode))
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
