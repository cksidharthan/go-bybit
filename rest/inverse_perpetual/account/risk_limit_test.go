package account

import (
	"context"
	"os"
	"testing"

	"github.com/cksidharthan/go-bybit/bybit"
	inverseperp "github.com/cksidharthan/go-bybit/rest/domain/inverse_perpetual"
	"github.com/stretchr/testify/assert"
)

func TestClient_InversePerpetual_Account_RiskLimit(t *testing.T) {
	t.Parallel()

	bybitClient := NewInversePerpetualAccountClient(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	t.Run("Get Risk limit - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.GetRiskLimit(context.Background(), &inverseperp.GetRiskLimitParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSD", response.Result[0].Symbol)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})

	t.Run("Set Risk limit - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		t.Skip("Skipping this test for now. It works but this test scenario has to be changed")
		response, err := bybitClient.SetRiskLimit(context.Background(), &inverseperp.SetRiskLimitParams{
			Symbol: "BTCUSD",
			RiskID: 1,
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})
}
