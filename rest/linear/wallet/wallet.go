package wallet

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

// GetWalletBalance - get user's wallet balance. [/v2/private/wallet/balance - GET]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-balance
func (c *LinearWalletClient) GetWalletBalance(ctx context.Context, params *linear.GetWalletBalanceParams) (res *linear.GetWalletBalanceResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetWalletBalancePath, params, &res)
	if err != nil {
		return
	}
	return
}

// GetWalletFundRecords - get user's wallet funding records. [/v2/private/wallet/fund/records - GET]
//
// Get wallet fund records. This endpoint also shows exchanges from the Asset Exchange, where the types for the exchange are ExchangeOrderWithdraw and ExchangeOrderDeposit.
//
// This endpoint returns incomplete information for transfers involving the derivatives wallet. Use the account asset API for creating and querying internal transfers.
//
// Find more detail for types Withdraw and Refund in the Withdraw Records endpoint.
// USDT records will not be returned unless you pass USDT with the coin parameter.
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-walletrecords
func (c *LinearWalletClient) GetWalletFundRecords(ctx context.Context, params *linear.GetWalletFundRecordsParams) (res *linear.GetWalletFundRecordsResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetWalletFundRecordsPath, params, &res)
	if err != nil {
		return
	}
	return
}

// GetWithdrawRecords - get user's withdraw records. [/v2/private/wallet/withdraw/list - GET]
//
// Get withdrawal records.
//
// The difference between data returned by this endpoint and data of type Withdraw in the Wallet Fund Records endpoint:
//
// This endpoint provides one withdrawal operation per record, and you can check the current withdrawal state with the status field.
//
// Once you have submitted a withdrawal application, there will be a record with type Withdraw, and if the application is CancelByUser, Reject or Expire in the Wallet Fund Records endpoint, with a corresponding record with type Refund.
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-withdrawrecords
func (c *LinearWalletClient) GetWithdrawRecords(ctx context.Context, params *linear.GetWithdrawRecordsParams) (res *linear.GetWithdrawRecordsResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearGetWithdrawRecordsPath, params, &res)
	if err != nil {
		return
	}
	return
}

// GetAssetExchangeRecords - get user's asset exchange records. [/v2/private/exchange-order/list - GET]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-assetexchangerecords
func (c *LinearWalletClient) GetAssetExchangeRecords(ctx context.Context, params *linear.GetAssetExchangeRecordsParams) (res *linear.GetAssetExchangeRecordsResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodGet, bybit.PrivateLinearAssetExchangeRecordsPath, params, &res)
	if err != nil {
		return
	}
	return
}
