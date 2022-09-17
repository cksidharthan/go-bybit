package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	inverseperp "github.com/cksidharthan/go-bybit/rest/domain/inverse_perpetual"
)

/*
PlaceActiveOrder - place active order. [ /v2/private/order/create  - POST ]

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-placeactive
*/
func (c *InversePerpetualAccountClient) PlaceActiveOrder(ctx context.Context, params *inverseperp.PlaceActiveOrderParams) (response *inverseperp.PlaceActiveOrderResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateInversePerpetualPlaceActiveOrderPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
GetActiveOrder - get active order. [ /v2/private/order/list  - GET ]

Because order creation/cancellation is asynchronous, there can be a data delay in this endpoint. You can get real-time order info with the Query Active Order (real-time) endpoint.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-getactive
*/
func (c *InversePerpetualAccountClient) GetActiveOrder(ctx context.Context, params *inverseperp.GetActiveOrderParams) (response *inverseperp.GetActiveOrderResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualGetActiveOrderPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
CancelActiveOrder - cancel active order. [ /v2/private/order/cancel  - POST ]

Either order_id or order_link_id are required for cancelling active orders. order_id - this unique 36 characters order ID was returned to you when the active order was created successfully.

You may cancel active orders that are unfilled or partially filled. Fully filled orders cannot be cancelled.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-cancelactive
*/
func (c *InversePerpetualAccountClient) CancelActiveOrder(ctx context.Context, params *inverseperp.CancelActiveOrderParams) (response *inverseperp.CancelActiveOrderResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateInversePerpetualCancelActiveOrderPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
CancelAllActiveOrders - cancel all active orders. [ /v2/private/order/cancelAll  - POST ]

Cancel all active orders that are unfilled or partially filled. Fully filled orders cannot be cancelled.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-cancelallactive
*/
func (c *InversePerpetualAccountClient) CancelAllActiveOrders(ctx context.Context, params *inverseperp.CancelAllActiveOrdersParams) (response *inverseperp.CancelAllActiveOrdersResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateInversePerpetualCancelAllActiveOrdersPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
ReplaceActiveOrder - replace active order. [ /v2/private/order/replace  - POST ]

Replace active orders that are unfilled or partially filled. Fully filled orders cannot be replaced.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-replaceactive
*/
func (c *InversePerpetualAccountClient) ReplaceActiveOrder(ctx context.Context, params *inverseperp.ReplaceActiveOrderParams) (response *inverseperp.ReplaceActiveOrderResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateInversePerpetualReplaceActiveOrderPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
QueryActiveOrder - query active order. [ /v2/private/order  - GET ]

Query real-time active order information. If only order_id or order_link_id are passed, a single order will be returned; otherwise, returns up to 500 unfilled orders.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-queryactive
*/
func (c *InversePerpetualAccountClient) QueryActiveOrder(ctx context.Context, params *inverseperp.QueryActiveOrderParams) (response *inverseperp.QueryActiveOrderResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualQueryActiveOrderPath, params, &response)
	if err != nil {
		return
	}
	return
}
