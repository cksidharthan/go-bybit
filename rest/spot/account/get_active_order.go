package account

import (
	"context"
	"encoding/json"
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
	"github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

func (c *Client) GetActiveOrder(ctx context.Context, params spot.GetActiveOrderParams) (res *spot.GetActiveOrderResponse, err error) {
	apiPath, err := url.Parse(bybit.PrivateSpotOrderPath)
	if err != nil {
		return
	}

	queryString, err := query.Values(params)
	if err != nil {
		return
	}
	apiPath.RawQuery = queryString.Encode()

	payload, err := c.transporter.SignedRequest(ctx, apiPath, http.MethodGet, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &res)
	if err != nil {
		return
	}
	return
}
