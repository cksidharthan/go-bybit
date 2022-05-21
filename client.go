package go_bybit

import "github.com/cksidharthan/go-bybit/public"

type BybitClientInterface interface {
	Public() public.PublicInterface
}

type BybitClient struct {
	public public.PublicInterface
}

func (c *BybitClient) BybitV2() public.PublicInterface {
	return c.public
}

func New(url, apiKey, apiSecret string) *BybitClient {
	return &BybitClient{
		public: public.New(url),
	}
}
