package wallet_test

import (
	"context"
	"os"
	"testing"

	"github.com/cksidharthan/go-bybit/bybit"
	domain "github.com/cksidharthan/go-bybit/rest/domain/inverse_perpetual"
	inverseperp "github.com/cksidharthan/go-bybit/rest/inverse_perpetual"
	"github.com/stretchr/testify/assert"
)

func TestClient_InversePerpetualWallet(t *testing.T) {
	t.Parallel()
	inversePerpClient := inverseperp.NewInversePerpetualClient(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	t.Run("Get Wallet Balance - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := inversePerpClient.Wallet().GetWalletBalance(context.Background(), &domain.WalletBalanceParams{
			Coin: "BTC",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Get Wallet Fund Records - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := inversePerpClient.Wallet().GetWalletFundRecords(context.Background(), &domain.WalletFundRecordsParams{})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Get Withdraw Records - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := inversePerpClient.Wallet().GetWithdrawRecords(context.Background(), &domain.WithdrawRecordsParams{})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Get Asset Exchange Records - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := inversePerpClient.Wallet().GetAssetExchangeRecords(context.Background(), &domain.AssetExchangeRecordsParams{})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})
}
