package wallet

import (
	"github.com/cksidharthan/go-bybit/transport"
	httpTransport "github.com/cksidharthan/go-bybit/transport/http"
)

type SpotWalletClient struct {
	Transporter transport.Transporter
}

// NewSpotWalletClient - creates a new bybit spot wallet client
// docs - https://bybit-exchange.github.io/docs/spot/#t-wallet
func NewSpotWalletClient(url, apiKey, apiSecret string) *SpotWalletClient {
	transporter := httpTransport.New(url, apiKey, apiSecret)
	return &SpotWalletClient{
		Transporter: transporter,
	}
}
