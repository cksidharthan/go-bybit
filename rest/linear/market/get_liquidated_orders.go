package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

func (c *Client) GetLiquidatedOrders(ctx context.Context, params *linear.GetLiquidatedOrdersParams) (res *linear.GetLiquidatedOrdersResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicLiquidatedOrdersPath, params, &res)
	if err != nil {
		return
	}
	return
}
