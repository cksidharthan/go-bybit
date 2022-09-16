package account

import (
	"context"
	"github.com/cksidharthan/go-bybit/bybit"
	inverseperp "github.com/cksidharthan/go-bybit/rest/domain/inverse_perpetual"
	"net/http"
)

/*
GetRiskLimit - get risk limit. [ /v2/private/risk-limit  - GET ]

Get user's risk limit.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-getrisklimit
*/
func (c *InversePerpetualAccountClient) GetRiskLimit(ctx context.Context, params *inverseperp.GetRiskLimitParams) (response *inverseperp.GetRiskLimitResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualGetRiskLimitPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
SetRiskLimit - set risk limit. [ /v2/private/risk-limit  - POST ]

Set user's risk limit.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-setrisklimit
*/
func (c *InversePerpetualAccountClient) SetRiskLimit(ctx context.Context, params *inverseperp.SetRiskLimitParams) (response *inverseperp.SetRiskLimitResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateInversePerpetualSetRiskLimitPath, params, &response)
	if err != nil {
		return
	}
	return
}
