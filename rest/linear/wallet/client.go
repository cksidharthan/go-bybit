package wallet

import (
	"github.com/cksidharthan/go-bybit/transport"
	httpTransport "github.com/cksidharthan/go-bybit/transport/http"
)

type LinearWalletClient struct {
	transporter transport.Transporter
}

// NewLinearWalletClient - create a new linear wallet client.
func NewLinearWalletClient(url, apiKey, apiSecret string) *LinearWalletClient {
	transporter := httpTransport.New(url, apiKey, apiSecret)
	return &LinearWalletClient{
		transporter: transporter,
	}
}
