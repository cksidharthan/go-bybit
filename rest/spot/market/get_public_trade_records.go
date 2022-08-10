package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/rest/domain/spot"

	"github.com/cksidharthan/go-bybit/bybit"
)

func (c *Client) GetPublicTradeRecords(ctx context.Context, params *spot.PublicTradeRecordsParams) (res *spot.PublicTradeRecordsResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteTradesPath, params, &res)
	if err != nil {
		return
	}
	return
}
