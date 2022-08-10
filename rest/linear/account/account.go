package account

import (
	"context"
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
	"net/http"
)

func (c *Client) PlaceActiveOrder(ctx context.Context, params *linear.PlaceActiveOrderParams) (res *linear.PlaceActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearPlaceOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}
