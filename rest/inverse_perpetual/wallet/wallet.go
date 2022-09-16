package wallet

import (
	"context"
	"github.com/cksidharthan/go-bybit/bybit"
	inverseperp "github.com/cksidharthan/go-bybit/rest/domain/inverse_perpetual"
	"net/http"
)

/*
GetWalletBalance - get wallet balance. [ /v2/private/wallet/balance  - GET ]

# Get wallet balance info

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-balance
*/
func (c *InversePerpetualWalletClient) GetWalletBalance(ctx context.Context, params *inverseperp.WalletBalanceParams) (response *inverseperp.WalletBalanceResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualWalletBalancePath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
GetWalletFundRecords - get wallet fund records. [ /v2/private/wallet/fund/records  - GET ]

Get wallet fund records. This endpoint also shows exchanges from the Asset Exchange, where the types for the exchange are ExchangeOrderWithdraw and ExchangeOrderDeposit.

This endpoint returns incomplete information for transfers involving the derivatives wallet. Use the account asset API for creating and querying internal transfers.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-walletrecords
*/
func (c *InversePerpetualWalletClient) GetWalletFundRecords(ctx context.Context, params *inverseperp.WalletFundRecordsParams) (response *inverseperp.WalletFundRecordsResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualWalletFundRecordsPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
GetWalletWithdrawRecords - get wallet withdraw records. [ /v2/private/wallet/withdraw/list  - GET ]

Get withdrawal records.

The difference between data returned by this endpoint and data of type Withdraw in the Wallet Fund Records endpoint:

This endpoint provides one withdrawal operation per record, and you can check the current withdrawal state with the status field.

Once you have submitted a withdrawal application, there will be a record with type Withdraw, and if the application is CancelByUser, Reject or Expire in the Wallet Fund Records endpoint, with a corresponding record with type Refund.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-withdrawrecords
*/
func (c *InversePerpetualWalletClient) GetWalletWithdrawRecords(ctx context.Context, params *inverseperp.WithdrawRecordsParams) (response *inverseperp.WithdrawRecordsResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualWithdrawRecordsPath, params, &response)
	if err != nil {
		return
	}
	return
}

/*
GetAssetExchangeRecords - get asset exchange records. [ /v2/private/asset/records  - GET ]

Get asset exchange records.

This endpoint returns the records of asset exchanges, including the exchanges from the Wallet Fund Records endpoint.

docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-assetexchangerecords
*/
func (c *InversePerpetualWalletClient) GetAssetExchangeRecords(ctx context.Context, params *inverseperp.AssetExchangeRecordsParams) (response *inverseperp.AssetExchangeRecordsResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PrivateInversePerpetualAssetExchangeRecords, params, &response)
	if err != nil {
		return
	}
	return
}
