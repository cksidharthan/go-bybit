package market

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cksidharthan/go-bybit/helpers"
	"github.com/cksidharthan/go-bybit/rest/domain/spot/market"
)

func (c *Client) GetBidAskPrice(ctx context.Context, symbol string) (bidAskPrice *market.PriceBidAskResponse, err error) {
	path := helpers.PublicSpotQuoteTickerBookTickerPath + "?symbol=" + symbol
	apiPath, err := url.Parse(path)
	if err != nil {
		return
	}
	payload, err := c.transporter.UnSignedRequest(ctx, apiPath, http.MethodGet, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &bidAskPrice)
	if err != nil {
		return
	}
	return
}
