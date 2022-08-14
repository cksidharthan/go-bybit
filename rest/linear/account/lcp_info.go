package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

func (c *LinearAccountClient) GetLCPInfo(ctx context.Context, params *linear.GetLCPInfoParams) (res *linear.GetLCPInfoResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetLCPInfoPath, params, &res)
	if err != nil {
		return
	}
	return
}
