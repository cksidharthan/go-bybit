package market

import (
	"github.com/cksidharthan/go-bybit/transport"
	httpTransport "github.com/cksidharthan/go-bybit/transport/http"
)

type SpotMarketClient struct {
	transporter transport.Transporter
}

// NewSpotMarketClient - creates a new bybit spot market client
//
// docs - https://bybit-exchange.github.io/docs/spot/#t-marketdata
func NewSpotMarketClient(url, apiKey, apiSecret string) *SpotMarketClient {
	transporter := httpTransport.New(url, apiKey, apiSecret)
	return &SpotMarketClient{
		transporter: transporter,
	}
}
