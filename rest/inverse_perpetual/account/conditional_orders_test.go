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

func TestClient_InversePerpetualConditionalOrders(t *testing.T) {
	t.Parallel()

	bybitClient := account.NewInversePerpetualAccountClient(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	t.Run("Get Conditional Orders - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.GetConditionalOrder(context.Background(), &inverseperp.GetConditionalOrderParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})

	// TODO: Fix this skipped test
	t.Run("Place Conditional Order - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		t.Skip("Skipping this test as it places an actual order and my test account doesn't have enough USD balance")
		response, err := bybitClient.PlaceConditionalOrder(context.Background(), &inverseperp.PlaceConditionalOrderParams{
			Symbol:         "BTCUSD",
			OrderType:      "Market",
			Side:           "Buy",
			Qty:            0.001,
			TimeInForce:    "GoodTillCancel",
			CloseOnTrigger: true,
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})

	// TODO: Fix this skipped test
	t.Run("Replace Conditional Order - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		t.Skip("Need to add a scenatio to test this. Currently the test account has some problems with USD asset")
		response, err := bybitClient.ReplaceConditionalOrder(context.Background(), &inverseperp.ReplaceConditionalOrderParams{
			StopOrderID:    "",
			OrderLinkID:    "",
			Symbol:         "",
			PRQty:          0,
			PRPrice:        "",
			PRTriggerPrice: "",
			TakeProfit:     0,
			StopLoss:       0,
			TPTriggerBy:    "",
			SLTriggerBy:    "",
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})

	t.Run("Get Conditional Order - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.GetConditionalOrder(context.Background(), &inverseperp.GetConditionalOrderParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})

	// TODO: Fix this skipped test
	t.Run("Cancel Conditional Order - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		t.Skip("Need to add a scenatio to test this. Currently the test account has some problems with USD asset")
		response, err := bybitClient.CancelConditionalOrder(context.Background(), &inverseperp.CancelConditionalOrderParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})

	t.Run("Cancel All Conditional Orders - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.CancelAllConditionalOrders(context.Background(), &inverseperp.CancelAllConditionalOrdersParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})
}
