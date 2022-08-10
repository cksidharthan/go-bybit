package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

func (c *Client) GetOrderHistory(ctx context.Context, params spot.GetOrderHistoryParams) (res *spot.GetOrderHistoryResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateOrderHistoryPath, params, &res)
	if err != nil {
		return
	}
	return
}
