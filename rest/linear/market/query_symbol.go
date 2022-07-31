package market

import (
	"context"
	"encoding/json"
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
	"github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

func (c *Client) QuerySymbol(ctx context.Context, params *linear.QuerySymbolParams) (res *linear.QuerySymbolResponse, err error) {
	apiPath, err := url.Parse(bybit.PublicQuerySymbolPath)
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
