package market

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cksidharthan/go-bybit/helpers"
	"github.com/cksidharthan/go-bybit/rest/domain/spot/market"
	"github.com/cksidharthan/go-bybit/rest/filters"
)

func (c *Client) GetOrderBookDepth(ctx context.Context, filters *filters.OrderBookDepthFilter) (orderBook *market.OrderBookResponse, err error) {
	apiPath, err := url.Parse(helpers.PublicSpotQuoteDepthPath)
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

	err = json.Unmarshal(payload, &orderBook)
	if err != nil {
		return
	}
	return
}
