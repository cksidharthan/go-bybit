package inverse_perpetual

import (
	"github.com/cksidharthan/go-bybit/rest/domain/inverse_perpetual"
	"github.com/cksidharthan/go-bybit/rest/inverse_perpetual/account"
	"github.com/cksidharthan/go-bybit/rest/inverse_perpetual/market"
	"github.com/cksidharthan/go-bybit/rest/inverse_perpetual/wallet"
)

type Client struct {
	market  inverse_perpetual.MarketInterface
	account inverse_perpetual.AccountInterface
	wallet  inverse_perpetual.WalletInterface
}

func (c *Client) Market() inverse_perpetual.MarketInterface {
	return c.market
}

func (c *Client) Account() inverse_perpetual.AccountInterface {
	return c.account
}

func (c *Client) Wallet() inverse_perpetual.WalletInterface {
	return c.wallet
}

func NewInversePerpetualClient(url, apiKey, apiSecret string) *Client {
	return &Client{
		market:  market.NewInversePerpetualMarketClient(url, apiKey, apiSecret),
		account: account.NewInversePerpetualAccountClient(url, apiKey, apiSecret),
		wallet:  wallet.NewInversePerpetualWalletClient(url, apiKey, apiSecret),
	}
}
