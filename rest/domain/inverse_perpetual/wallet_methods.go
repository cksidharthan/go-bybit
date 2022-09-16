package inverseperp

import "context"

type WalletInterface interface {
	GetWalletBalance(ctx context.Context, params *WalletBalanceParams) (*WalletBalanceResponse, error)
	GetWalletFundRecords(ctx context.Context, params *WalletFundRecordsParams) (*WalletFundRecordsResponse, error)
	GetWithdrawRecords(ctx context.Context, params *WithdrawRecordsParams) (*WithdrawRecordsResponse, error)
	GetAssetExchangeRecords(ctx context.Context, params *AssetExchangeRecordsParams) (*AssetExchangeRecordsResponse, error)
}
