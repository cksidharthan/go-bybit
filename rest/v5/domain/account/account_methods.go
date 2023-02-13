package account

import "context"

type Client interface {
	GetWalletBalance(ctx context.Context, params *GetWalletBalanceParams) (*GetWalletBalanceResponse, error)
}
