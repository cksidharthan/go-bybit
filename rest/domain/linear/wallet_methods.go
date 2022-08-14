package linear

import "context"

type WalletInterface interface {
	GetWalletBalance(ctx context.Context, params *GetWalletBalanceParams) (*GetWalletBalanceResponse, error)
	GetWalletFundRecords(ctx context.Context, params *GetWalletFundRecordsParams) (*GetWalletFundRecordsResponse, error)
	GetWithdrawRecords(ctx context.Context, params *GetWithdrawRecordsParams) (*GetWithdrawRecordsResponse, error)
	GetAssetExchangeRecords(ctx context.Context, params *GetAssetExchangeRecordsParams) (*GetAssetExchangeRecordsResponse, error)
}
