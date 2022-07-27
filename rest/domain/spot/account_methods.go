package spot

import (
	"context"
)

type AccountInterface interface {
	PlaceActiveOrder(ctx context.Context, order ActiveOrderParams) (*ActiveOrderResponse, error)
}
