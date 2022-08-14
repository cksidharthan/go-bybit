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
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearSwitchMarginPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) SwitchPositionMode(ctx context.Context, params *linear.SwitchPositionModeParams) (res *linear.SwitchPositionModeResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearSwitchPositionModePath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) PositionTpSlSwitch(ctx context.Context, params *linear.PositionTpSlSwitchParams) (res *linear.PositionTpSlSwitchResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearPositionTpSlSwitch, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) AddReduceMargin(ctx context.Context, params *linear.AddReduceMarginParams) (res *linear.AddReduceMarginResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearAddReduceMarginPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) SetLeverage(ctx context.Context, params *linear.SetLeverageParams) (res *linear.SetLeverageResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearSetLeveragePath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) SetTradingStop(ctx context.Context, params *linear.SetTradingStopParams) (res *linear.SetTradingStopResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearSetTradingStopPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) GetUserTradeRecords(ctx context.Context, params *linear.GetUserTradeRecordsParams) (res *linear.GetUserTradeRecordsResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetUserTradeRecordsPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) GetExtendedUserTradeRecords(ctx context.Context, params *linear.GetExtendedUserTradeRecordsParams) (res *linear.GetExtendedUserTradeRecordsResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetExtendedUserTradeRecordsPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) GetClosedProfitLoss(ctx context.Context, params *linear.GetClosedProfitLossParams) (res *linear.GetClosedProfitLossResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetClosedProfitLossPath, params, &res)
	if err != nil {
		return
	}
	return
}
