package rest

import (
	"context"
	"github.com/cksidharthan/go-bybit/rest/filters"
	"github.com/cksidharthan/go-bybit/testdata"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSpotClient_GetSymbols(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		server := testdata.FakeBybitServer(
			"/spot/v1/symbols",
			"../testdata/public/spot/get_symbols_response_200.json",
			http.StatusOK)
		bybitClient := New(server.URL, "", "")
		symbols, err := bybitClient.Spot().GetSymbols(context.Background())
		assert.NoError(t, err)
		assert.NotEmpty(t, symbols)
		assert.NotNil(t, symbols)
		assert.Equal(t, 54, len(symbols.Result))
	})
}

func TestSpotClient_GetOrderBookDepth(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		server := testdata.FakeBybitServer(
			"/spot/quote/v1/depth?limit=10&symbol=BTCUSDT",
			"../testdata/public/spot/get_order_book_depth_200.json",
			http.StatusOK)
		bybitClient := New(server.URL, "", "")
		orderBook, err := bybitClient.Spot().GetOrderBookDepth(context.Background(), &filters.OrderBookDepthFilter{
			Symbol: "BTCUSDT",
			Limit:  "10",
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, orderBook)
		assert.NotNil(t, orderBook)
		assert.Equal(t, 10, len(orderBook.Result.Asks))
	})

	t.Run("invalid path - 400", func(t *testing.T) {
		server := testdata.FakeBybitServer(
			"/spot/quote/v1/depth?limit=10&symbol=BTCUSD",
			"../testdata/public/spot/get_order_book_depth_400.json",
			http.StatusBadRequest)
		bybitClient := New(server.URL, "", "")
		orderBook, err := bybitClient.Spot().GetOrderBookDepth(context.Background(), &filters.OrderBookDepthFilter{
			Symbol: "BTCUSD",
			Limit:  "10",
		})
		assert.Error(t, err)
		assert.Empty(t, orderBook)
		assert.Nil(t, orderBook)
		assert.Equal(t, "HTTP 400: Not supported symbols", err.Error())
	})
}

func TestSpotClient_GetMergedOrderBook(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		server := testdata.FakeBybitServer(
			"/spot/quote/v1/depth/merged?limit=10&scale=1&symbol=BTCUSDT",
			"../testdata/public/spot/get_merged_order_book_200.json",
			http.StatusOK)
		bybitClient := New(server.URL, "", "")
		orderBook, err := bybitClient.Spot().GetMergedOrderBook(context.Background(), &filters.MergedOrderBookFilter{
			Symbol: "BTCUSDT",
			Limit:  "10",
			Scale:  "1",
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, orderBook)
		assert.NotNil(t, orderBook)
		assert.Equal(t, 10, len(orderBook.Result.Asks))
	})

	t.Run("error path - 400", func(t *testing.T) {
		server := testdata.FakeBybitServer(
			"/spot/quote/v1/depth/merged?limit=10&scale=0.5&symbol=BTCUSDT",
			"../testdata/public/spot/get_merged_order_book_400.json",
			http.StatusBadRequest)
		bybitClient := New(server.URL, "", "")
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
}

func TestSpotClient_GetPublicTradeRecords(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		server := testdata.FakeBybitServer(
			"/spot/quote/v1/trades?limit=10&symbol=BTCUSDT",
			"../testdata/public/spot/get_public_trade_records_200.json",
			http.StatusOK)
		bybitClient := New(server.URL, "", "")
		publicTradeRecords, err := bybitClient.Spot().GetPublicTradeRecords(context.Background(), &filters.TradeRecordsFilter{
			Symbol: "BTCUSDT",
			Limit:  "10",
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, publicTradeRecords)
		assert.NotNil(t, publicTradeRecords)
		assert.Equal(t, 10, len(publicTradeRecords.Result))
	})

	t.Run("error path - 400", func(t *testing.T) {
		server := testdata.FakeBybitServer(
			"/spot/quote/v1/trades?limit=10&symbol=BTCUSDT",
			"../testdata/public/spot/get_public_trade_records_400.json",
			http.StatusBadRequest)
		bybitClient := New(server.URL, "", "")
		publicTradeRecords, err := bybitClient.Spot().GetPublicTradeRecords(context.Background(), &filters.TradeRecordsFilter{
			Symbol: "BTCUSDT",
			Limit:  "10",
		})
		assert.Error(t, err)
		assert.Empty(t, publicTradeRecords)
		assert.Nil(t, publicTradeRecords)
		assert.Equal(t, "HTTP 400: Invalid Symbols!", err.Error())
	})
}

