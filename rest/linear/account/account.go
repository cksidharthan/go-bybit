package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

func (c *LinearAccountClient) PlaceActiveOrder(ctx context.Context, params *linear.PlaceActiveOrderParams) (res *linear.PlaceActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearPlaceOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}
