package market

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cksidharthan/go-bybit/rest/domain/spot"

	"github.com/cksidharthan/go-bybit/bybit"
)

func (c *Client) GetTickerInfo24hr(ctx context.Context, symbol string) (tickerInfo24hr *spot.TickerResponse, err error) {
	path := bybit.PublicSpotQuoteTicker24HrPath + "?symbol=" + symbol
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
