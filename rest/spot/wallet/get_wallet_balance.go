package wallet

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

func (c *Client) GetWalletBalance(ctx context.Context) (res *spot.GetWalletBalanceResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateWalletBalancePath, nil, &res)
	if err != nil {
		return
	}
	return
}
