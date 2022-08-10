package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

func (c *Client) GetBidAskPrice(ctx context.Context, params *spot.GetBidAskPriceParams) (res *spot.BidAskPriceResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteTickerBookTickerPath, params, &res)
	if err != nil {
		return
	}
	return
}
