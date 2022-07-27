package spot

import (
	"context"
)

type AccountInterface interface {
	PlaceActiveOrder(ctx context.Context, params PlaceActiveOrderParams) (*PlaceActiveOrderResponse, error)
	GetActiveOrder(ctx context.Context, params GetActiveOrderParams) (*GetActiveOrderResponse, error)
}
