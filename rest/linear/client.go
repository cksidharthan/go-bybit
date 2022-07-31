package linear

import (
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
	"github.com/cksidharthan/go-bybit/rest/spot/account"
	"github.com/cksidharthan/go-bybit/rest/spot/market"
	"github.com/cksidharthan/go-bybit/rest/spot/wallet"
)

type Interface interface {
	Market() spot.MarketInterface
	Account() spot.AccountInterface
	Wallet() spot.WalletInterface
}

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

func New(url, apiKey, apiSecret string) Interface {
	return &Client{
		market:  market.New(url, apiKey, apiSecret),
		account: account.New(url, apiKey, apiSecret),
		wallet:  wallet.New(url, apiKey, apiSecret),
	}
}
