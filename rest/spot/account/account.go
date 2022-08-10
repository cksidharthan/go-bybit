package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
	"github.com/google/go-querystring/query"
)

func (c *Client) PlaceActiveOrder(ctx context.Context, params spot.PlaceActiveOrderParams) (res *spot.PlaceActiveOrderResponse, err error) {
	queryString, err := query.Values(params)
	if err != nil {
		return
	}

	if err := c.transporter.SignedPostForm(bybit.PrivateSpotOrderPath, queryString, &res); err != nil {
		return nil, err
	}
	return
}

func (c *Client) GetActiveOrder(ctx context.Context, params spot.GetActiveOrderParams) (res *spot.GetActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateSpotOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) CancelActiveOrder(ctx context.Context, params spot.CancelActiveOrderParams) (res *spot.CancelActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotOrderPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) FastCancelActiveOrder(ctx context.Context, params spot.FastCancelActiveOrderParams) (res *spot.FastCancelActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotOrderFastPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) BatchCancelActiveOrder(ctx context.Context, params spot.BatchCancelActiveOrderParams) (res *spot.BatchCancelActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotOrderBatchCancelPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) BatchFastCancelActiveOrder(ctx context.Context, params spot.BatchFastCancelActiveOrderParams) (res *spot.BatchFastCancelActiveOrderResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotOrderBatchFastCancelPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) BatchCancelActiveOrdersByIDs(ctx context.Context, params spot.BatchCancelActiveOrdersByIDsParams) (res *spot.BatchCancelActiveOrdersByIDsResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotBatchCancelByIdsPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetOpenOrders(ctx context.Context, params spot.GetOpenOrdersParams) (res *spot.GetOpenOrdersResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateOpenOrdersPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetOrderHistory(ctx context.Context, params spot.GetOrderHistoryParams) (res *spot.GetOrderHistoryResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateOrderHistoryPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetTradeHistory(ctx context.Context, params spot.GetTradeHistoryParams) (res *spot.GetTradeHistoryResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateTradeHistoryPath, params, &res)
	if err != nil {
		return
	}
	return
}
