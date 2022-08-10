package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

func (c *Client) BatchCancelActiveOrder(ctx context.Context, params spot.BatchCancelActiveOrderParams) (res *spot.BatchCancelActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotOrderBatchCancelPath, params, &res)
	if err != nil {
		return
	}
	return
}
