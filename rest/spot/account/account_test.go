package account_test

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/cksidharthan/go-bybit/rest/domain/spot"
	spotRest "github.com/cksidharthan/go-bybit/rest/spot"
	"github.com/stretchr/testify/assert"

	"github.com/cksidharthan/go-bybit/bybit"
)

// the tests are not run in parallel as we test the buy and sell orders
func TestSpotClient_Account(t *testing.T) {
	client := spotRest.NewSpotClient(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	currentBTCPrice, err := getSpotBTCBuyPriceForTest(client)
	if err != nil {
		t.Fatalf("failed to get current BTC price: %v", err)
	}
	// Do we need an error here? since the currentBTC price is always float
	preferredBTCBuyPrice, _ := bybit.GetPrecision(*currentBTCPrice - 1000.00)

	t.Run("Place Active Order - Buy - SPOT", func(t *testing.T) {
		order := spot.PlaceActiveOrderParams{
			Symbol: "BTCUSDT",
			Side:   bybit.SideSpotBuy,
			Type:   bybit.OrderTypeSpotLimit,
			Price:  preferredBTCBuyPrice,
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
		assert.Equal(t, preferredBTCBuyPrice, result.Result.Price)
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
				Price:  preferredBTCBuyPrice,
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
				Price:  preferredBTCBuyPrice,
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

	t.Run("Fast Cancel Active Order - SPOT", func(t *testing.T) {
		var orderID string
		{
			order := spot.PlaceActiveOrderParams{
				Symbol: "BTCUSDT",
				Side:   bybit.SideSpotBuy,
				Type:   bybit.OrderTypeSpotLimit,
				Price:  preferredBTCBuyPrice,
				Qty:    0.004,
			}
			result, err := client.Account().PlaceActiveOrder(context.Background(), order)
			assert.NoError(t, err)
			assert.NotNil(t, result)
			orderID = result.Result.OrderID
		}
		params := spot.FastCancelActiveOrderParams{
			SymbolID: "BTCUSDT",
			OrderID:  orderID,
		}
		result, err := client.Account().FastCancelActiveOrder(context.Background(), params)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 0, result.RetCode)
		assert.Equal(t, "OK", result.RetMsg)
		assert.Equal(t, true, result.Result.IsCancelled)
	})

	t.Run("Batch Cancel Active Order - SPOT", func(t *testing.T) {
		{
			order := spot.PlaceActiveOrderParams{
				Symbol: "BTCUSDT",
				Side:   bybit.SideSpotBuy,
				Type:   bybit.OrderTypeSpotLimit,
				Price:  preferredBTCBuyPrice,
				Qty:    0.004,
			}
			result, err := client.Account().PlaceActiveOrder(context.Background(), order)
			assert.NoError(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, "OK", result.RetMsg)
			assert.Equal(t, 0, result.RetCode)
		}
		params := spot.BatchCancelActiveOrderParams{
			Symbol: "BTCUSDT",
		}
		result, err := client.Account().BatchCancelActiveOrder(context.Background(), params)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 0, result.RetCode)
		assert.Equal(t, "OK", result.RetMsg)
		assert.Equal(t, true, result.Result.Success)
	})

	t.Run("Batch Fast Cancel Active Order - SPOT", func(t *testing.T) {
		{
			order := spot.PlaceActiveOrderParams{
				Symbol: "BTCUSDT",
				Side:   bybit.SideSpotBuy,
				Type:   bybit.OrderTypeSpotLimit,
				Price:  preferredBTCBuyPrice,
				Qty:    0.004,
			}
			result, err := client.Account().PlaceActiveOrder(context.Background(), order)
			assert.NoError(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, "OK", result.RetMsg)
			assert.Equal(t, 0, result.RetCode)
		}
		{
			price, _ := bybit.GetPrecision(preferredBTCBuyPrice - 2000.00)
			order := spot.PlaceActiveOrderParams{
				Symbol: "BTCUSDT",
				Side:   bybit.SideSpotBuy,
				Type:   bybit.OrderTypeSpotLimit,
				Price:  price,
				Qty:    0.004,
			}
			result, err := client.Account().PlaceActiveOrder(context.Background(), order)
			assert.NoError(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, "OK", result.RetMsg)
			assert.Equal(t, 0, result.RetCode)
		}
		params := spot.BatchFastCancelActiveOrderParams{
			Symbol: "BTCUSDT",
		}
		result, err := client.Account().BatchFastCancelActiveOrder(context.Background(), params)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 0, result.RetCode)
		assert.Equal(t, "OK", result.RetMsg)
		assert.Equal(t, true, result.Result.Success)
	})

	t.Run("Batch Fast Cancel Active Order by IDs - SPOT", func(t *testing.T) {
		var orderIds []string
		{
			order := spot.PlaceActiveOrderParams{
				Symbol: "BTCUSDT",
				Side:   bybit.SideSpotBuy,
				Type:   bybit.OrderTypeSpotLimit,
				Price:  preferredBTCBuyPrice,
				Qty:    0.004,
			}
			result, err := client.Account().PlaceActiveOrder(context.Background(), order)
			assert.NoError(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, "OK", result.RetMsg)
			assert.Equal(t, 0, result.RetCode)
			orderIds = append(orderIds, result.Result.OrderID)
		}
		{
			price, _ := bybit.GetPrecision(preferredBTCBuyPrice - 2000.00)
			order := spot.PlaceActiveOrderParams{
				Symbol: "BTCUSDT",
				Side:   bybit.SideSpotBuy,
				Type:   bybit.OrderTypeSpotLimit,
				Price:  price,
				Qty:    0.004,
			}
			result, err := client.Account().PlaceActiveOrder(context.Background(), order)
			assert.NoError(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, "OK", result.RetMsg)
			assert.Equal(t, 0, result.RetCode)
			orderIds = append(orderIds, result.Result.OrderID)
		}
		params := spot.BatchCancelActiveOrdersByIDsParams{
			OrderIDs: strings.Join(orderIds, ","),
		}
		result, err := client.Account().BatchCancelActiveOrdersByIDs(context.Background(), params)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 0, result.RetCode)
		assert.Equal(t, "OK", result.RetMsg)
	})

	t.Run("Get Open Orders - SPOT", func(t *testing.T) {
		{
			order := spot.PlaceActiveOrderParams{
				Symbol: "BTCUSDT",
				Side:   bybit.SideSpotBuy,
				Type:   bybit.OrderTypeSpotLimit,
				Price:  preferredBTCBuyPrice,
				Qty:    0.004,
			}
			result, err := client.Account().PlaceActiveOrder(context.Background(), order)
			assert.NoError(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, "OK", result.RetMsg)
			assert.Equal(t, 0, result.RetCode)
		}
		{
			price, _ := bybit.GetPrecision(preferredBTCBuyPrice - 2000.00)
			order := spot.PlaceActiveOrderParams{
				Symbol: "BTCUSDT",
				Side:   bybit.SideSpotBuy,
				Type:   bybit.OrderTypeSpotLimit,
				Price:  price,
				Qty:    0.004,
			}
			result, err := client.Account().PlaceActiveOrder(context.Background(), order)
			assert.NoError(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, "OK", result.RetMsg)
			assert.Equal(t, 0, result.RetCode)
		}
		params := spot.GetOpenOrdersParams{
			Symbol: "BTCUSDT",
		}
		result, err := client.Account().GetOpenOrders(context.Background(), params)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 0, result.RetCode)
		assert.Equal(t, "OK", result.RetMsg)
		assert.Equal(t, 2, len(result.Result))

		// cancel all orders after testing
		cancelParams := spot.BatchFastCancelActiveOrderParams{
			Symbol: "BTCUSDT",
		}
		cancelResult, err := client.Account().BatchFastCancelActiveOrder(context.Background(), cancelParams)
		assert.Nil(t, err)
		assert.NotNil(t, cancelResult)
		assert.Equal(t, 0, cancelResult.RetCode)
		assert.Equal(t, "OK", cancelResult.RetMsg)
		assert.Equal(t, true, cancelResult.Result.Success)
	})

	t.Run("Get Order History - SPOT", func(t *testing.T) {
		params := spot.GetOrderHistoryParams{
			Symbol: "BTCUSDT",
		}
		result, err := client.Account().GetOrderHistory(context.Background(), params)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 0, result.RetCode)
		assert.Equal(t, "OK", result.RetMsg)
		assert.LessOrEqual(t, 0, len(result.Result))
	})

	t.Run("Get Trade History - SPOT", func(t *testing.T) {
		params := spot.GetTradeHistoryParams{
			Symbol: "BTCUSDT",
		}
		result, err := client.Account().GetTradeHistory(context.Background(), params)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 0, result.RetCode)
		assert.Equal(t, "OK", result.RetMsg)
		assert.LessOrEqual(t, 0, len(result.Result))
	})
}

func getSpotBTCBuyPriceForTest(client spotRest.Interface) (*float64, error) {
	response, err := client.Market().GetLastTradedPrice(context.Background(), &spot.GetLastTradedPriceParams{
		Symbol: "BTCUSDT",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get current BTC price: %v", err)
	}
	currentBTCPrice, err := strconv.ParseFloat(response.Result.Price, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse current BTC price: %v", err)
	}
	currentBTCPrice, err = bybit.GetPrecision(currentBTCPrice)
	if err != nil {
		return nil, fmt.Errorf("failed to get current BTC price: %v", err)
	}
	fmt.Println("current BTC price:", currentBTCPrice)
	return &currentBTCPrice, nil
}
