package market

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cksidharthan/go-bybit/rest/domain/spot/market/filters"
	"github.com/cksidharthan/go-bybit/rest/domain/spot/market/types"

	"github.com/cksidharthan/go-bybit/helpers"
)

func (c *Client) GetKline(ctx context.Context, filters *filters.KlineFilter) (kline *types.KlineResponse, err error) {
	apiPath, err := url.Parse(helpers.PublicSpotQuoteKlinePath)
	if err != nil {
		return
	}
	if filters != nil {
		filters.ToQuery(apiPath)
	}

	payload, err := c.transporter.UnSignedRequest(ctx, apiPath, http.MethodGet, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &kline)
	if err != nil {
		return
	}
	return
}
