package wallet

import (
	"github.com/cksidharthan/go-bybit/transport"
	httpTransport "github.com/cksidharthan/go-bybit/transport/http"
)

type Client struct {
	transporter transport.Transporter
}

func New(url, apiKey, apiSecret string) *Client {
	transporter := httpTransport.New(url, apiKey, apiSecret)
	return &Client{
		transporter: transporter,
	}
}
