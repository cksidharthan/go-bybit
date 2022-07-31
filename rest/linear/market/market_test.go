package market_test

import (
	"context"
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
	linearRest "github.com/cksidharthan/go-bybit/rest/linear"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestClient_Spot_Market(t *testing.T) {
	bybitClient := linearRest.New(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	t.Run("Get Order Book - LINEAR", func(t *testing.T) {
		t.Parallel()
		symbols, err := bybitClient.Market().GetOrderBook(context.Background(), &linear.OrderBookParams{
			Symbol:   "BTCUSDT",
			Interval: bybit.Interval1Min,
			From:     time.Now().Add(-10 * time.Minute).Second(),
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, symbols.RetCode)
		assert.NotEmpty(t, symbols)
		assert.NotNil(t, symbols)
	})
}
