package market

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
	"github.com/google/go-querystring/query"
)

func (c *Client) QueryKline(ctx context.Context, params *linear.QueryKlineParams) (res *linear.QueryKlineResponse, err error) {
	apiPath, err := url.Parse(bybit.PublicQueryKlinePath)
	if err != nil {
		return
	}

	queryString, err := query.Values(params)
	if err != nil {
		return
	}
	apiPath.RawQuery = queryString.Encode()

	payload, err := c.transporter.UnSignedRequest(ctx, apiPath, http.MethodGet, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &res)
	if err != nil {
		return
	}
	return
}
