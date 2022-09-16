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

docs - https://bybit-exchange.github.io/docs/inverse/#t-myposition
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

docs - https://bybit-exchange.github.io/docs/inverse/#t-myposition
*/
func (c *InversePerpetualAccountClient) GetPositionWithSymbol(ctx context.Context, params *inverseperp.GetPositionWithSymbolParams) (response *inverseperp.GetPositionWithSymbolResponse, err error) {
	err = c.Transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualGetPositionPath, params, &response)
	if err != nil {
		return
	}
	return
}
