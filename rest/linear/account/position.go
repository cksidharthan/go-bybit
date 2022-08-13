package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
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

func (c *LinearAccountClient) SetAutoAddMargin(ctx context.Context, params *linear.SetAutoAddMarginParams) (res *linear.SetAutoAddMarginResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearSetAutoAddMarginPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) SwitchMargin(ctx context.Context, params *linear.SwitchMarginParams) (res *linear.SwitchMarginResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateSwithcMarginPath, params, &res)
	if err != nil {
		return
	}
	return
}
