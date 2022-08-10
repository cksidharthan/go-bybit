package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

func (c *Client) GetOrderBookDepth(ctx context.Context, params *spot.OrderBookDepthParams) (res *spot.OrderBookDepthResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteDepthPath, params, &res)
	if err != nil {
		return
	}
	return
}
