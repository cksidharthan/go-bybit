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

func (c *Client) GetPublicTradeRecords(ctx context.Context, filters *filters.TradeRecordsFilter) (publicTradeRecords *market.TradeRecordsResponse, err error) {
	apiPath, err := url.Parse(helpers.PublicSpotQuoteTradesPath)
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

	err = json.Unmarshal(payload, &publicTradeRecords)
	if err != nil {
		return
	}
	return
}
