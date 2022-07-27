package account_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cksidharthan/go-bybit/rest/domain/spot"
	spotRest "github.com/cksidharthan/go-bybit/rest/spot"

	"github.com/stretchr/testify/assert"

	"github.com/cksidharthan/go-bybit/bybit"
)

// the tests are not run in parallel as we test the buy and sell orders
func TestClient_Account(t *testing.T) {
	client := spotRest.New(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	t.Run("Place Active Order - Buy", func(t *testing.T) {
		order := spot.PlaceActiveOrderParams{
			Symbol: "BTCUSDT",
			Side:   bybit.SideSpotBuy,
			Type:   bybit.OrderTypeSpotLimit,
			Price:  22000,
			Qty:    0.004,
		}
		result, err := client.Account().PlaceActiveOrder(context.Background(), order)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, 0, result.RetCode)
		assert.Equal(t, "OK", result.RetMsg)
		assert.Equal(t, "BTCUSDT", result.Result.Symbol)
		assert.Equal(t, "BUY", result.Result.Side)
		assert.Equal(t, bybit.OrderTypeSpotLimit, result.Result.Type)
		assert.Equal(t, "22000", result.Result.Price)
		assert.Equal(t, fmt.Sprintf("%.3f", 0.004), result.Result.OrigQty)
		assert.Equal(t, bybit.TimeInForce("GTC"), result.Result.TimeInForce)
	})

	// Skipped Test
	t.Run("Place Active Order - Sell", func(t *testing.T) {
		t.Skip("Skipping this test as it will fail if there is a lag in buy order")
		order := spot.PlaceActiveOrderParams{
			Symbol: "BTCUSDT",
			Side:   bybit.SideSell,
			Type:   bybit.OrderTypeSpotLimit,
			Price:  22000,
			Qty:    0.004,
		}
		result, err := client.Account().PlaceActiveOrder(context.Background(), order)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, 0, result.RetCode)
		assert.Equal(t, "OK", result.RetMsg)
		assert.Equal(t, "BTCUSDT", result.Result.Symbol)
		assert.Equal(t, "SELL", result.Result.Side)
		assert.Equal(t, bybit.OrderTypeSpotLimit, result.Result.Type)
		assert.Equal(t, "22000", result.Result.Price)
		assert.Equal(t, fmt.Sprintf("%.3f", 0.004), result.Result.OrigQty)
		assert.Equal(t, bybit.TimeInForce("GTC"), result.Result.TimeInForce)
	})

	// Place an active order and get it
	t.Run("Get Active Order - No orders exist", func(t *testing.T) {
		var orderID string
		{
			order := spot.PlaceActiveOrderParams{
				Symbol: "BTCUSDT",
				Side:   bybit.SideSpotBuy,
				Type:   bybit.OrderTypeSpotLimit,
				Price:  22000,
				Qty:    0.004,
			}
			result, err := client.Account().PlaceActiveOrder(context.Background(), order)
			assert.NoError(t, err)
			assert.NotNil(t, result)
			orderID = result.Result.OrderID
		}

		params := spot.GetActiveOrderParams{
			OrderID: orderID,
		}
		result, err := client.Account().GetActiveOrder(context.Background(), params)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 0, result.RetCode)
		assert.Equal(t, "OK", result.RetMsg)
	})
}
