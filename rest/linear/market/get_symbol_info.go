package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

func (c *Client) GetSymbolInformation(ctx context.Context, params *linear.GetSymbolInformationParams) (res *linear.GetSymbolInformationResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicGetSymbolInformationPath, params, &res)
	if err != nil {
		return
	}
	return
}
