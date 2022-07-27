package market

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cksidharthan/go-bybit/helpers"
	"github.com/cksidharthan/go-bybit/rest/domain/spot/market"
)

func (c *Client) GetSymbols(ctx context.Context) (symbols *market.SymbolsResponse, err error) {
	apiPath, err := url.Parse(helpers.PublicSpotSymbolsPath)
	if err != nil {
		return
	}

	payload, err := c.transporter.UnSignedRequest(ctx, apiPath, http.MethodGet, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &symbols)
	if err != nil {
		return
	}
	return
}
