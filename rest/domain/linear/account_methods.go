package linear

import "context"

type AccountInterface interface {
	PlaceActiveOrder(ctx context.Context, params *PlaceActiveOrderParams) (*PlaceActiveOrderResponse, error)
	GetActiveOrder(ctx context.Context, params *GetActiveOrderParams) (*GetActiveOrderResponse, error)
	CancelActiveOrder(ctx context.Context, params *CancelActiveOrderParams) (*CancelActiveOrderResponse, error)
	CancelAllActiveOrders(ctx context.Context, params *CancelAllActiveOrdersParams) (*CancelAllActiveOrdersResponse, error)
	ReplaceActiveOrder(ctx context.Context, params *ReplaceActiveOrderParams) (*ReplaceActiveOrderResponse, error)
	QueryActiveOrder(ctx context.Context, params *QueryActiveOrderParams) (*QueryActiveOrderResponse, error)
}
