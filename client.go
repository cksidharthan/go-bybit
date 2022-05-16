package go_bybit

import v2 "github.com/cksidharthan/go-bybit/v2"

type BybitClientInterface interface {
	BybitV1() v2.BybitV2ClientInterface
}

type BybitClient struct {
	bybitV2 v2.BybitV2ClientInterface
}

func (c *BybitClient) BybitV2() v2.BybitV2ClientInterface {
	return c.bybitV2
}

func New(url, apiKey, apiSecret string) *BybitClient {
	return &BybitClient{
		bybitV2: v2.New(url, apiKey, apiSecret),
	}
}
