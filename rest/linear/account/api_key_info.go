package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

// GetAPIKeyInfo - get users api key info. [/linear/v1/account - GET]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-key
func (c *LinearAccountClient) GetAPIKeyInfo(ctx context.Context) (res *linear.GetAPIKeyInfoResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetAPIKeyInfoPath, nil, &res)
	if err != nil {
		return
	}
	return
}
