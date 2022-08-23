package inverseperp

import (
	inverseperp "github.com/cksidharthan/go-bybit/rest/domain/inverse_perpetual"
	"github.com/cksidharthan/go-bybit/rest/inverse_perpetual/account"
	"github.com/cksidharthan/go-bybit/rest/inverse_perpetual/market"
	"github.com/cksidharthan/go-bybit/rest/inverse_perpetual/wallet"
)

type Client struct {
	market  inverseperp.MarketInterface
	account inverseperp.AccountInterface
	wallet  inverseperp.WalletInterface
}

func (c *Client) Market() inverseperp.MarketInterface {
	return c.market
}

func (c *Client) Account() inverseperp.AccountInterface {
	return c.account
}

func (c *Client) Wallet() inverseperp.WalletInterface {
	return c.wallet
}

func NewInversePerpetualClient(url, apiKey, apiSecret string) *Client {
	return &Client{
		market:  market.NewInversePerpetualMarketClient(url, apiKey, apiSecret),
		account: account.NewInversePerpetualAccountClient(url, apiKey, apiSecret),
		wallet:  wallet.NewInversePerpetualWalletClient(url, apiKey, apiSecret),
	}
}
