package spot

import "context"

type WalletInterface interface {
	GetWalletBalance(ctx context.Context) (*GetWalletBalanceResponse, error)
}
