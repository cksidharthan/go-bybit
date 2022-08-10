package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

func (c *Client) GetKline(ctx context.Context, params *spot.GetKlineParams) (res *spot.KlineResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteKlinePath, params, &res)
	if err != nil {
		return
	}
	return
}
