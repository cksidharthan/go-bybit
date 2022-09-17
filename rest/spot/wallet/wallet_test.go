package wallet_test

import (
	"context"
	"os"
	"testing"

	"github.com/cksidharthan/go-bybit/bybit"
	spotRest "github.com/cksidharthan/go-bybit/rest/spot"
	"github.com/stretchr/testify/assert"
)

func TestSpotClient_Wallet(t *testing.T) {
	t.Parallel()
	client := spotRest.NewSpotClient(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	t.Run("happy path - get symbols", func(t *testing.T) {
		t.Parallel()
		balance, err := client.Wallet().GetWalletBalance(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 0, balance.RetCode)
		assert.Equal(t, "OK", balance.RetMsg)
	})
}
