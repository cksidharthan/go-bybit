package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

func (c *Client) GetOpenOrders(ctx context.Context, params spot.GetOpenOrdersParams) (res *spot.GetOpenOrdersResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateOpenOrdersPath, params, &res)
	if err != nil {
		return
	}
	return
}
