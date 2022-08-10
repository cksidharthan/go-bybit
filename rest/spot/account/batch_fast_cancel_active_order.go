package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

func (c *Client) BatchFastCancelActiveOrder(ctx context.Context, params spot.BatchFastCancelActiveOrderParams) (res *spot.BatchFastCancelActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotOrderBatchFastCancelPath, params, &res)
	if err != nil {
		return
	}
	return
}
