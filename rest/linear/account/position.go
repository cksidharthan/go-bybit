package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

// GetPositionsBySymbol - get my positions by symbol. [/private/linear/position/list - GET]
//
// docs  - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-myposition
func (c *LinearAccountClient) GetPositionsBySymbol(ctx context.Context, params *linear.GetPositionsBySymbolParams) (res *linear.GetPositionsBySymbolResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetPositionPath, params, &res)
	if err != nil {
		return
	}
	return
}

// GetPositions - get my positions. [/private/linear/position/list - GET]
//
// docs  - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-myposition
func (c *LinearAccountClient) GetPositions(ctx context.Context) (res *linear.GetPositionsResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetPositionPath, nil, &res)
	if err != nil {
		return
	}
	return
}

// SetAutoAddMargin - Set auto add margin, or Auto-Margin Replenishment. [/private/linear/position/set-auto-add-margin - POST]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-setautoaddmargin
func (c *LinearAccountClient) SetAutoAddMargin(ctx context.Context, params *linear.SetAutoAddMarginParams) (res *linear.SetAutoAddMarginResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearSetAutoAddMarginPath, params, &res)
	if err != nil {
		return
	}
	return
}

// SwitchMargin - Switch Cross/Isolated; must set leverage value when switching from Cross to Isolated. [/private/linear/position/switch-isolated - POST]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-marginswitch
func (c *LinearAccountClient) SwitchMargin(ctx context.Context, params *linear.SwitchMarginParams) (res *linear.SwitchMarginResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearSwitchMarginPath, params, &res)
	if err != nil {
		return
	}
	return
}

// SwitchPositionMode -switched position mode between isolated and cross. [//private/linear/position/switch-mode - POST]
//
// If you are in One-Way Mode, you can only open one position on Buy or Sell side. If you are in Hedge Mode, you can open both Buy and Sell side positions simultaneously.
// Supports switching between One-Way Mode and Hedge Mode at the coin level.
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-switchpositionmode
func (c *LinearAccountClient) SwitchPositionMode(ctx context.Context, params *linear.SwitchPositionModeParams) (res *linear.SwitchPositionModeResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearSwitchPositionModePath, params, &res)
	if err != nil {
		return
	}
	return
}

// PositionTpSlSwitch - Switch mode between Full or Partial. [/private/linear/tpsl/switch-mode - POST]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-switchmode
func (c *LinearAccountClient) PositionTpSlSwitch(ctx context.Context, params *linear.PositionTpSlSwitchParams) (res *linear.PositionTpSlSwitchResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearPositionTpSlSwitch, params, &res)
	if err != nil {
		return
	}
	return
}

// AddReduceMargin - Add or Reduce Margin. [/private/linear/position/add-margin - POST]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-addmargin
func (c *LinearAccountClient) AddReduceMargin(ctx context.Context, params *linear.AddReduceMarginParams) (res *linear.AddReduceMarginResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearAddReduceMarginPath, params, &res)
	if err != nil {
		return
	}
	return
}

// SetLeverage - Set leverage. [/private/linear/position/set-leverage - POST]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-setleverage
func (c *LinearAccountClient) SetLeverage(ctx context.Context, params *linear.SetLeverageParams) (res *linear.SetLeverageResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearSetLeveragePath, params, &res)
	if err != nil {
		return
	}
	return
}

// SetTradingStop - Set Trading Stop. [/private/linear/position/trading-stop - POST]
//
// Set take profit, stop loss, and trailing stop for your open position. If using partial mode, TP/SL/TS orders will not close your entire position.
//
// docs - /private/linear/position/trading-stop
func (c *LinearAccountClient) SetTradingStop(ctx context.Context, params *linear.SetTradingStopParams) (res *linear.SetTradingStopResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearSetTradingStopPath, params, &res)
	if err != nil {
		return
	}
	return
}

// GetUserTradeRecords - Get user trade records. [/private/linear/trade/execution/list - GET]
//
// Get user's trading records. The results are ordered in descending order (the first item is the latest). Returns records up to 2 years old.
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-usertraderecords
func (c *LinearAccountClient) GetUserTradeRecords(ctx context.Context, params *linear.GetUserTradeRecordsParams) (res *linear.GetUserTradeRecordsResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetUserTradeRecordsPath, params, &res)
	if err != nil {
		return
	}
	return
}

// GetExtendedUserTradeRecords - Get extended user trade records. [/private/linear/trade/execution/list-extended - GET]
//
// Get user's trading records. The results are ordered in descending order (the first item is the latest). Return records up to 2 years old.
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-userhistorytraderecords
func (c *LinearAccountClient) GetExtendedUserTradeRecords(ctx context.Context, params *linear.GetExtendedUserTradeRecordsParams) (res *linear.GetExtendedUserTradeRecordsResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetExtendedUserTradeRecordsPath, params, &res)
	if err != nil {
		return
	}
	return
}

// GetClosedProfitLoss - Get closed profit loss. [/private/linear/trade/closed-profit-loss - GET]
//
// Get user's closed profit and loss records. The results are ordered in descending order (the first item is the latest).
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-closedprofitandloss
func (c *LinearAccountClient) GetClosedProfitLoss(ctx context.Context, params *linear.GetClosedProfitLossParams) (res *linear.GetClosedProfitLossResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetClosedProfitLossPath, params, &res)
	if err != nil {
		return
	}
	return
}
