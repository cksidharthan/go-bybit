package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	inverseperp "github.com/cksidharthan/go-bybit/rest/domain/inverse_perpetual"
)

/*
GetPosition - get position. [ /v2/private/position/list  - GET ]

Get my position list.

docs - https://bybit-exchange.github.io/futuresV2/docs/inverse/#t-myposition
*/
func (c *InversePerpetualAccountClient) GetPosition(ctx context.Context) (response *inverseperp.GetPositionResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualGetPositionPath, nil, &response)
	if err != nil {
		return
	}
	return
}

/*
GetPositionWithSymbol - get position with symbol. [ /v2/private/position/list  - GET ]

Get my position list.

docs - https://bybit-exchange.github.io/futuresV2/docs/inverse/#t-myposition
*/
func (c *InversePerpetualAccountClient) GetPositionWithSymbol(ctx context.Context, params *inverseperp.GetPositionWithSymbolParams) (response *inverseperp.GetPositionWithSymbolResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualGetPositionPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
ChangeMargin - change margin. [ /v2/private/position/change-position-margin  - POST ]

Change user's position margin.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-changemargin
*/
func (c *InversePerpetualAccountClient) ChangeMargin(ctx context.Context, params *inverseperp.ChangeMarginParams) (response *inverseperp.ChangeMarginResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateInversePerpetualChangeMarginPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
SetTradingStop - set trading stop. [ /v2/private/position/trading-stop  - POST ]

Set user's position trading-stop.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-tradingstop
*/
func (c *InversePerpetualAccountClient) SetTradingStop(ctx context.Context, params *inverseperp.SetTradingStopParams) (response *inverseperp.SetTradingStopResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateInversePerpetualSetTradingStopPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
SetLeverage - set leverage. [ /v2/private/position/leverage/save  - POST ]

Set user's position leverage.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-setleverage
*/
func (c *InversePerpetualAccountClient) SetLeverage(ctx context.Context, params *inverseperp.SetLeverageParams) (response *inverseperp.SetLeverageResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateInversePerpetualSetLeveragePath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
GetUserTradeRecords - get user trade records. [ /v2/private/execution/list  - GET ]

Get user's trade records.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-usertrades
*/
func (c *InversePerpetualAccountClient) GetUserTradeRecords(ctx context.Context, params *inverseperp.GetUserTradeRecordsParams) (response *inverseperp.GetUserTradeRecordsResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualGetUserTradeRecordsPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
ClosedProfitAndLoss - closed profit and loss. [ /v2/private/position/closed-pnl/list  - GET ]

Get user's closed profit and loss records.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-closedprofitandloss
*/
func (c *InversePerpetualAccountClient) ClosedProfitAndLoss(ctx context.Context, params *inverseperp.ClosedProfitAndLossParams) (response *inverseperp.ClosedProfitAndLossResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualClosedProfitLossPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
PositionTpSlSwitch - position tp sl switch. [ /v2/private/position/tpsl  - POST ]

Switch user's position's tp/sl.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-tpslswitch
*/
func (c *InversePerpetualAccountClient) PositionTpSlSwitch(ctx context.Context, params *inverseperp.PositionTpSlSwitchParams) (response *inverseperp.PositionTpSlSwitchResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateInversePerpetualPositionTpSlSwitchPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
MarginSwitch - margin switch. [ /v2/private/position/margin-switch  - POST ]

Switch user's position's margin mode.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-marginswitch
*/
func (c *InversePerpetualAccountClient) MarginSwitch(ctx context.Context, params *inverseperp.MarginSwitchParams) (response *inverseperp.MarginSwitchResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateInversePerpetualMarginSwitchPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
QueryTradingFeeRate - query trading fee rate. [ /v2/private/execution/fee  - GET ]

Query user's trading fee rate.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-querytradingfeerate
*/
func (c *InversePerpetualAccountClient) QueryTradingFeeRate(ctx context.Context, params *inverseperp.QueryTradingFeeRateParams) (response *inverseperp.QueryTradingFeeRateResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualQueryTradingFeeRatePath, params, &response)
	if err != nil {
		return
	}
	return
}
