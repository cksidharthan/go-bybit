package market_test

import (
	"context"
	"fmt"
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
		kline, err := bybitClient.Market().QueryKline(context.Background(), &linear.QueryKlineParams{
			Symbol:   "BTCUSDT",
			Interval: bybit.Interval1Min,
			From:     time.Now().Add(-time.Hour).Second(),
			Limit:    3,
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, kline.RetCode)
		assert.NotEmpty(t, kline)
		assert.NotNil(t, kline)
		assert.Equal(t, "BTCUSDT", kline.Result[0].Symbol)
		assert.Equal(t, 3, len(kline.Result))
	})

	t.Run("Get Symbol Info - Single Symbol - LINEAR", func(t *testing.T) {
		t.Parallel()
		symbolInfo, err := bybitClient.Market().GetSymbolInformation(context.Background(), &linear.GetSymbolInformationParams{
			Symbol: "BTCUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, symbolInfo.RetCode)
		assert.NotEmpty(t, symbolInfo)
		assert.NotNil(t, symbolInfo)
		assert.Equal(t, "BTCUSDT", symbolInfo.Result[0].Symbol)
		assert.Equal(t, 1, len(symbolInfo.Result))
	})

	t.Run("Get Symbol Info - Multiple Symbols - LINEAR", func(t *testing.T) {
		t.Parallel()
		symbolInfo, err := bybitClient.Market().GetSymbolInformation(context.Background(), &linear.GetSymbolInformationParams{})
		assert.NoError(t, err)
		assert.Equal(t, 0, symbolInfo.RetCode)
		assert.NotEmpty(t, symbolInfo)
		assert.NotNil(t, symbolInfo)
		assert.GreaterOrEqual(t, len(symbolInfo.Result), 0)
	})

	t.Run("Get Public Trading Records - LINEAR", func(t *testing.T) {
		t.Parallel()
		symbolInfo, err := bybitClient.Market().GetPublicTradingRecords(context.Background(), &linear.GetPublicTradingRecordsParams{
			Symbol: "BTCUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, symbolInfo.RetCode)
		assert.NotEmpty(t, symbolInfo)
		assert.NotNil(t, symbolInfo)
		assert.GreaterOrEqual(t, len(symbolInfo.Result), 0)
	})

	t.Run("Get Public Trading Records - Limit - LINEAR", func(t *testing.T) {
		t.Parallel()
		symbolInfo, err := bybitClient.Market().GetPublicTradingRecords(context.Background(), &linear.GetPublicTradingRecordsParams{
			Symbol: "BTCUSDT",
			Limit:  3,
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, symbolInfo.RetCode)
		assert.NotEmpty(t, symbolInfo)
		assert.NotNil(t, symbolInfo)
		assert.Equal(t, 3, len(symbolInfo.Result))
	})

	t.Run("Query Symbol -  LINEAR", func(t *testing.T) {
		t.Parallel()
		symbolInfo, err := bybitClient.Market().QuerySymbol(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 0, symbolInfo.RetCode)
		assert.NotEmpty(t, symbolInfo)
		assert.NotNil(t, symbolInfo)
	})

	t.Run("Get Liquidated Orders -  LINEAR", func(t *testing.T) {
		t.Parallel()
		t.Skip("this endpoint is not available and sometimes goes offline")
		liquidatedOrders, err := bybitClient.Market().GetLiquidatedOrders(context.Background(), &linear.GetLiquidatedOrdersParams{
			Symbol: "BTCUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, liquidatedOrders.RetCode)
		assert.NotEmpty(t, liquidatedOrders)
		assert.Equal(t, "OK", liquidatedOrders.RetMsg)
		fmt.Println(liquidatedOrders)
		assert.NotNil(t, liquidatedOrders)
	})
}
