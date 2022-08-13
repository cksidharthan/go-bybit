package account

import (
	"context"
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
	"net/http"
)

func (c *LinearAccountClient) GetPositionsBySymbol(ctx context.Context, params *linear.GetPositionsBySymbolParams) (res *linear.GetPositionsBySymbolResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetPositionPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) GetPositions(ctx context.Context) (res *linear.GetPositionsResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetPositionPath, nil, &res)
	if err != nil {
		return
	}
	return
}
