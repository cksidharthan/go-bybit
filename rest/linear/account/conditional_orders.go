package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

func (c *LinearAccountClient) PlaceConditionalOrder(ctx context.Context, params *linear.PlaceConditionalOrderParams) (res *linear.PlaceConditionalOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearPlaceConditionalOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) GetConditionalOrder(ctx context.Context, params *linear.GetConditionalOrderParams) (res *linear.GetConditionalOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetConditionalOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) CancelConditionalOrder(ctx context.Context, params *linear.CancelConditionalOrderParams) (res *linear.CancelConditionalOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearCancelConditionalOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) CancelAllConditionalOrders(ctx context.Context, params *linear.CancelAllConditionalOrdersParams) (res *linear.CancelAllConditionalOrdersResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearCancelAllConditionalOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) ReplaceConditionalOrder(ctx context.Context, params *linear.ReplaceConditionalOrderParams) (res *linear.ReplaceConditionalOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearReplaceConditionalOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) QueryConditionalOrderBySymbol(ctx context.Context, params *linear.QueryConditionalOrderBySymbolParams) (res *linear.QueryConditionalOrderBySymbolResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearQueryConditionalOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) QueryConditionalOrderWithIDs(ctx context.Context, params *linear.QueryConditionalOrderWithIDsParams) (res *linear.QueryConditionalOrderWithIDsResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearQueryConditionalOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}
