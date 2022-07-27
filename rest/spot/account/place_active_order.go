package account

import (
	"context"

	"github.com/cksidharthan/go-bybit/rest/domain/spot"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/google/go-querystring/query"
)

func (c *Client) PlaceActiveOrder(ctx context.Context, order spot.ActiveOrderParams) (res *spot.ActiveOrderResponse, err error) {
	queryString, err := query.Values(order)
	if err != nil {
		return
	}

	if err := c.transporter.SignedPostFormRequest(bybit.PrivateSpotOrderPath, queryString, &res); err != nil {
		return nil, err
	}
	return
}
