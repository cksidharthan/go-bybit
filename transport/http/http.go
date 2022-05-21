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

type Http struct {
	Uri       string
	ApiKey    string
	ApiSecret string
}

func New(url, apiKey, apiSecret string) *Http {
	return &Http{
		Uri:       url,
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
	}
}

func (h *Http) UnSignedRequest(ctx context.Context, apiPath *url.URL, method string, payload []byte, headers map[string]string) ([]byte, error) {
	c := http.Client{
		Timeout: 5 * time.Second,
	}

	base, err := url.Parse(h.Uri)
	if err != nil {
		return nil, &transport.TransportError{
			ErrorMsg: err.Error(),
			HTTPCode: -1,
		}
	}
	apiURL := base.ResolveReference(apiPath)

	req, err := http.NewRequestWithContext(ctx, method, apiURL.String(), bytes.NewReader(payload))
	if err != nil {
		return nil, &transport.TransportError{
			ErrorMsg: err.Error(),
			HTTPCode: -1,
		}
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, &transport.TransportError{
			ErrorMsg: err.Error(),
			HTTPCode: -1,
		}
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &transport.TransportError{
			ErrorMsg: err.Error(),
			HTTPCode: -1,
		}
	}

	// TODO: Handle error codes
	if resp.StatusCode > 399 {
		var bybitError transport.BybitError
		err = json.Unmarshal(respData, &bybitError)
		if err != nil {
			return nil, &transport.TransportError{
				ErrorMsg: err.Error(),
				HTTPCode: -1,
			}
		}
		return nil, &transport.TransportError{
			ErrorMsg: bybitError.RetMsg,
			HTTPCode: resp.StatusCode,
		}
	}

	return respData, nil
}
