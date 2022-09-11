package account

import (
	"github.com/cksidharthan/go-bybit/transport"
	httpTransport "github.com/cksidharthan/go-bybit/transport/http"
)

type SpotAccountClient struct {
	Transporter transport.Transporter
}

// NewSpotAccountClient - creates a new bybit spot account client
func NewSpotAccountClient(url, apiKey, apiSecret string) *SpotAccountClient {
	transporter := httpTransport.New(url, apiKey, apiSecret)
	return &SpotAccountClient{
		Transporter: transporter,
	}
}
