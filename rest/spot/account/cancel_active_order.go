package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

func (c *Client) CancelActiveOrder(ctx context.Context, params spot.CancelActiveOrderParams) (res *spot.CancelActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}
