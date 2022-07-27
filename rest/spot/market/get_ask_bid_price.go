package market

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cksidharthan/go-bybit/rest/domain/spot"

	"github.com/cksidharthan/go-bybit/bybit"
)

func (c *Client) GetBidAskPrice(ctx context.Context, symbol string) (bidAskPrice *spot.PriceBidAskResponse, err error) {
	path := bybit.PublicSpotQuoteTickerBookTickerPath + "?symbol=" + symbol
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
