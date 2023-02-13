package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Do(ctx context.Context, method string, url string, payload interface{}, headers map[string]string, response interface{}) error {
	var body io.Reader
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		body = io.NopCloser(bytes.NewReader(b))
	}

	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s%s", c.BaseURL, url), body)
	if err != nil {
		return &Error{
			StatusCode: http.StatusBadGateway,
			Message:    err.Error(),
		}
	}

	req.Header.Set("User-Agent", c.Useragent)
	// Add additional headers if specified by the caller
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := c.Http.Do(req)
	if err != nil {
		return &Error{
			StatusCode: http.StatusBadGateway,
			Message:    err.Error(),
		}
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return &Error{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return &Error{
			StatusCode: resp.StatusCode,
			Message:    string(respData),
		}
	}

	// If the caller has specified a response object, unmarshal the response data into it
	if response != nil {
		if err := json.Unmarshal(respData, response); err != nil {
			return &Error{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}
	}

	return nil
}

func addHeaders(req *http.Request) {
	req.Header.Set("User-Agent", "go-bybit")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-BAPI-API-KEY", "API_KEY")
	req.Header.Set("X-BAPI-SIGN", "SIGNATURE")
	req.Header.Set("X-BAPI-TIMESTAMP", "TIMESTAMP")
	req.Header.Set("X-BAPI-RECV-WINDOW", "EXPIRE")
}
