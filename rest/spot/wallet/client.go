package wallet

import (
	"github.com/cksidharthan/go-bybit/transport"
	httpTransport "github.com/cksidharthan/go-bybit/transport/http"
)

type SpotWalletClient struct {
	transporter transport.Transporter
}

func NewSpotWalletClient(url, apiKey, apiSecret string) *SpotWalletClient {
	transporter := httpTransport.New(url, apiKey, apiSecret)
	return &SpotWalletClient{
		transporter: transporter,
	}
}
