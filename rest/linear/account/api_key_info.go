package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

func (c *LinearAccountClient) GetAPIKeyInfo(ctx context.Context) (res *linear.GetAPIKeyInfoResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetAPIKeyInfoPath, nil, &res)
	if err != nil {
		return
	}
	return
}
