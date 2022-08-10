package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/rest/domain/spot"

	"github.com/cksidharthan/go-bybit/bybit"
)

func (c *Client) GetServerTime(ctx context.Context) (res *spot.ServerTimeResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicServerTimePath, nil, &res)
	if err != nil {
		return
	}
	return
}
