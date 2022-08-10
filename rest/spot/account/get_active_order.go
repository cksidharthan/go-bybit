package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

func (c *Client) GetActiveOrder(ctx context.Context, params spot.GetActiveOrderParams) (res *spot.GetActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateSpotOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}
