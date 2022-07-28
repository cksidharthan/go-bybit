package account_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/cksidharthan/go-bybit/rest/domain/spot"
	spotRest "github.com/cksidharthan/go-bybit/rest/spot"

	"github.com/stretchr/testify/assert"

	"github.com/cksidharthan/go-bybit/bybit"
)

// the tests are not run in parallel as we test the buy and sell orders
func TestSpotClient_Account(t *testing.T) {
	// client := spotRest.New(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))
	client := spotRest.New(bybit.BybitTestnetBaseURL, "J4f3nQx9dAaKy9hs0P", "kamI8ZtvxfXE6K3hcPfkh9YlFMyenfBOl5jC")

	t.Run("Place Active Order - Buy - SPOT", func(t *testing.T) {
		order := spot.PlaceActiveOrderParams{
			Symbol: "BTCUSDT",
			Side:   bybit.SideSpotBuy,
			Type:   bybit.OrderTypeSpotLimit,
			Price:  15000,
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
		assert.Equal(t, "15000", result.Result.Price)
		assert.Equal(t, fmt.Sprintf("%.3f", 0.004), result.Result.OrigQty)
		assert.Equal(t, bybit.TimeInForce("GTC"), result.Result.TimeInForce)
	})

	// Skipped Test
	t.Run("Place Active Order - Sell - SPOT", func(t *testing.T) {
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
	t.Run("Get Active Order - SPOT", func(t *testing.T) {
		var orderID string
		{
			order := spot.PlaceActiveOrderParams{
				Symbol: "BTCUSDT",
				Side:   bybit.SideSpotBuy,
				Type:   bybit.OrderTypeSpotLimit,
				Price:  15000,
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

	t.Run("Cancel Active Order - SPOT", func(t *testing.T) {
		var orderID string
		{
			order := spot.PlaceActiveOrderParams{
				Symbol: "BTCUSDT",
				Side:   bybit.SideSpotBuy,
				Type:   bybit.OrderTypeSpotLimit,
				Price:  15000,
				Qty:    0.004,
			}
			result, err := client.Account().PlaceActiveOrder(context.Background(), order)
			assert.NoError(t, err)
			assert.NotNil(t, result)
			orderID = result.Result.OrderID
		}

		params := spot.CancelActiveOrderParams{
			OrderID: orderID,
		}
		result, err := client.Account().CancelActiveOrder(context.Background(), params)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 0, result.RetCode)
		assert.Equal(t, "OK", result.RetMsg)
	})
}
