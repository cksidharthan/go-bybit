package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

func (c *Client) GetLastFundingRate(ctx context.Context, params *linear.GetLastFundingRateParams) (res *linear.GetLastFundingRateResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicGetLastFundingRatePath, params, &res)
	if err != nil {
		return
	}
	return
}
