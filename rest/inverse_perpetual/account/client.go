package account

import (
	"github.com/cksidharthan/go-bybit/transport"
	httpTransport "github.com/cksidharthan/go-bybit/transport/http"
)

type InversePerpetualAccountClient struct {
	transporter transport.Transporter
}

func NewInversePerpetualAccountClient(url, apiKey, apiSecret string) *InversePerpetualAccountClient {
	transporter := httpTransport.New(url, apiKey, apiSecret)
	return &InversePerpetualAccountClient{
		transporter: transporter,
	}
}
