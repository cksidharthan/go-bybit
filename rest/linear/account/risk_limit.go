package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

// GetRiskLimit - get risk limit. [/public/linear/risk-limit - GET]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-getrisklimit
func (c *LinearAccountClient) GetRiskLimit(ctx context.Context, params *linear.GetRiskLimitParams) (res *linear.GetRiskLimitResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicLinearGetRiskLimitPath, params, &res)
	if err != nil {
		return
	}
	return
}

// SetRiskLimit - set risk limit. [/private/linear/risk-limit - POST]
//
// docs - /private/linear/position/set-risk
func (c *LinearAccountClient) SetRiskLimit(ctx context.Context, params *linear.SetRiskLimitParams) (res *linear.SetRiskLimitResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearSetRiskLimitPath, params, &res)
	if err != nil {
		return
	}
	return
}
