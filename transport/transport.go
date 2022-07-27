package transport

import (
	"context"
	"fmt"
	"net/url"
)

type BybitError struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	Result  string `json:"result"`
	Token   string `json:"token"`
}

type Error struct {
	HTTPCode int    `json:"httpCode"`
	ErrorMsg string `json:"errorMsg"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.HTTPCode, e.ErrorMsg)
}

type Transporter interface {
	UnSignedRequest(ctx context.Context, apiPath *url.URL, method string, payload []byte, headers map[string]string) ([]byte, error)
	SignedRequest(ctx context.Context, apiPath *url.URL, method string, payload []byte, headers map[string]string) ([]byte, error)
}
