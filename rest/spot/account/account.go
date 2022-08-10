package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
	"github.com/google/go-querystring/query"
)

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

func (c *SpotAccountClient) GetActiveOrder(ctx context.Context, params spot.GetActiveOrderParams) (res *spot.GetActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateSpotOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *SpotAccountClient) CancelActiveOrder(ctx context.Context, params spot.CancelActiveOrderParams) (res *spot.CancelActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *SpotAccountClient) FastCancelActiveOrder(ctx context.Context, params spot.FastCancelActiveOrderParams) (res *spot.FastCancelActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotOrderFastPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *SpotAccountClient) BatchCancelActiveOrder(ctx context.Context, params spot.BatchCancelActiveOrderParams) (res *spot.BatchCancelActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotOrderBatchCancelPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *SpotAccountClient) BatchFastCancelActiveOrder(ctx context.Context, params spot.BatchFastCancelActiveOrderParams) (res *spot.BatchFastCancelActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotOrderBatchFastCancelPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *SpotAccountClient) BatchCancelActiveOrdersByIDs(ctx context.Context, params spot.BatchCancelActiveOrdersByIDsParams) (res *spot.BatchCancelActiveOrdersByIDsResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotBatchCancelByIdsPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *SpotAccountClient) GetOpenOrders(ctx context.Context, params spot.GetOpenOrdersParams) (res *spot.GetOpenOrdersResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateOpenOrdersPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *SpotAccountClient) GetOrderHistory(ctx context.Context, params spot.GetOrderHistoryParams) (res *spot.GetOrderHistoryResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateOrderHistoryPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *SpotAccountClient) GetTradeHistory(ctx context.Context, params spot.GetTradeHistoryParams) (res *spot.GetTradeHistoryResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateTradeHistoryPath, params, &res)
	if err != nil {
		return
	}
	return
}