func TestSpotClient_GetKlines(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		server := testdata.FakeBybitServer(
			"/spot/quote/v1/kline?interval=1m&limit=10&symbol=BTCUSDT",
			"../testdata/public/spot/get_klines_200.json",
			http.StatusOK)
		bybitClient := New(server.URL, "", "")
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

	t.Run("error path - 400", func(t *testing.T) {
		server := testdata.FakeBybitServer(
			"/spot/quote/v1/kline?interval=1&limit=10&symbol=BTCUSDT",
			"../testdata/public/spot/get_klines_400.json",
			http.StatusBadRequest)
		bybitClient := New(server.URL, "", "")
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
}

func TestSpotClient_GetTickerInfo24hr(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		server := testdata.FakeBybitServer(
			"/spot/quote/v1/ticker/24hr?symbol=BTCUSDT",
			"../testdata/public/spot/get_ticker_info_24hr_200.json",
			http.StatusOK)
		bybitClient := New(server.URL, "", "")
		tickerInfo, err := bybitClient.Spot().GetTickerInfo24hr(context.Background(), "BTCUSDT")
		assert.NoError(t, err)
		assert.NotEmpty(t, tickerInfo)
		assert.NotNil(t, tickerInfo)
		assert.Equal(t, "BTCUSDT", tickerInfo.Result.Symbol)
	})

	t.Run("error path - 400", func(t *testing.T) {
		server := testdata.FakeBybitServer(
			"/spot/quote/v1/ticker/24hr?symbol=BTCUSD",
			"../testdata/public/spot/get_ticker_info_24hr_400.json",
			http.StatusBadRequest)
		bybitClient := New(server.URL, "", "")
		tickerInfo, err := bybitClient.Spot().GetTickerInfo24hr(context.Background(), "BTCUSD")
		assert.Error(t, err)
		assert.Empty(t, tickerInfo)
		assert.Nil(t, tickerInfo)
		assert.Equal(t, "HTTP 400: Invalid Symbols!", err.Error())
	})
}

func TestSpotClient_GetLastTradedPrice(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		server := testdata.FakeBybitServer(
			"/spot/quote/v1/ticker/price?symbol=BTCUSDT",
			"../testdata/public/spot/get_last_traded_price_200.json",
			http.StatusOK)
		bybitClient := New(server.URL, "", "")
		tickerPrice, err := bybitClient.Spot().GetLastTradedPrice(context.Background(), "BTCUSDT")
		assert.NoError(t, err)
		assert.NotEmpty(t, tickerPrice)
		assert.NotNil(t, tickerPrice)
		assert.Equal(t, "BTCUSDT", tickerPrice.Result.Symbol)
	})

	t.Run("error path - 400", func(t *testing.T) {
		server := testdata.FakeBybitServer(
			"/spot/quote/v1/ticker/price?symbol=BTCUSD",
			"../testdata/public/spot/get_last_traded_price_400.json",
			http.StatusBadRequest)
		bybitClient := New(server.URL, "", "")
		tickerPrice, err := bybitClient.Spot().GetLastTradedPrice(context.Background(), "BTCUSD")
		assert.Error(t, err)
		assert.Empty(t, tickerPrice)
		assert.Nil(t, tickerPrice)
		assert.Equal(t, "HTTP 400: Invalid Symbols!", err.Error())
	})
}

func TestSpotClient_GetBidAskPrice(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		server := testdata.FakeBybitServer(
			"/spot/quote/v1/ticker/book_ticker?symbol=BTCUSDT",
			"../testdata/public/spot/get_bid_ask_price_200.json",
			http.StatusOK)
		bybitClient := New(server.URL, "", "")
		tickerBook, err := bybitClient.Spot().GetBidAskPrice(context.Background(), "BTCUSDT")
		assert.NoError(t, err)
		assert.NotEmpty(t, tickerBook)
		assert.NotNil(t, tickerBook)
		assert.Equal(t, "BTCUSDT", tickerBook.Result.Symbol)
	})

	t.Run("error path - 400", func(t *testing.T) {
		server := testdata.FakeBybitServer(
			"/spot/quote/v1/ticker/book_ticker?symbol=BTCUSD",
			"../testdata/public/spot/get_bid_ask_price_400.json",
			http.StatusBadRequest)
		bybitClient := New(server.URL, "", "")
		tickerBook, err := bybitClient.Spot().GetBidAskPrice(context.Background(), "BTCUSD")
		assert.Error(t, err)
		assert.Empty(t, tickerBook)
		assert.Nil(t, tickerBook)
		assert.Equal(t, "HTTP 400: Invalid Symbols!", err.Error())
	})
}

func TestSpotClient_GetServerTime(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		server := testdata.FakeBybitServer(
			"/spot/v1/time",
			"../testdata/public/spot/get_server_time_200.json",
			http.StatusOK)
		bybitClient := New(server.URL, "", "")
		symbols, err := bybitClient.Spot().GetServerTime(context.Background())
		assert.NoError(t, err)
		assert.NotEmpty(t, symbols)
		assert.NotNil(t, symbols)
	})
}
