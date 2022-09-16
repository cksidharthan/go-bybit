package account_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cksidharthan/go-bybit/bybit"
	inverseperp "github.com/cksidharthan/go-bybit/rest/domain/inverse_perpetual"
	"github.com/cksidharthan/go-bybit/rest/inverse_perpetual/account"
	"github.com/stretchr/testify/assert"
)

func TestClient_ActiveOrders(t *testing.T) {
	t.Parallel()

	bybitClient := account.NewInversePerpetualAccountClient(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	t.Run("Place Active Order - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		t.Skipf("Skipping test %s, as there is some issue in adding USD in testnet account", t.Name())
		response, err := bybitClient.PlaceActiveOrder(context.Background(), &inverseperp.PlaceActiveOrderParams{
			Side:        string(bybit.SideBuy),
			Symbol:      "ADAUSD",
			OrderType:   string(bybit.OrderTypeMarket),
			Qty:         1,
			TimeInForce: string(bybit.TimeInForceGoodTillCancel),
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
		fmt.Println(response)
		assert.Equal(t, "BTCUSD", response.Result.Symbol)
	})

	t.Run("Get Active Orders - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.GetActiveOrder(context.Background(), &inverseperp.GetActiveOrderParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Cancel all Active orders - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.CancelAllActiveOrders(context.Background(), &inverseperp.CancelAllActiveOrdersParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Cancel Active Order - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		t.Skipf("Skipping test %s, as there is some issue in adding USD in testnet account", t.Name())
		response, err := bybitClient.CancelActiveOrder(context.Background(), &inverseperp.CancelActiveOrderParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Replace Active Order - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		t.Skipf("Skipping test %s, as there is some issue in adding USD in testnet account", t.Name())
		response, err := bybitClient.ReplaceActiveOrder(context.Background(), &inverseperp.ReplaceActiveOrderParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Query Active Order - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.QueryActiveOrder(context.Background(), &inverseperp.QueryActiveOrderParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})
}
