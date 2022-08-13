package account

import (
	"context"
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
	"net/http"
)

func (c *LinearAccountClient) PlaceConditionalOrder(ctx context.Context, params *linear.PlaceConditionalOrderParams) (res *linear.PlaceConditionalOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivatePlaceConditionalOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) GetConditionalOrder(ctx context.Context, params *linear.GetConditionalOrderParams) (res *linear.GetConditionalOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateGetConditionalOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) CancelConditionalOrder(ctx context.Context, params *linear.CancelConditionalOrderParams) (res *linear.CancelConditionalOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateCancelConditionalOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}
