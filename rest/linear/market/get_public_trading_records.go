package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

func (c *Client) GetPublicTradingRecords(ctx context.Context, params *linear.GetPublicTradingRecordsParams) (res *linear.GetPublicTradingRecordsResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicGetPublicTradingRecordsPath, params, &res)
	if err != nil {
		return
	}
	return
}
