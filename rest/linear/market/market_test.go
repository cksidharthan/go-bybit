package market_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
	linearRest "github.com/cksidharthan/go-bybit/rest/linear"
	"github.com/stretchr/testify/assert"
)

func TestClient_Linear_Market(t *testing.T) {
	t.Parallel()
	bybitClient := linearRest.New(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	t.Run("Get Order Book - LINEAR", func(t *testing.T) {
		t.Parallel()
		orderBook, err := bybitClient.Market().GetOrderBook(context.Background(), &linear.OrderBookParams{
			Symbol: "BTCUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, orderBook.RetCode)
		assert.NotEmpty(t, orderBook)
		assert.NotNil(t, orderBook)
		assert.Equal(t, "BTCUSDT", orderBook.Result[0].Symbol)
	})

	t.Run("Query Kline - LINEAR", func(t *testing.T) {
		t.Parallel()
		orderBook, err := bybitClient.Market().QueryKline(context.Background(), &linear.QueryKlineParams{
			Symbol:   "BTCUSDT",
			Interval: bybit.Interval1Min,
			From:     time.Now().Add(-time.Hour).Second(),
			Limit:    3,
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, orderBook.RetCode)
		assert.NotEmpty(t, orderBook)
		assert.NotNil(t, orderBook)
		assert.Equal(t, "BTCUSDT", orderBook.Result[0].Symbol)
		assert.Equal(t, 3, len(orderBook.Result))
	})
}
