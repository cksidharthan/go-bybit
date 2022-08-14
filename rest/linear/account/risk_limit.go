package account

import (
	"context"
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
	"net/http"
)

func (c *LinearAccountClient) GetRiskLimit(ctx context.Context, params *linear.GetRiskLimitParams) (res *linear.GetRiskLimitResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicLinearGetRiskLimitPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) SetRiskLimit(ctx context.Context, params *linear.SetRiskLimitParams) (res *linear.SetRiskLimitResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodPost, bybit.PrivateLinearSetRiskLimitPath, params, &res)
	if err != nil {
		return
	}
	return
}
