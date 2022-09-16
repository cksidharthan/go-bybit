package account

import (
	"context"
	"github.com/cksidharthan/go-bybit/bybit"
	inverseperp "github.com/cksidharthan/go-bybit/rest/domain/inverse_perpetual"
	"net/http"
)

/*
GetFundingRate - get funding rate. [ /v2/private/funding/prev-funding-rate  - GET ]

Get user's funding rate.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-getfundingrate
*/
func (c *InversePerpetualAccountClient) GetFundingRate(ctx context.Context, params *inverseperp.GetFundingRateParams) (response *inverseperp.GetFundingRateResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualGetLastFundingRatePath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
GetLastFundingFee - get last funding fee. [ /v2/private/funding/prev-funding  - GET ]

Get user's last funding fee.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-getlastfundingfee
*/
func (c *InversePerpetualAccountClient) GetLastFundingFee(ctx context.Context, params *inverseperp.GetLastFundingFeeParams) (response *inverseperp.GetLastFundingFeeResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualGetLastFundingFeePath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
GetPredictedFunding - get predicted funding. [ /v2/private/funding/predicted-funding  - GET ]

Get user's predicted funding.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-getpredictedfunding
*/
func (c *InversePerpetualAccountClient) GetPredictedFunding(ctx context.Context, params *inverseperp.GetPredictedFundingParams) (response *inverseperp.GetPredictedFundingResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualGetPredictedFundingPath, params, &response)
	if err != nil {
		return
	}
	return
}
