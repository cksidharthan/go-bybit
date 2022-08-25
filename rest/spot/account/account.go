package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
	"github.com/google/go-querystring/query"
)

// PlaceActiveOrder - place an active order. [/spot/v1/order]
func (c *SpotAccountClient) PlaceActiveOrder(ctx context.Context, params spot.PlaceActiveOrderParams) (res *spot.PlaceActiveOrderResponse, err error) {
	queryString, err := query.Values(params)
	if err != nil {
		return
	}

	if err := c.transporter.SignedPostForm(bybit.PrivateSpotOrderPath, queryString, &res); err != nil {
		return nil, err
	}
	return
}

// GetActiveOrder - get a list of all active order. [/spot/v1/order]
func (c *SpotAccountClient) GetActiveOrder(ctx context.Context, params spot.GetActiveOrderParams) (res *spot.GetActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateSpotOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

// CancelActiveOrder - cancel an active order. [/spot/v1/order]
func (c *SpotAccountClient) CancelActiveOrder(ctx context.Context, params spot.CancelActiveOrderParams) (res *spot.CancelActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

// FastCancelActiveOrder - cancel an active order quickly. [/spot/v1/order/fast]
func (c *SpotAccountClient) FastCancelActiveOrder(ctx context.Context, params spot.FastCancelActiveOrderParams) (res *spot.FastCancelActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotOrderFastPath, params, &res)
	if err != nil {
		return
	}
	return
}

// BatchCancelActiveOrder - cancel multiple active orders. [/spot/order/batch-cancel]
func (c *SpotAccountClient) BatchCancelActiveOrder(ctx context.Context, params spot.BatchCancelActiveOrderParams) (res *spot.BatchCancelActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotOrderBatchCancelPath, params, &res)
	if err != nil {
		return
	}
	return
}

// BatchFastCancelActiveOrder - cancel multiple active orders quickly. [/spot/order/batch-fast-cancel]
func (c *SpotAccountClient) BatchFastCancelActiveOrder(ctx context.Context, params spot.BatchFastCancelActiveOrderParams) (res *spot.BatchFastCancelActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotOrderBatchFastCancelPath, params, &res)
	if err != nil {
		return
	}
	return
}

// BatchCancelActiveOrdersByIDs - cancel multiple active orders by IDs. [/spot/order/batch-cancel-by-ids]
func (c *SpotAccountClient) BatchCancelActiveOrdersByIDs(ctx context.Context, params spot.BatchCancelActiveOrdersByIDsParams) (res *spot.BatchCancelActiveOrdersByIDsResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotBatchCancelByIdsPath, params, &res)
	if err != nil {
		return
	}
	return
}

// GetOpenOrders - get a list of all open orders. [/spot/v1/open-orders]
func (c *SpotAccountClient) GetOpenOrders(ctx context.Context, params spot.GetOpenOrdersParams) (res *spot.GetOpenOrdersResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateOpenOrdersPath, params, &res)
	if err != nil {
		return
	}
	return
}

// GetOrderHistory - get a list of all order history. [/spot/v1/history-orders]
func (c *SpotAccountClient) GetOrderHistory(ctx context.Context, params spot.GetOrderHistoryParams) (res *spot.GetOrderHistoryResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateOrderHistoryPath, params, &res)
	if err != nil {
		return
	}
	return
}

// GetTradeHistory - get a list of all trade history. [/spot/v1/myTrades]
func (c *SpotAccountClient) GetTradeHistory(ctx context.Context, params spot.GetTradeHistoryParams) (res *spot.GetTradeHistoryResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateTradeHistoryPath, params, &res)
	if err != nil {
		return
	}
	return
}
