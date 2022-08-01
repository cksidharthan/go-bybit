package linear

import (
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
	"github.com/cksidharthan/go-bybit/rest/linear/account"
	"github.com/cksidharthan/go-bybit/rest/linear/market"
	"github.com/cksidharthan/go-bybit/rest/linear/wallet"
)

type Interface interface {
	Market() linear.MarketInterface
	Account() linear.AccountInterface
	Wallet() linear.WalletInterface
}

type Client struct {
	market  linear.MarketInterface
	account linear.AccountInterface
	wallet  linear.WalletInterface
}

func (c *Client) Market() linear.MarketInterface {
	return c.market
}

func (c *Client) Account() linear.AccountInterface {
	return c.account
}

func (c *Client) Wallet() linear.WalletInterface {
	return c.wallet
}

func New(url, apiKey, apiSecret string) Interface {
	return &Client{
		market:  market.New(url, apiKey, apiSecret),
		account: account.New(url, apiKey, apiSecret),
		wallet:  wallet.New(url, apiKey, apiSecret),
	}
}
