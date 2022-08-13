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

func (c *LinearAccountClient) GetActiveOrder(ctx context.Context, params *linear.GetActiveOrderParams) (res *linear.GetActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetActiveOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) CancelActiveOrder(ctx context.Context, params *linear.CancelActiveOrderParams) (res *linear.CancelActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearOrderCancelPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) CancelAllActiveOrders(ctx context.Context, params *linear.CancelAllActiveOrdersParams) (res *linear.CancelAllActiveOrdersResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearOrderCancelAllPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) ReplaceActiveOrder(ctx context.Context, params *linear.ReplaceActiveOrderParams) (res *linear.ReplaceActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearReplaceActiveOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) QueryActiveOrder(ctx context.Context, params *linear.QueryActiveOrderParams) (res *linear.QueryActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearQueryActiveOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}