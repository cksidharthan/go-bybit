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

	t.Run("Switch Position Mode - LINEAR", func(t *testing.T) {
		t.Skip("the api endpoint works. but the test scenario has to be changed")
		response, err := bybitClient.Account().SwitchPositionMode(context.Background(), &linear.SwitchPositionModeParams{
			Symbol: "ADAUSDT",
			Mode:   "BothSide",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Position TP SL Switch - LINEAR", func(t *testing.T) {
		t.Skip("the api endpoint works. but the test scenario has to be changed")
		response, err := bybitClient.Account().PositionTpSlSwitch(context.Background(), &linear.PositionTpSlSwitchParams{
			Symbol:   "ADAUSDT",
			TpSlMode: "Partial",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Add Reduce Margin - LINEAR", func(t *testing.T) {
		t.Skip("the api endpoint works. but the test scenario has to be changed")
		response, err := bybitClient.Account().AddReduceMargin(context.Background(), &linear.AddReduceMarginParams{
			Symbol: "ADAUSDT",
			Side:   bybit.SideBuy,
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Set Leverage - LINEAR", func(t *testing.T) {
		t.Skip("the api endpoint works. but the test scenario has to be changed")
		response, err := bybitClient.Account().SetLeverage(context.Background(), &linear.SetLeverageParams{
			Symbol: "ADAUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Set Trading stop - LINEAR", func(t *testing.T) {
		t.Skip("the api endpoint works. but the test scenario has to be changed")
		response, err := bybitClient.Account().SetTradingStop(context.Background(), &linear.SetTradingStopParams{
			Symbol: "ADAUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Get User Trade Records - LINEAR", func(t *testing.T) {
		response, err := bybitClient.Account().GetUserTradeRecords(context.Background(), &linear.GetUserTradeRecordsParams{
			Symbol: "ADAUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Get Extended User Trade Records - LINEAR", func(t *testing.T) {
		response, err := bybitClient.Account().GetExtendedUserTradeRecords(context.Background(), &linear.GetExtendedUserTradeRecordsParams{
			Symbol: "ADAUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Get Closed Profit and Loss - LINEAR", func(t *testing.T) {
		response, err := bybitClient.Account().GetClosedProfitLoss(context.Background(), &linear.GetClosedProfitLossParams{
			Symbol: "ADAUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

}
