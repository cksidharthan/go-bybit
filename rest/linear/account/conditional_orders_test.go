package account_test

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
	linearRest "github.com/cksidharthan/go-bybit/rest/linear"
	"github.com/stretchr/testify/assert"
)

func TestClient_Linear_Account(t *testing.T) {
	bybitClient := linearRest.NewLinearClient(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	currentADAPrice, err := getLinearADABuyPriceForTest(bybitClient)
	if err != nil {
		t.Fatalf("failed to get current ADA price: %v", err)
	}
	preferredADABuyPrice, _ := bybit.GetPrecision(*currentADAPrice + 0.8)

	t.Run("Place Conditional Order - LINEAR", func(t *testing.T) {
		response, err := bybitClient.Account().PlaceConditionalOrder(context.Background(), &linear.PlaceConditionalOrderParams{
			Side:           bybit.SideBuy,
			Symbol:         "ADAUSDT",
			OrderType:      bybit.OrderTypeLimit,
			Qty:            10,
			Price:          preferredADABuyPrice,
			TimeInForce:    bybit.TimeInForceGoodTillCancel,
			BasePrice:      preferredADABuyPrice,
			StopPx:         preferredADABuyPrice + 0.5,
			TriggerBy:      "MarkPrice",
			ReduceOnly:     false,
			CloseOnTrigger: false,
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Get Conditional Order - LINEAR", func(t *testing.T) {
		response, err := bybitClient.Account().GetConditionalOrder(context.Background(), &linear.GetConditionalOrderParams{
			Symbol: "ADAUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Cancel Conditional Order - LINEAR", func(t *testing.T) {
		stopOrderID := ""
		{
			response, err := bybitClient.Account().PlaceConditionalOrder(context.Background(), &linear.PlaceConditionalOrderParams{
				Side:           bybit.SideBuy,
				Symbol:         "ADAUSDT",
				OrderType:      bybit.OrderTypeLimit,
				Qty:            10,
				Price:          preferredADABuyPrice,
				TimeInForce:    bybit.TimeInForceGoodTillCancel,
				BasePrice:      preferredADABuyPrice,
				StopPx:         preferredADABuyPrice + 0.5,
				TriggerBy:      "LastPrice",
				ReduceOnly:     false,
				CloseOnTrigger: false,
			})
			assert.NoError(t, err)
			assert.Equal(t, 0, response.RetCode)
			assert.NotEmpty(t, response)
			assert.NotNil(t, response)
			stopOrderID = response.Result.StopOrderID
		}
		{
			response, err := bybitClient.Account().CancelConditionalOrder(context.Background(), &linear.CancelConditionalOrderParams{
				Symbol:      "ADAUSDT",
				StopOrderID: stopOrderID,
			})
			assert.NoError(t, err)
			assert.Equal(t, 0, response.RetCode)
			assert.NotEmpty(t, response)
			assert.NotNil(t, response)
		}
	})

	t.Run("Cancel All Conditional Order - LINEAR", func(t *testing.T) {
		{
			response, err := bybitClient.Account().PlaceConditionalOrder(context.Background(), &linear.PlaceConditionalOrderParams{
				Side:           bybit.SideBuy,
				Symbol:         "ADAUSDT",
				OrderType:      bybit.OrderTypeLimit,
				Qty:            10,
				Price:          preferredADABuyPrice,
				TimeInForce:    bybit.TimeInForceGoodTillCancel,
				BasePrice:      preferredADABuyPrice,
				StopPx:         preferredADABuyPrice + 0.5,
				TriggerBy:      "LastPrice",
				ReduceOnly:     false,
				CloseOnTrigger: false,
			})
			assert.NoError(t, err)
			assert.Equal(t, 0, response.RetCode)
			assert.NotEmpty(t, response)
			assert.NotNil(t, response)
		}
		{
			response, err := bybitClient.Account().CancelAllConditionalOrders(context.Background(), &linear.CancelAllConditionalOrdersParams{
				Symbol: "ADAUSDT",
			})
			assert.NoError(t, err)
			assert.Equal(t, 0, response.RetCode)
			assert.NotEmpty(t, response)
			assert.NotNil(t, response)
		}
	})

	t.Run("Replace Conditional Order - LINEAR", func(t *testing.T) {
		stopOrderID := ""
		{
			response, err := bybitClient.Account().PlaceConditionalOrder(context.Background(), &linear.PlaceConditionalOrderParams{
				Side:           bybit.SideBuy,
				Symbol:         "ADAUSDT",
				OrderType:      bybit.OrderTypeLimit,
				Qty:            10,
				Price:          preferredADABuyPrice,
				TimeInForce:    bybit.TimeInForceGoodTillCancel,
				BasePrice:      preferredADABuyPrice,
				StopPx:         preferredADABuyPrice + 0.5,
				TriggerBy:      "LastPrice",
				ReduceOnly:     false,
				CloseOnTrigger: false,
			})
			assert.NoError(t, err)
			assert.Equal(t, 0, response.RetCode)
			assert.NotEmpty(t, response)
			assert.NotNil(t, response)
			stopOrderID = response.Result.StopOrderID
		}
		{
			response, err := bybitClient.Account().ReplaceConditionalOrder(context.Background(), &linear.ReplaceConditionalOrderParams{
				Symbol:      "ADAUSDT",
				StopOrderID: stopOrderID,
				TpTriggerBy: "MarkPrice",
			})
			assert.NoError(t, err)
			assert.Equal(t, 0, response.RetCode)
			assert.NotEmpty(t, response)
			assert.NotNil(t, response)
		}
	})

	t.Run("Query Conditional Order By Symbol - LINEAR", func(t *testing.T) {
		response, err := bybitClient.Account().QueryConditionalOrderBySymbol(context.Background(), &linear.QueryConditionalOrderBySymbolParams{
			Symbol: "ADAUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Query Conditional Order By Symbol - LINEAR", func(t *testing.T) {
		stopOrderID := ""
		{
			response, err := bybitClient.Account().PlaceConditionalOrder(context.Background(), &linear.PlaceConditionalOrderParams{
				Side:           bybit.SideBuy,
				Symbol:         "ADAUSDT",
				OrderType:      bybit.OrderTypeLimit,
				Qty:            10,
				Price:          preferredADABuyPrice,
				TimeInForce:    bybit.TimeInForceGoodTillCancel,
				BasePrice:      preferredADABuyPrice,
				StopPx:         preferredADABuyPrice + 0.5,
				TriggerBy:      "LastPrice",
				ReduceOnly:     false,
				CloseOnTrigger: false,
			})
			assert.NoError(t, err)
			assert.Equal(t, 0, response.RetCode)
			assert.NotEmpty(t, response)
			assert.NotNil(t, response)
			stopOrderID = response.Result.StopOrderID
		}
		{
			response, err := bybitClient.Account().QueryConditionalOrderWithIDs(context.Background(), &linear.QueryConditionalOrderWithIDsParams{
				Symbol:      "ADAUSDT",
				StopOrderID: stopOrderID,
			})
			assert.NoError(t, err)
			assert.Equal(t, 0, response.RetCode)
			assert.NotEmpty(t, response)
			assert.NotNil(t, response)
		}
	})
}

func getLinearADABuyPriceForTest(client *linearRest.Client) (*float64, error) {
	response, err := client.Market().GetSymbolInformation(context.Background(), &linear.GetSymbolInformationParams{
		Symbol: "ADAUSDT",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get current ADA price: %v", err)
	}
	currentADAPrice, err := strconv.ParseFloat(response.Result[0].LastPrice, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse current ADA price: %v", err)
	}
	currentADAPrice, err = bybit.GetPrecision(currentADAPrice)
	if err != nil {
		return nil, fmt.Errorf("failed to get current ADA price: %v", err)
	}
	return &currentADAPrice, nil
}
