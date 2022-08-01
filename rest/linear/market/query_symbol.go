package market

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

func (c *Client) QuerySymbol(ctx context.Context) (res *linear.QuerySymbolResponse, err error) {
	apiPath, err := url.Parse(bybit.PublicQuerySymbolPath)
	if err != nil {
		return
	}

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
