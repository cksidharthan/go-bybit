package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

func (c *Client) QueryMarkPriceKline(ctx context.Context, params *linear.QueryMarkPriceKlineParams) (res *linear.QueryMarkPriceKlineResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicQueryMarkPriceKline, params, &res)
	if err != nil {
		return
	}
	return
}
