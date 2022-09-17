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

func TestClient_InversePerpetualAccount_Position(t *testing.T) {
	t.Parallel()

	bybitClient := account.NewInversePerpetualAccountClient(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	t.Run("Get Position - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.GetPosition(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})

	t.Run("Get Position With Symbol - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.GetPositionWithSymbol(context.Background(), &inverseperp.GetPositionWithSymbolParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})

	// TODO: Fix this skipped test
	t.Run("Change Position Margin - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		t.Skip("Skipping this test as it is not possible to test this without actually changing the margin")
		response, err := bybitClient.ChangeMargin(context.Background(), &inverseperp.ChangeMarginParams{
			Symbol: "BTCUSD",
			Margin: "12",
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})

	// TODO: Fix this skipped test
	t.Run("Set Trading Stop - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		t.Skip("Skipping this test as there is no position in USD and USD asset is not working in my test account")
		response, err := bybitClient.SetTradingStop(context.Background(), &inverseperp.SetTradingStopParams{
			Symbol:   "BTCUSD",
			StopLoss: 100,
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})

	// TODO: Fix this skipped test
	t.Run("Set Leverage - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		t.Skip("Need to check this endpoint")
		response, err := bybitClient.SetLeverage(context.Background(), &inverseperp.SetLeverageParams{
			Symbol:       "BTCUSD",
			Leverage:     1,
			LeverageOnly: true,
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})

	t.Run("Get User Trade Records - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.GetUserTradeRecords(context.Background(), &inverseperp.GetUserTradeRecordsParams{
			Symbol:  "BTCUSD",
			OrderID: "123",
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})

	t.Run("Closed Profit and Loss - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.ClosedProfitAndLoss(context.Background(), &inverseperp.ClosedProfitAndLossParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})

	// TODO: Fix this skipped test
	t.Run("Position TP SL Switch - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		t.Skip("Skipping this test as there is no position in USD and USD asset is not working in my test account")
		response, err := bybitClient.PositionTpSlSwitch(context.Background(), &inverseperp.PositionTpSlSwitchParams{
			Symbol:   "BTCUSD",
			TpSlMode: "Full",
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})

	// TODO: Fix this skipped test
	t.Run("Margin Switch - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		t.Skip("Skipping this test as there is no position in USD and USD asset is not working in my test account")
		response, err := bybitClient.MarginSwitch(context.Background(), &inverseperp.MarginSwitchParams{
			Symbol:       "BTCUSD",
			IsIsolated:   true,
			BuyLeverage:  10,
			SellLeverage: 10,
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})

	t.Run("Query Trading Fee Rate - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.QueryTradingFeeRate(context.Background(), &inverseperp.QueryTradingFeeRateParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})
}
