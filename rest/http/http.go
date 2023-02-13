package http

import (
	"net/http"
)

type Client struct {
	BaseURL   string
	Http      *http.Client
	Useragent string
}

type ClientOpts struct {
	// BaseURL is the base URL for the API.
	BaseURL string
	// UserAgent is the user agent used when communicating with the API.
	UserAgent string
	// HTTPClient is the HTTP client used to communicate with the API.
	HTTPClient *http.Client
	// Debug is a flag that enables debug logging.
}

func New(opts *ClientOpts) *Client {
	if opts == nil {
		opts = &ClientOpts{}
	}

	if opts.HTTPClient == nil {
		opts.HTTPClient = http.DefaultClient
	}

	return &Client{
		BaseURL:   opts.BaseURL,
		Http:      opts.HTTPClient,
		Useragent: opts.UserAgent,
	}
}
