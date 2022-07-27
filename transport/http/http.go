package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/cksidharthan/go-bybit/transport"
)

type HTTP struct {
	URI       string
	APIKey    string
	APISecret string
}

func New(url, apiKey, apiSecret string) *HTTP {
	return &HTTP{
		URI:       url,
		APIKey:    apiKey,
		APISecret: apiSecret,
	}
}

func (h *HTTP) UnSignedRequest(ctx context.Context, apiPath *url.URL, method string, payload []byte, headers map[string]string) ([]byte, error) {
	c := http.Client{
		Timeout: 5 * time.Second,
	}

	base, err := url.Parse(h.URI)
	if err != nil {
		return nil, &transport.Error{
			ErrorMsg: err.Error(),
			HTTPCode: -1,
		}
	}
	apiURL := base.ResolveReference(apiPath)

	req, err := http.NewRequestWithContext(ctx, method, apiURL.String(), bytes.NewReader(payload))
	if err != nil {
		return nil, &transport.Error{
			ErrorMsg: err.Error(),
			HTTPCode: -1,
		}
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, &transport.Error{
			ErrorMsg: err.Error(),
			HTTPCode: -1,
		}
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &transport.Error{
			ErrorMsg: err.Error(),
			HTTPCode: -1,
		}
	}

	// TODO: Handle error codes
	if resp.StatusCode > 399 {
		var bybitError transport.BybitError
		err = json.Unmarshal(respData, &bybitError)
		if err != nil {
			return nil, &transport.Error{
				ErrorMsg: err.Error(),
				HTTPCode: -1,
			}
		}
		return nil, &transport.Error{
			ErrorMsg: bybitError.RetMsg,
			HTTPCode: resp.StatusCode,
		}
	}

	return respData, nil
}
