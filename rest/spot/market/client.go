package market

import (
	"github.com/cksidharthan/go-bybit/transport"
	httpTransport "github.com/cksidharthan/go-bybit/transport/http"
)

type SpotMarketClient struct {
	transporter transport.Transporter
}

func NewSpotMarketClient(url, apiKey, apiSecret string) *SpotMarketClient {
	transporter := httpTransport.New(url, apiKey, apiSecret)
	return &SpotMarketClient{
		transporter: transporter,
	}
}
