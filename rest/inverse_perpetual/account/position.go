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
