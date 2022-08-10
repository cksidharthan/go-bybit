package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

func (c *Client) GetOrderBook(ctx context.Context, params *linear.OrderBookParams) (res *linear.OrderBookResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicLinearOrderBookPath, params, &res)
	if err != nil {
		return
	}
	return
}
