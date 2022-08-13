package account_test

import (
	"context"
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
	linearRest "github.com/cksidharthan/go-bybit/rest/linear"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClient_Linear_Account_Position(t *testing.T) {
	bybitClient := linearRest.NewLinearClient(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	t.Run("Get Positions - LINEAR", func(t *testing.T) {
		response, err := bybitClient.Account().GetPositions(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Get Positions By Symbol - LINEAR", func(t *testing.T) {
		response, err := bybitClient.Account().GetPositionsBySymbol(context.Background(), &linear.GetPositionsBySymbolParams{
			Symbol: "ADAUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Set Auto Add Margin - LINEAR", func(t *testing.T) {
		t.Skip("the api endpoint works. but the test scenario has to be changed")
		response, err := bybitClient.Account().SetAutoAddMargin(context.Background(), &linear.SetAutoAddMarginParams{
			Symbol:        "ADAUSDT",
			Side:          bybit.SideBuy,
			AutoAddMargin: true,
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Switch Margin - LINEAR", func(t *testing.T) {
		t.Skip("the api endpoint works. but the test scenario has to be changed")
		response, err := bybitClient.Account().SwitchMargin(context.Background(), &linear.SwitchMarginParams{
			Symbol:       "ADAUSDT",
			IsIsolated:   true,
			BuyLeverage:  2,
			SellLeverage: 2,
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

}
