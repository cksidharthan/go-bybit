package account

import (
	"context"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot/account/types"
	"github.com/google/go-querystring/query"
)

func (c *Client) PlaceActiveOrder(ctx context.Context, order types.SpotActiveOrderParams) (res *types.SpotActiveOrderResponse, err error) {
	queryString, err := query.Values(order)
	if err != nil {
		return
	}

	if err := c.transporter.SignedPostFormRequest(bybit.PrivateSpotOrderPath, queryString, &res); err != nil {
		return nil, err
	}
	return
}
