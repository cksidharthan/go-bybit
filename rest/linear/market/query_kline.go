package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

func (c *Client) QueryKline(ctx context.Context, params *linear.QueryKlineParams) (res *linear.QueryKlineResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicQueryKlinePath, params, &res)
	if err != nil {
		return
	}
	return
}
