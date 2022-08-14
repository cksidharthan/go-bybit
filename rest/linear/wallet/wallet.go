package wallet

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

func (c *LinearWalletClient) GetWalletBalance(ctx context.Context, params *linear.GetWalletBalanceParams) (res *linear.GetWalletBalanceResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetWalletBalancePath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearWalletClient) GetWalletFundRecords(ctx context.Context, params *linear.GetWalletFundRecordsParams) (res *linear.GetWalletFundRecordsResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetWalletFundRecordsPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearWalletClient) GetWithdrawRecords(ctx context.Context, params *linear.GetWithdrawRecordsParams) (res *linear.GetWithdrawRecordsResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetWithdrawRecordsPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearWalletClient) GetAssetExchangeRecords(ctx context.Context, params *linear.GetAssetExchangeRecordsParams) (res *linear.GetAssetExchangeRecordsResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearAssetExchangeRecordsPath, params, &res)
	if err != nil {
		return
	}
	return
}
