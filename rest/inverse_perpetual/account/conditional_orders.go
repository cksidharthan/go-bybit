package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	inverseperp "github.com/cksidharthan/go-bybit/rest/domain/inverse_perpetual"
)

/*
PlaceConditionalOrder - place conditional order. [ /v2/private/stop-order/create  - POST ]

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-placecond
*/
func (c *InversePerpetualAccountClient) PlaceConditionalOrder(ctx context.Context, params *inverseperp.PlaceConditionalOrderParams) (response *inverseperp.PlaceConditionalOrderResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateInversePerpetualPlaceConditionalOrderPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
GetConditionalOrder - get conditional order. [ /v2/private/stop-order/list  - GET ]

Get my conditional order list.

Because order creation/cancellation is asynchronous, there can be a data delay in this endpoint. You can get real-time order info with the Query Conditional Order (real-time) endpoint.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-getcond
*/
func (c *InversePerpetualAccountClient) GetConditionalOrder(ctx context.Context, params *inverseperp.GetConditionalOrderParams) (response *inverseperp.GetConditionalOrderResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualGetConditionalOrderPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
CancelConditionalOrder - cancel conditional order. [ /v2/private/stop-order/cancel  - POST ]

You may cancel all untriggered conditional orders or take profit/stop loss order. Essentially, after a conditional order is triggered, it will become an active order. So, when a conditional order is triggered, cancellation has to be done through the active order endpoint for any unfilled or partially filled active order. As always, orders that have been fully filled cannot be cancelled.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-cancelcond
*/
func (c *InversePerpetualAccountClient) CancelConditionalOrder(ctx context.Context, params *inverseperp.CancelConditionalOrderParams) (response *inverseperp.CancelConditionalOrderResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateInversePerpetualCancelConditionalOrderPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
CancelAllConditionalOrders - cancel all conditional orders. [ /v2/private/stop-order/cancelAll  - POST ]

Cancel all untriggered conditional orders.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-cancelallcond
*/
func (c *InversePerpetualAccountClient) CancelAllConditionalOrders(ctx context.Context, params *inverseperp.CancelAllConditionalOrdersParams) (response *inverseperp.CancelAllConditionalOrdersResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateInversePerpetualCancelAllConditionalOrdersPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
	ReplaceConditionalOrder - replace conditional order. [ /v2/private/stop-order/replace  - POST ]

Replace conditional order can modify/amend your conditional orders.

order_id and symbol are required for identifying a conditional order.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-replacecond
*/
func (c *InversePerpetualAccountClient) ReplaceConditionalOrder(ctx context.Context, params *inverseperp.ReplaceConditionalOrderParams) (response *inverseperp.ReplaceConditionalOrderResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateInversePerpetualReplaceConditionalOrderPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
QueryConditionalOrder - query conditional order. [ /v2/private/stop-order  - GET ]

Query real-time stop order information. When passing the parameter order_id or order_link_id, a single order data will be returned; otherwise, returns up to 10 unfilled orders.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-querycond
*/
func (c *InversePerpetualAccountClient) QueryConditionalOrder(ctx context.Context, params *inverseperp.QueryConditionalOrderParams) (response *inverseperp.QueryConditionalOrderResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualQueryConditionalOrderPath, params, &response)
	if err != nil {
		return
	}
	return
}
