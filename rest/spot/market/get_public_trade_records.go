package market

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot/market/filters"
	"github.com/cksidharthan/go-bybit/rest/domain/spot/market/types"
)

func (c *Client) GetPublicTradeRecords(ctx context.Context, filters *filters.TradeRecordsFilter) (publicTradeRecords *types.TradeRecordsResponse, err error) {
	apiPath, err := url.Parse(bybit.PublicSpotQuoteTradesPath)
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
