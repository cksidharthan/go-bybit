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

func TestClient_Linear_Market(t *testing.T) {
	t.Parallel()
	bybitClient := linearRest.NewLinearClient(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	currentBTCPrice, err := getLinearBTCBuyPriceForTest(bybitClient)
	if err != nil {
		t.Fatalf("failed to get current BTC price: %v", err)
	}

	t.Run("Place Active Order - LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Account().PlaceActiveOrder(context.Background(), &linear.PlaceActiveOrderParams{
			Side:           bybit.SideBuy,
			Symbol:         "BTCUSDT",
			OrderType:      bybit.OrderTypeLimit,
			Qty:            0.004,
			Price:          *currentBTCPrice,
			TimeInForce:    bybit.TimeInForceGoodTillCancel,
			ReduceOnly:     false,
			CloseOnTrigger: false,
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Get Active Order - LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Account().GetActiveOrder(context.Background(), &linear.GetActiveOrderParams{
			Symbol: "BTCUSDT",
			Limit:  1,
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
		assert.Equal(t, 1, len(response.Result.Data))
		assert.Equal(t, "BTCUSDT", response.Result.Data[0].Symbol)
	})

	t.Run("Cancel all Active orders - LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Account().CancelAllActiveOrders(context.Background(), &linear.CancelAllActiveOrdersParams{
			Symbol: "BTCUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Replace Active Order - LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Account().ReplaceActiveOrder(context.Background(), &linear.ReplaceActiveOrderParams{
			Symbol: "BTCUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Query Active Order - LINEAR", func(t *testing.T) {
		t.Parallel()
		// place active order to query it
		{
			response, err := bybitClient.Account().PlaceActiveOrder(context.Background(), &linear.PlaceActiveOrderParams{
				Side:           bybit.SideBuy,
				Symbol:         "BTCUSDT",
				OrderType:      bybit.OrderTypeLimit,
				Qty:            0.004,
				Price:          *currentBTCPrice,
				TimeInForce:    bybit.TimeInForceGoodTillCancel,
				ReduceOnly:     false,
				CloseOnTrigger: false,
			})
			assert.NoError(t, err)
			assert.Equal(t, 0, response.RetCode)
			assert.NotEmpty(t, response)
			assert.NotNil(t, response)
		}
		{
			response, err := bybitClient.Account().PlaceActiveOrder(context.Background(), &linear.PlaceActiveOrderParams{
				Side:           bybit.SideBuy,
				Symbol:         "BTCUSDT",
				OrderType:      bybit.OrderTypeLimit,
				Qty:            0.004,
				Price:          *currentBTCPrice,
				TimeInForce:    bybit.TimeInForceGoodTillCancel,
				ReduceOnly:     false,
				CloseOnTrigger: false,
			})
			assert.NoError(t, err)
			assert.Equal(t, 0, response.RetCode)
			assert.NotEmpty(t, response)
			assert.NotNil(t, response)
		}
		{
			response, err := bybitClient.Account().QueryActiveOrder(context.Background(), &linear.QueryActiveOrderParams{
				Symbol: "BTCUSDT",
			})
			assert.NoError(t, err)
			assert.Equal(t, 0, response.RetCode)
			assert.NotEmpty(t, response)
			assert.NotNil(t, response)
		}
	})
}

func getLinearBTCBuyPriceForTest(client *linearRest.Client) (*float64, error) {
	response, err := client.Market().GetSymbolInformation(context.Background(), &linear.GetSymbolInformationParams{
		Symbol: "BTCUSDT",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get current BTC price: %v", err)
	}
	currentBTCPrice, err := strconv.ParseFloat(response.Result[0].LastPrice, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse current BTC price: %v", err)
	}
	currentBTCPrice, err = bybit.GetPrecision(currentBTCPrice)
	if err != nil {
		return nil, fmt.Errorf("failed to get current BTC price: %v", err)
	}
	return &currentBTCPrice, nil
}
