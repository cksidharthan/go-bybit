package methods

import (
	"context"

	"github.com/cksidharthan/go-bybit/rest/domain/spot/account/types"
)

type Interface interface {
	PlaceActiveOrder(ctx context.Context, order types.SpotActiveOrderParams) (*types.SpotActiveOrderResponse, error)
}
