package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

func (c *Client) QuerySymbol(ctx context.Context) (res *linear.QuerySymbolResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicQuerySymbolPath, nil, &res)
	if err != nil {
		return
	}
	return
}
