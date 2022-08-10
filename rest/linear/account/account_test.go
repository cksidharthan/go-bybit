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
	// Do we need an error here? since the currentBTC price is always float
	preferredBTCBuyPrice, _ := bybit.GetPrecision(*currentBTCPrice - 1000.00)
	fmt.Println("preferred BTC buy price:", preferredBTCBuyPrice)

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
}

func getLinearBTCBuyPriceForTest(client linearRest.Interface) (*float64, error) {
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
	fmt.Println("current BTC price:", currentBTCPrice)
	return &currentBTCPrice, nil
}
