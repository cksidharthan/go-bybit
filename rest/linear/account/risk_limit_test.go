package account_test

import (
	"context"
	"os"
	"testing"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
	linearRest "github.com/cksidharthan/go-bybit/rest/linear"
	"github.com/stretchr/testify/assert"
)

func TestClient_Linear_Account_RiskLimit(t *testing.T) {
	bybitClient := linearRest.NewLinearClient(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	t.Run("Get Risk Limit - LINEAR", func(t *testing.T) {
		response, err := bybitClient.Account().GetRiskLimit(context.Background(), &linear.GetRiskLimitParams{
			Symbol: "ADAUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Set Risk Limit - LINEAR", func(t *testing.T) {
		t.Skip("need to create test for correct scenario")
		response, err := bybitClient.Account().SetRiskLimit(context.Background(), &linear.SetRiskLimitParams{
			Symbol: "ADAUSDT",
			Side:   bybit.SideBuy,
			RiskID: 1,
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

}
