package market

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cksidharthan/go-bybit/helpers"
	"github.com/cksidharthan/go-bybit/rest/domain/spot/market"
)

func (c *Client) GetServerTime(ctx context.Context) (serverTime *market.ServerTimeResponse, err error) {
	apiPath, err := url.Parse(helpers.PublicServerTimePath)
	if err != nil {
		return
	}

	payload, err := c.transporter.UnSignedRequest(ctx, apiPath, http.MethodGet, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &serverTime)
	if err != nil {
		return
	}
	return
}
