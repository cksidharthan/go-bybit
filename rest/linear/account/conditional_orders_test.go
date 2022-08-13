package account_test

import (
	"context"
	"fmt"
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
	linearRest "github.com/cksidharthan/go-bybit/rest/linear"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestClient_Linear_Account(t *testing.T) {
	t.Parallel()
	// bybitClient := linearRest.NewLinearClient(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))
	bybitClient := linearRest.NewLinearClient(bybit.BybitTestnetBaseURL, "J4f3nQx9dAaKy9hs0P", "kamI8ZtvxfXE6K3hcPfkh9YlFMyenfBOl5jC")
	currentADAPrice, err := getLinearADABuyPriceForTest(bybitClient)
	if err != nil {
		t.Fatalf("failed to get current ADA price: %v", err)
	}
	preferredADABuyPrice, _ := bybit.GetPrecision(*currentADAPrice + 1)
	fmt.Println("preferred ADA buy price:", preferredADABuyPrice)

	t.Run("Place Conditional Order - LINEAR", func(t *testing.T) {
		t.Parallel()
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
	})

}

func getLinearADABuyPriceForTest(client linearRest.Interface) (*float64, error) {
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
	fmt.Println("current ADA price:", currentADAPrice)
	return &currentADAPrice, nil
}
