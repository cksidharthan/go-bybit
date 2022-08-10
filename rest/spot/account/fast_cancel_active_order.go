package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

func (c *Client) FastCancelActiveOrder(ctx context.Context, params spot.FastCancelActiveOrderParams) (res *spot.FastCancelActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotOrderFastPath, params, &res)
	if err != nil {
		return
	}
	return
}
