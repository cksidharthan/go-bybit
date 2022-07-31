package account

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
	"github.com/google/go-querystring/query"
)

func (c *Client) BatchFastCancelActiveOrder(ctx context.Context, params spot.BatchFastCancelActiveOrderParams) (result *spot.BatchFastCancelActiveOrderResponse, err error) {
	apiPath, err := url.Parse(bybit.PrivateSpotOrderBatchFastCancelPath)
	if err != nil {
		return
	}

	queryString, err := query.Values(params)
	if err != nil {
		return
	}
	apiPath.RawQuery = queryString.Encode()

	payload, err := c.transporter.SignedRequest(ctx, apiPath, http.MethodDelete, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &result)
	if err != nil {
		return
	}
	return
}
