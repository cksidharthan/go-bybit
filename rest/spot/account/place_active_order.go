package account

import (
	"context"
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
	"github.com/google/go-querystring/query"
)

func (c *Client) PlaceActiveOrder(ctx context.Context, params spot.PlaceActiveOrderParams) (res *spot.PlaceActiveOrderResponse, err error) {
	queryString, err := query.Values(params)
	if err != nil {
		return
	}

	if err := c.transporter.SignedPostForm(bybit.PrivateSpotOrderPath, queryString, &res); err != nil {
		return nil, err
	}
	return
}
