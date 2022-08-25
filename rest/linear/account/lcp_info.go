package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

// GetLCPInfo - get user's lcp info. [/v2/private/account/lcp - GET]
//
// this is a shared endpoint between USDT Perpetual(linear) and Inverse Perpetual(inverse)
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-lcp
func (c *LinearAccountClient) GetLCPInfo(ctx context.Context, params *linear.GetLCPInfoParams) (res *linear.GetLCPInfoResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetLCPInfoPath, params, &res)
	if err != nil {
		return
	}
	return
}
