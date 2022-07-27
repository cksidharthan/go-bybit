package market

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot/market/types"
)

func (c *Client) GetBidAskPrice(ctx context.Context, symbol string) (bidAskPrice *types.PriceBidAskResponse, err error) {
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
