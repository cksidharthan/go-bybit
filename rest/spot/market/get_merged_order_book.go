package market

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cksidharthan/go-bybit/helpers"
	"github.com/cksidharthan/go-bybit/rest/domain/spot/market/filters"
	"github.com/cksidharthan/go-bybit/rest/domain/spot/market/types"
)

func (c *Client) GetMergedOrderBook(ctx context.Context, filters *filters.MergedOrderBookFilter) (orderBook *types.OrderBookResponse, err error) {
	apiPath, err := url.Parse(helpers.PublicSpotQuoteDepthMergedPath)
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
