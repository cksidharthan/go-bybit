package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

// GetPredictedFundingRate - Get predicted funding rate and my predicted funding fee.. [/private/linear/funding/predicted-funding - GET]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-predictedfunding
func (c *LinearAccountClient) GetPredictedFundingRate(ctx context.Context, params *linear.PredictedFundingRateParams) (res *linear.PredictedFundingRateResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearPredictedFundingRatePath, params, &res)
	if err != nil {
		return
	}
	return
}

// GetLastFundingFee - Get last funding fee. [/private/linear/funding/prev-funding - GET]
//
// Funding settlement occurs every 8 hours at 00:00 UTC, 08:00 UTC and 16:00 UTC. The current interval's fund fee settlement is based on the previous interval's fund rate.
//
// For example, at 16:00, the settlement is based on the fund rate generated at 8:00. The fund rate generated at 16:00 will be used at 0:00 the next day.
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/linear/#t-mylastfundingfee
func (c *LinearAccountClient) GetLastFundingFee(ctx context.Context, params *linear.GetLastFundingFeeParams) (res *linear.GetLastFundingFeeResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetLastFundingFeePath, params, &res)
	if err != nil {
		return
	}
	return
}
