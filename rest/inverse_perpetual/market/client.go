package market

import (
	"github.com/cksidharthan/go-bybit/transport"
	httpTransport "github.com/cksidharthan/go-bybit/transport/http"
)

type InversePerpetualMarketClient struct {
	transporter transport.Transporter
}

func NewInversePerpetualMarketClient(url, apiKey, apiSecret string) *InversePerpetualMarketClient {
	transporter := httpTransport.New(url, apiKey, apiSecret)
	return &InversePerpetualMarketClient{
		transporter: transporter,
	}
}