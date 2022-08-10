package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

func (c *Client) GetTradeHistory(ctx context.Context, params spot.GetTradeHistoryParams) (res *spot.GetTradeHistoryResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateTradeHistoryPath, params, &res)
	if err != nil {
		return
	}
	return
}
