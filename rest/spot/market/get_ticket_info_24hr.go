package market

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cksidharthan/go-bybit/helpers"
	"github.com/cksidharthan/go-bybit/rest/domain/spot/market"
)

func (c *Client) GetTickerInfo24hr(ctx context.Context, symbol string) (tickerInfo24hr *market.TickerResponse, err error) {
	path := helpers.PublicSpotQuoteTicker24HrPath + "?symbol=" + symbol
	apiPath, err := url.Parse(path)
	if err != nil {
		return
	}
	payload, err := c.transporter.UnSignedRequest(ctx, apiPath, http.MethodGet, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &tickerInfo24hr)
	if err != nil {
		return
	}
	return
}
