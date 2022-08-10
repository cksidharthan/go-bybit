package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/rest/domain/spot"

	"github.com/cksidharthan/go-bybit/bybit"
)

func (c *Client) GetSymbols(ctx context.Context) (symbols *spot.GetSymbolsResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotSymbolsPath, nil, &symbols)
	if err != nil {
		return
	}
	return
}
