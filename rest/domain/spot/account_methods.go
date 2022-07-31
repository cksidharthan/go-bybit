package spot

import (
	"context"
)

type AccountInterface interface {
	PlaceActiveOrder(ctx context.Context, params PlaceActiveOrderParams) (*PlaceActiveOrderResponse, error)
	GetActiveOrder(ctx context.Context, params GetActiveOrderParams) (*GetActiveOrderResponse, error)
	CancelActiveOrder(ctx context.Context, params CancelActiveOrderParams) (*CancelActiveOrderResponse, error)
	FastCancelActiveOrder(ctx context.Context, params FastCancelActiveOrderParams) (*FastCancelActiveOrderResponse, error)
	BatchCancelActiveOrder(ctx context.Context, params BatchCancelActiveOrderParams) (*BatchCancelActiveOrderResponse, error)
	BatchFastCancelActiveOrder(ctx context.Context, params BatchFastCancelActiveOrderParams) (*BatchFastCancelActiveOrderResponse, error)
	BatchCancelActiveOrdersByIDs(ctx context.Context, params BatchCancelActiveOrdersByIDsParams) (*BatchCancelActiveOrdersByIDsResponse, error)
	GetOpenOrders(ctx context.Context, params GetOpenOrdersParams) (*GetOpenOrdersResponse, error)
	GetOrderHistory(ctx context.Context, params GetOrderHistoryParams) (*GetOrderHistoryResponse, error)
	GetTradeHistory(ctx context.Context, params GetTradeHistoryParams) (*GetTradeHistoryResponse, error)
}
