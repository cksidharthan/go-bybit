package wallet

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

// GetWalletBalance - get wallet balance. [/spot/v1/account - GET]
//
// docs  - https://bybit-exchange.github.io/docs/spot/#t-balance
func (c *SpotWalletClient) GetWalletBalance(ctx context.Context) (res *spot.GetWalletBalanceResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateWalletBalancePath, nil, &res)
	if err != nil {
		return
	}
	return
}
