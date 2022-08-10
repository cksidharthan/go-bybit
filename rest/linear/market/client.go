package market

import (
	"github.com/cksidharthan/go-bybit/transport"
	httpTransport "github.com/cksidharthan/go-bybit/transport/http"
)

type LinearMarketClient struct {
	transporter transport.Transporter
}

func NewLinearMarketClient(url, apiKey, apiSecret string) *LinearMarketClient {
	transporter := httpTransport.New(url, apiKey, apiSecret)
	return &LinearMarketClient{
		transporter: transporter,
	}
}
