package market

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cksidharthan/go-bybit/rest/domain/spot/market/types"

	"github.com/cksidharthan/go-bybit/bybit"
)

func (c *Client) GetLastTradedPrice(ctx context.Context, symbol string) (lastTradedPrice *types.PriceResponse, err error) {
	path := bybit.PublicSpotQuoteTickerPrice + "?symbol=" + symbol
	apiPath, err := url.Parse(path)
	if err != nil {
		return
	}
	payload, err := c.transporter.UnSignedRequest(ctx, apiPath, http.MethodGet, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &lastTradedPrice)
	if err != nil {
		return
	}
	return
}
