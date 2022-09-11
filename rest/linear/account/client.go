package account

import (
	"github.com/cksidharthan/go-bybit/transport"
	httpTransport "github.com/cksidharthan/go-bybit/transport/http"
)

type LinearAccountClient struct {
	Transporter transport.Transporter
}

func NewLinearAccountClient(url, apiKey, apiSecret string) *LinearAccountClient {
	transporter := httpTransport.New(url, apiKey, apiSecret)
	return &LinearAccountClient{
		Transporter: transporter,
	}
}
