package transport

import (
	"context"
	"fmt"
	"net/url"
)

type Transporter interface {
	Call(ctx context.Context, apiPath *url.URL, method string, payload []byte, headers map[string]string) ([]byte, error)
}

type TransportError struct {
	StatusCode int    `json:"statusCode"`
	Msg        string `json:"msg"`
}

func (e *TransportError) Error() string {
	return fmt.Sprintf("Error: %s Statuscode:%d", e.Msg, e.StatusCode)
}
