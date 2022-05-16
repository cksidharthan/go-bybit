package http

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/cksidharthan/go-bybit/transport"
)

type Http struct {
	Uri   string
	Token string
}

func New(url, token string) *Http {
	return &Http{
		Uri:   url,
		Token: token,
	}
}

func (h *Http) Call(ctx context.Context, apiPath *url.URL, method string, payload []byte, headers map[string]string) ([]byte, error) {
	c := http.Client{
		Timeout: 30 * time.Second,
	}

	base, err := url.Parse(h.Uri)
	if err != nil {
		return nil, &transport.TransportError{
			StatusCode: 0,
			Msg:        err.Error(),
		}
	}
	apiURL := base.ResolveReference(apiPath)

	req, err := http.NewRequestWithContext(ctx, method, apiURL.String(), bytes.NewReader(payload))
	if err != nil {
		return nil, &transport.TransportError{
			StatusCode: 0,
			Msg:        err.Error(),
		}
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, &transport.TransportError{
			StatusCode: 0,
			Msg:        err.Error(),
		}
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &transport.TransportError{
			StatusCode: 0,
			Msg:        err.Error(),
		}
	}

	// TODO: Handle error codes
	if resp.StatusCode > 399 {
		return nil, &transport.TransportError{
			StatusCode: resp.StatusCode,
			Msg:        "error getting response from bybit",
		}
	}

	return respData, nil
}
