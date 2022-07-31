package wallet

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

func (c *Client) GetWalletBalance(ctx context.Context) (res *spot.GetWalletBalanceResponse, err error) {
	apiPath, err := url.Parse(bybit.PrivateWalletBalancePath)
	if err != nil {
		return
	}

	payload, err := c.transporter.SignedRequest(ctx, apiPath, http.MethodGet, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &res)
	if err != nil {
		return
	}
	return
}
