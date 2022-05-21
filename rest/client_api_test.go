package rest

import (
	"context"
	"github.com/cksidharthan/go-bybit/helpers"
	"github.com/cksidharthan/go-bybit/rest/filters"
	"github.com/stretchr/testify/assert"
	"testing"
)

// The tests in this file will hit the bybit testnet API.
// These tests are skipped by default.
// Will be run manually to check if there is any changes to the Bybit API

func TestSpotClient_API_GetSymbols(t *testing.T) {
	// t.Skip("Skipping tests to be run manually")

	bybitClient := New(helpers.BYBIT_TESTNET_BASE_URL, "rwEwhfC6mDFYIGfcyb", "yfNJSzGapfFwbJyvguAyVXLJSIOCIegBg42Z")

	t.Run("happy path - get symbols", func(t *testing.T) {
		t.Parallel()
		symbols, err := bybitClient.Spot().GetSymbols(context.Background())
		assert.NoError(t, err)
		assert.NotEmpty(t, symbols)
		assert.NotNil(t, symbols)
	})

	t.Run("happy path - get order book depth", func(t *testing.T) {
		t.Parallel()
		orderBook, err := bybitClient.Spot().GetOrderBookDepth(context.Background(), &filters.OrderBookDepthFilter{
			Symbol: "BTCUSDT",
			Limit:  "10",
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, orderBook)
		assert.NotNil(t, orderBook)
	})

	t.Run("400 - get order book depth", func(t *testing.T) {
		orderBook, err := bybitClient.Spot().GetOrderBookDepth(context.Background(), &filters.OrderBookDepthFilter{
			Symbol: "BTCUSD",
			Limit:  "10",
		})
		assert.Error(t, err)
		assert.Empty(t, orderBook)
		assert.Nil(t, orderBook)
		assert.Equal(t, "HTTP 400: Not supported symbols", err.Error())
	})

	t.Run("happy path - get merged order book", func(t *testing.T) {
		t.Parallel()
		orderBook, err := bybitClient.Spot().GetMergedOrderBook(context.Background(), &filters.MergedOrderBookFilter{
			Symbol: "BTCUSDT",
			Limit:  "10",
			Scale:  "1",
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, orderBook)
		assert.NotNil(t, orderBook)
	})

	t.Run("400 - get merged order book ", func(t *testing.T) {
		orderBook, err := bybitClient.Spot().GetMergedOrderBook(context.Background(), &filters.MergedOrderBookFilter{
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
		publicTradeRecords, err := bybitClient.Spot().GetPublicTradeRecords(context.Background(), &filters.TradeRecordsFilter{
			Symbol: "BTCUSDT",
			Limit:  "10",
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, publicTradeRecords)
		assert.NotNil(t, publicTradeRecords)
		assert.Equal(t, 10, len(publicTradeRecords.Result))
	})

	t.Run("400 - get public records", func(t *testing.T) {
		publicTradeRecords, err := bybitClient.Spot().GetPublicTradeRecords(context.Background(), &filters.TradeRecordsFilter{
			Symbol: "BTCUSD",
			Limit:  "10",
		})
		assert.Error(t, err)
		assert.Empty(t, publicTradeRecords)
		assert.Nil(t, publicTradeRecords)
		assert.Equal(t, "HTTP 400: Invalid Symbols!", err.Error())
	})

	t.Run("happy path - get kline", func(t *testing.T) {
		klines, err := bybitClient.Spot().GetKline(context.Background(), &filters.KlineFilter{
			Symbol:   "BTCUSDT",
			Limit:    "10",
			Interval: "1m",
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, klines)
		assert.NotNil(t, klines)
		assert.Equal(t, 10, len(klines.Result))
	})

	t.Run("400 - get kline", func(t *testing.T) {
		klines, err := bybitClient.Spot().GetKline(context.Background(), &filters.KlineFilter{
			Symbol:   "BTCUSDT",
			Limit:    "10",
			Interval: "1",
		})
		assert.Error(t, err)
		assert.Empty(t, klines)
		assert.Nil(t, klines)
		assert.Equal(t, "HTTP 400: Invalid period!", err.Error())
	})

	t.Run("happy path - get ticker 24hr", func(t *testing.T) {
		tickerInfo, err := bybitClient.Spot().GetTickerInfo24hr(context.Background(), "BTCUSDT")
		assert.NoError(t, err)
		assert.NotEmpty(t, tickerInfo)
		assert.NotNil(t, tickerInfo)
		assert.NotNil(t, tickerInfo.Result)
	})

	t.Run("400 - get ticker 24hr", func(t *testing.T) {
		tickerInfo, err := bybitClient.Spot().GetTickerInfo24hr(context.Background(), "BTCUSD")
		assert.Error(t, err)
		assert.Empty(t, tickerInfo)
		assert.Nil(t, tickerInfo)
		assert.Equal(t, "HTTP 400: Invalid Symbols!", err.Error())
	})

	t.Run("happy path - get last traded price", func(t *testing.T) {
		tradePrice, err := bybitClient.Spot().GetLastTradedPrice(context.Background(), "BTCUSDT")
		assert.NoError(t, err)
		assert.NotEmpty(t, tradePrice)
		assert.NotNil(t, tradePrice)
		assert.NotNil(t, tradePrice.Result)
	})

	t.Run("400 - get last traded price", func(t *testing.T) {
		tradePrice, err := bybitClient.Spot().GetLastTradedPrice(context.Background(), "BTCUSD")
		assert.Error(t, err)
		assert.Empty(t, tradePrice)
		assert.Nil(t, tradePrice)
		assert.Equal(t, "HTTP 400: Invalid Symbols!", err.Error())
	})

	t.Run("happy path - get bid ask price", func(t *testing.T) {
		bidAskPrice, err := bybitClient.Spot().GetBidAskPrice(context.Background(), "BTCUSDT")
		assert.NoError(t, err)
		assert.NotEmpty(t, bidAskPrice)
		assert.NotNil(t, bidAskPrice)
		assert.NotNil(t, bidAskPrice.Result)
	})

	t.Run("400 - get last traded price", func(t *testing.T) {
		bidAskPrice, err := bybitClient.Spot().GetBidAskPrice(context.Background(), "BTCUSD")
		assert.Error(t, err)
		assert.Empty(t, bidAskPrice)
		assert.Nil(t, bidAskPrice)
		assert.Equal(t, "HTTP 400: Invalid Symbols!", err.Error())
	})

	t.Run("happy path - get server time", func(t *testing.T) {
		t.Parallel()
		serverTime, err := bybitClient.Spot().GetServerTime(context.Background())
		assert.NoError(t, err)
		assert.NotEmpty(t, serverTime)
		assert.NotNil(t, serverTime)
	})
}
