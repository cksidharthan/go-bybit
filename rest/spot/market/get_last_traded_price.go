package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

func (c *Client) GetLastTradedPrice(ctx context.Context, params *spot.GetLastTradedPriceParams) (res *spot.GetLastTradedPriceResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteTickerPrice, params, &res)
	if err != nil {
		return
	}
	return
}
