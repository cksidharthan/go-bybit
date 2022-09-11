package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

// PlaceActiveOrder - Place an active linear order. [/private/linear/order/create - POST]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-placeactive
func (c *LinearAccountClient) PlaceActiveOrder(ctx context.Context, params *linear.PlaceActiveOrderParams) (res *linear.PlaceActiveOrderResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearPlaceOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

// GetActiveOrder - Get active linear orders. [/private/linear/order/list - GET]
//
// Get my active order list.
// Because order creation/cancellation is asynchronous, there can be a data delay in this endpoint. You can get real-time order info with the Query Active Order (real-time) endpoint.
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-getactive
func (c *LinearAccountClient) GetActiveOrder(ctx context.Context, params *linear.GetActiveOrderParams) (res *linear.GetActiveOrderResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetActiveOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

// CancelActiveOrder - Cancel an active linear order. [/private/linear/order/cancel - POST]
//
// Either order_id or order_link_id are required for cancelling active orders. order_id - this unique 36 characters order ID was returned to you when the active order was created successfully.
// You may cancel active orders that are unfilled or partially filled. Fully filled orders cannot be cancelled.
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-cancelactive
func (c *LinearAccountClient) CancelActiveOrder(ctx context.Context, params *linear.CancelActiveOrderParams) (res *linear.CancelActiveOrderResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearOrderCancelPath, params, &res)
	if err != nil {
		return
	}
	return
}

// CancelAllActiveOrders - Cancel all active linear orders. [/private/linear/order/cancel-all - POST]
//
// Cancel all active orders that are unfilled or partially filled. Fully filled orders cannot be cancelled.
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-cancelactive
func (c *LinearAccountClient) CancelAllActiveOrders(ctx context.Context, params *linear.CancelAllActiveOrdersParams) (res *linear.CancelAllActiveOrdersResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearOrderCancelAllPath, params, &res)
	if err != nil {
		return
	}
	return
}

// ReplaceActiveOrder - Replace order can modify/amend your active orders. [/private/linear/order/replace - POST]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-replaceactive
func (c *LinearAccountClient) ReplaceActiveOrder(ctx context.Context, params *linear.ReplaceActiveOrderParams) (res *linear.ReplaceActiveOrderResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearReplaceActiveOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

// QueryActiveOrder - Query active linear orders. [/private/linear/order/query - GET]
//
// Query real-time active order information. If only order_id or order_link_id are passed, a single order will be returned; otherwise, returns up to 500 unfilled orders.
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-queryactive
func (c *LinearAccountClient) QueryActiveOrder(ctx context.Context, params *linear.QueryActiveOrderParams) (res *linear.QueryActiveOrderResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearQueryActiveOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}
