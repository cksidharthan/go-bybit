package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

func (c *Client) GetMergedOrderBook(ctx context.Context, params *spot.MergedOrderBookParams) (res *spot.MergedOrderBookResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteDepthMergedPath, params, &res)
	if err != nil {
		return
	}
	return
}
