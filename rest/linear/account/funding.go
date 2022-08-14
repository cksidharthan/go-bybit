package account

import (
	"context"
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
	"net/http"
)

func (c *LinearAccountClient) GetPredictedFundingRate(ctx context.Context, params *linear.PredictedFundingRateParams) (res *linear.PredictedFundingRateResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearPredictedFundingRatePath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearAccountClient) GetLastFundingFee(ctx context.Context, params *linear.GetLastFundingFeeParams) (res *linear.GetLastFundingFeeResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetLastFundingFeePath, params, &res)
	if err != nil {
		return
	}
	return
}
