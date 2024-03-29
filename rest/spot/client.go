package spot

import (
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
	"github.com/cksidharthan/go-bybit/rest/spot/account"
	"github.com/cksidharthan/go-bybit/rest/spot/market"
	"github.com/cksidharthan/go-bybit/rest/spot/wallet"
)

type Client struct {
	market  spot.MarketInterface
	account spot.AccountInterface
	wallet  spot.WalletInterface
}

func (c *Client) Market() spot.MarketInterface {
	return c.market
}

func (c *Client) Account() spot.AccountInterface {
	return c.account
}

func (c *Client) Wallet() spot.WalletInterface {
	return c.wallet
}

func NewSpotClient(url, apiKey, apiSecret string) *Client {
	return &Client{
		market:  market.NewSpotMarketClient(url, apiKey, apiSecret),
		account: account.NewSpotAccountClient(url, apiKey, apiSecret),
		wallet:  wallet.NewSpotWalletClient(url, apiKey, apiSecret),
	}
}
