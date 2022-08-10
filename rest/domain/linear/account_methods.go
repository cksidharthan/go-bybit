package linear

import "context"

type AccountInterface interface {
	PlaceActiveOrder(ctx context.Context, params *PlaceActiveOrderParams) (*PlaceActiveOrderResponse, error)
}
