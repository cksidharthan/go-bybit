package market

import (
	"github.com/cksidharthan/go-bybit/transport"
	httpTransport "github.com/cksidharthan/go-bybit/transport/http"
)

type InversePerpetualMarketClient struct {
	Transporter transport.Transporter
}

// NewInversePerpetualMarketClient - create a new inverse perpetual market client.
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-marketdata
func NewInversePerpetualMarketClient(url, apiKey, apiSecret string) *InversePerpetualMarketClient {
	transporter := httpTransport.New(url, apiKey, apiSecret)
	return &InversePerpetualMarketClient{
		Transporter: transporter,
	}
}
