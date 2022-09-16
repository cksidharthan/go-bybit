package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	inverseperp "github.com/cksidharthan/go-bybit/rest/domain/inverse_perpetual"
)

/*
GetAPIKeyInfo - get api key info. [ /v2/private/account/api-key  - GET ]

Get user's api key info.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-getapikeyinfo
*/
func (c *InversePerpetualAccountClient) GetAPIKeyInfo(ctx context.Context) (response *inverseperp.GetAPIKeyInfoResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualGetAPIKeyInfoPath, nil, &response)
	if err != nil {
		return
	}
	return
}

/*
GetLCPInfo - get lcp info. [ /v2/private/account/lcp  - GET ]

Get user's lcp info.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-getlcpinfo
*/
func (c *InversePerpetualAccountClient) GetLCPInfo(ctx context.Context, params *inverseperp.GetLCPInfoParams) (response *inverseperp.GetLCPInfoResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualGetLCPInfoPath, params, &response)
	if err != nil {
		return
	}
	return
}
