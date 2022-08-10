package wallet

import (
	"context"
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
	"net/http"
)

func (c *Client) GetWalletBalance(ctx context.Context) (res *spot.GetWalletBalanceResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateWalletBalancePath, nil, &res)
	if err != nil {
		return
	}
	return
}
