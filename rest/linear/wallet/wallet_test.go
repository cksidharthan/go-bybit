package wallet_test

import (
	"context"
	"os"
	"testing"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
	linearRest "github.com/cksidharthan/go-bybit/rest/linear"
	"github.com/stretchr/testify/assert"
)

func TestClient_Linear_Wallet(t *testing.T) {
	t.Parallel()
	bybitClient := linearRest.NewLinearClient(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	t.Run("Get Wallet Balance - LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Wallet().GetWalletBalance(context.Background(), &linear.GetWalletBalanceParams{
			Coin: "BTC",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)

	})

	t.Run("Get Wallet Fund Records - LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Wallet().GetWalletFundRecords(context.Background(), &linear.GetWalletFundRecordsParams{})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Get Withdraw Records - LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Wallet().GetWithdrawRecords(context.Background(), &linear.GetWithdrawRecordsParams{})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Get Asset Exchange Records - LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Wallet().GetAssetExchangeRecords(context.Background(), &linear.GetAssetExchangeRecordsParams{
			Limit:     5,
			From:      0,
			Direction: "",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

}
