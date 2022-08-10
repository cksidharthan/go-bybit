package account

import (
	"github.com/cksidharthan/go-bybit/transport"
	httpTransport "github.com/cksidharthan/go-bybit/transport/http"
)

type SpotAccountClient struct {
	transporter transport.Transporter
}

func NewSpotAccountClient(url, apiKey, apiSecret string) *SpotAccountClient {
	transporter := httpTransport.New(url, apiKey, apiSecret)
	return &SpotAccountClient{
		transporter: transporter,
	}
}
