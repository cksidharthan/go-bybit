package account

import (
	"github.com/cksidharthan/go-bybit/rest/http"
	"github.com/cksidharthan/go-bybit/rest/v5/domain/account"
)

type Client struct {
	http *http.Client
}

// New returns a new accounts client
func New(opts *http.ClientOpts) account.Client {
	return &Client{
		http: http.New(opts),
	}
}
