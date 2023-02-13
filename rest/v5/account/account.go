package account

import (
	"context"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/v5/domain/account"
)

func (c Client) GetWalletBalance(ctx context.Context, params *account.GetWalletBalanceParams) (*account.GetWalletBalanceResponse, error) {
	var resp account.GetWalletBalanceResponse
	err := c.http.Do(ctx, "GET", bybit.GetWalletBalance, params, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
