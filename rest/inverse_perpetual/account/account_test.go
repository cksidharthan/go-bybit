package account_test

import (
	"context"
	"os"
	"testing"

	"github.com/cksidharthan/go-bybit/bybit"
	inverseperp "github.com/cksidharthan/go-bybit/rest/domain/inverse_perpetual"
	"github.com/cksidharthan/go-bybit/rest/inverse_perpetual/account"
	"github.com/stretchr/testify/assert"
)

func TestClient_InversePerpetualAccount(t *testing.T) {
	t.Parallel()

	bybitClient := account.NewInversePerpetualAccountClient(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	t.Run("Get API Key Info - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.GetAPIKeyInfo(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})

	t.Run("Get LCP Info - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.GetLCPInfo(context.Background(), &inverseperp.GetLCPInfoParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})
}
