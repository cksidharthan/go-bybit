package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

// PlaceConditionalOrder - place conditional order. [/private/linear/stop-order/create - POST]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-placecond
func (c *LinearAccountClient) PlaceConditionalOrder(ctx context.Context, params *linear.PlaceConditionalOrderParams) (res *linear.PlaceConditionalOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearPlaceConditionalOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

// GetConditionalOrder - Gets conditional order [/private/linear/stop-order/list - GET]
//
// Get my conditional order list.
// Because order creation/cancellation is asynchronous, there can be a data delay in this endpoint. You can get real-time order info with the Query Conditional Order (real-time) endpoint.
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-getcond
func (c *LinearAccountClient) GetConditionalOrder(ctx context.Context, params *linear.GetConditionalOrderParams) (res *linear.GetConditionalOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetConditionalOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

// CancelConditionalOrder - cancel conditional order. [/private/linear/stop-order/cancel - POST]
//
// You may cancel all untriggered conditional orders or take profit/stop loss order. Essentially, after a conditional order is triggered, it will become an active order.
// So, when a conditional order is triggered, cancellation has to be done through the active order endpoint for any unfilled or partially filled active order. As always, orders that have been fully filled cannot be cancelled.
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-cancelcond
func (c *LinearAccountClient) CancelConditionalOrder(ctx context.Context, params *linear.CancelConditionalOrderParams) (res *linear.CancelConditionalOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearCancelConditionalOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

// CancelAllConditionalOrders - Cancel all untriggered conditional orders.. [/private/linear/stop-order/cancel-all - POST]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-cancelallcond
func (c *LinearAccountClient) CancelAllConditionalOrders(ctx context.Context, params *linear.CancelAllConditionalOrdersParams) (res *linear.CancelAllConditionalOrdersResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearCancelAllConditionalOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

// ReplaceConditionalOrder - replace conditional order. [/private/linear/stop-order/replace - POST]
//
// Replace conditional order can modify/amend your conditional orders.
// order_id and symbol are required for identifying a conditional order.
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-replacecond
func (c *LinearAccountClient) ReplaceConditionalOrder(ctx context.Context, params *linear.ReplaceConditionalOrderParams) (res *linear.ReplaceConditionalOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearReplaceConditionalOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

// QueryConditionalOrderBySymbol - Query conditional order by symbol. [/private/linear/stop-order/list - GET]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-querycond
func (c *LinearAccountClient) QueryConditionalOrderBySymbol(ctx context.Context, params *linear.QueryConditionalOrderBySymbolParams) (res *linear.QueryConditionalOrderBySymbolResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearQueryConditionalOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

// QueryConditionalOrderWithIDs - Query conditional order with IDs. [/private/linear/stop-order/list - GET]
//
// Query real-time stop order information. When passing the parameter order_id or order_link_id, a single order data will be returned; otherwise, returns up to 10 unfilled orders.
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-querycond
func (c *LinearAccountClient) QueryConditionalOrderWithIDs(ctx context.Context, params *linear.QueryConditionalOrderWithIDsParams) (res *linear.QueryConditionalOrderWithIDsResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearQueryConditionalOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}
