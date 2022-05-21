package public

import "github.com/cksidharthan/go-bybit/public/spot"

type PublicInterface interface {
	Spot() spot.SpotInterface
}

type PublicClient struct {
	spot spot.SpotInterface
}

func (c *PublicClient) Spot() spot.SpotInterface {
	return c.spot
}

func New(url string) PublicInterface {
	return &PublicClient{
		spot: spot.New(url),
	}
}
