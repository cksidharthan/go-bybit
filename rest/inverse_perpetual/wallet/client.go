package wallet

import (
	"github.com/cksidharthan/go-bybit/transport"
	httpTransport "github.com/cksidharthan/go-bybit/transport/http"
)

type InversePerpetualWalletClient struct {
	transporter transport.Transporter
}

func NewInversePerpetualWalletClient(url, apiKey, apiSecret string) *InversePerpetualWalletClient {
	transporter := httpTransport.New(url, apiKey, apiSecret)
	return &InversePerpetualWalletClient{
		transporter: transporter,
	}
}
