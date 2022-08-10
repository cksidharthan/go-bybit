package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

func (c *Client) GetTickerInfo24hr(ctx context.Context, params *spot.GetTickerInfo24hrParams) (res *spot.GetTickerInfoResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteTicker24HrPath, params, &res)
	if err != nil {
		return
	}
	return
}
