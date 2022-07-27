package market_test

import (
	"context"
	"os"
	"testing"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot/market/filters"
	"github.com/cksidharthan/go-bybit/rest/spot"
	"github.com/stretchr/testify/assert"
)

// The tests in this file will hit the bybit testnet API.
func TestClient_Spot(t *testing.T) {
	bybitClient := spot.New(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	t.Run("happy path - get symbols", func(t *testing.T) {
		t.Parallel()
		symbols, err := bybitClient.Market().GetSymbols(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 0, symbols.RetCode)
		assert.NotEmpty(t, symbols)
		assert.NotNil(t, symbols)
	})

	t.Run("happy path - get order book depth", func(t *testing.T) {
		t.Parallel()
		orderBook, err := bybitClient.Market().GetOrderBookDepth(context.Background(), &filters.OrderBookDepthFilter{
			Symbol: "BTCUSDT",
			Limit:  "10",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, orderBook.RetCode)
		assert.NotEmpty(t, orderBook)
		assert.NotNil(t, orderBook)
	})

	t.Run("400 - get order book depth", func(t *testing.T) {
		orderBook, err := bybitClient.Market().GetOrderBookDepth(context.Background(), &filters.OrderBookDepthFilter{
			Symbol: "BTCUSD",
			Limit:  "10",
		})
		assert.NoError(t, err)
		assert.NotEqual(t, 0, orderBook.RetCode)
		assert.Equal(t, -100011, orderBook.RetCode)
		assert.Equal(t, "Not supported symbols", orderBook.RetMsg)
	})

	t.Run("happy path - get merged order book", func(t *testing.T) {
		t.Parallel()
		orderBook, err := bybitClient.Market().GetMergedOrderBook(context.Background(), &filters.MergedOrderBookFilter{
			Symbol: "BTCUSDT",
			Limit:  "10",
			Scale:  "1",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, orderBook.RetCode)
		assert.NotEmpty(t, orderBook)
		assert.NotNil(t, orderBook)
	})

	t.Run("400 - get merged order book ", func(t *testing.T) {
		orderBook, err := bybitClient.Market().GetMergedOrderBook(context.Background(), &filters.MergedOrderBookFilter{
			Symbol: "BTCUSDT",
			Limit:  "10",
			Scale:  "0.5",
		})
		assert.Error(t, err)
		assert.Empty(t, orderBook)
		assert.Nil(t, orderBook)
		assert.Equal(t, "HTTP 400: Param scale should be Integer.", err.Error())
	})

	t.Run("happy path - get public records", func(t *testing.T) {
		publicTradeRecords, err := bybitClient.Market().GetPublicTradeRecords(context.Background(), &filters.TradeRecordsFilter{
			Symbol: "BTCUSDT",
			Limit:  "10",
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, publicTradeRecords)
		assert.Equal(t, 0, publicTradeRecords.RetCode)
		assert.NotNil(t, publicTradeRecords)
		assert.Equal(t, 10, len(publicTradeRecords.Result))
	})

	t.Run("400 - get public records", func(t *testing.T) {
		publicTradeRecords, err := bybitClient.Market().GetPublicTradeRecords(context.Background(), &filters.TradeRecordsFilter{
			Symbol: "BTCUSD",
			Limit:  "10",
		})
		assert.NoError(t, err)
		assert.NotEqual(t, 0, publicTradeRecords.RetCode)
		assert.Equal(t, -100011, publicTradeRecords.RetCode)
		assert.Equal(t, "Not supported symbols", publicTradeRecords.RetMsg)
	})

	t.Run("happy path - get kline", func(t *testing.T) {
		klines, err := bybitClient.Market().GetKline(context.Background(), &filters.KlineFilter{
			Symbol:   "BTCUSDT",
			Limit:    "10",
			Interval: "1m",
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, klines)
		assert.Equal(t, 0, klines.RetCode)
		assert.NotNil(t, klines)
		assert.Equal(t, 10, len(klines.Result))
	})

	t.Run("400 - get kline", func(t *testing.T) {
		klines, err := bybitClient.Market().GetKline(context.Background(), &filters.KlineFilter{
			Symbol:   "BTCUSDT",
			Limit:    "10",
			Interval: "1",
		})
		assert.NoError(t, err)
		assert.Equal(t, -10009, klines.RetCode)
		assert.NotEqual(t, 0, klines.RetCode)
		assert.Equal(t, "Period required!", klines.RetMsg)
	})

	t.Run("happy path - get ticker 24hr", func(t *testing.T) {
		tickerInfo, err := bybitClient.Market().GetTickerInfo24hr(context.Background(), "BTCUSDT")
		assert.NoError(t, err)
		assert.NotEmpty(t, tickerInfo)
		assert.Equal(t, 0, tickerInfo.RetCode)
		assert.NotNil(t, tickerInfo)
		assert.NotNil(t, tickerInfo.Result)
	})

	t.Run("400 - get ticker 24hr", func(t *testing.T) {
		tickerInfo, err := bybitClient.Market().GetTickerInfo24hr(context.Background(), "BTCUSD")
		assert.NoError(t, err)
		assert.NotEqual(t, 0, tickerInfo.RetCode)
		assert.Equal(t, -100011, tickerInfo.RetCode)
		assert.Equal(t, "Not supported symbols", tickerInfo.RetMsg)
	})

	t.Run("happy path - get last traded price", func(t *testing.T) {
		tradePrice, err := bybitClient.Market().GetLastTradedPrice(context.Background(), "BTCUSDT")
		assert.NoError(t, err)
		assert.NotEmpty(t, tradePrice)
		assert.NotNil(t, tradePrice)
		assert.Equal(t, 0, tradePrice.RetCode)
		assert.NotNil(t, tradePrice.Result)
	})

	t.Run("400 - get last traded price", func(t *testing.T) {
		tradePrice, err := bybitClient.Market().GetLastTradedPrice(context.Background(), "BTCUSD")
		assert.NoError(t, err)
		assert.NotEqual(t, 0, tradePrice.RetCode)
		assert.Equal(t, -100011, tradePrice.RetCode)
		assert.Equal(t, "Not supported symbols", tradePrice.RetMsg)
	})

	t.Run("happy path - get bid ask price", func(t *testing.T) {
		bidAskPrice, err := bybitClient.Market().GetBidAskPrice(context.Background(), "BTCUSDT")
		assert.NoError(t, err)
		assert.NotEmpty(t, bidAskPrice)
		assert.Equal(t, 0, bidAskPrice.RetCode)
		assert.NotNil(t, bidAskPrice)
		assert.NotNil(t, bidAskPrice.Result)
	})

	t.Run("400 - get last traded price", func(t *testing.T) {
		bidAskPrice, err := bybitClient.Market().GetBidAskPrice(context.Background(), "BTCUSD")
		assert.NoError(t, err)
		assert.NotEqual(t, 0, bidAskPrice.RetCode)
		assert.Equal(t, -100011, bidAskPrice.RetCode)
		assert.Equal(t, "Not supported symbols", bidAskPrice.RetMsg)
	})

	t.Run("happy path - get server time", func(t *testing.T) {
		t.Parallel()
		serverTime, err := bybitClient.Market().GetServerTime(context.Background())
		assert.NoError(t, err)
		assert.NotEmpty(t, serverTime)
		assert.Equal(t, 0, serverTime.RetCode)
		assert.NotNil(t, serverTime)
	})
}