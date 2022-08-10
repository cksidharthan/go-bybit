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
		response, err := bybitClient.Market().GetOrderBook(context.Background(), &linear.OrderBookParams{
			Symbol: "BTCUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSDT", response.Result[0].Symbol)
	})

	t.Run("Query Kline - LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().QueryKline(context.Background(), &linear.QueryKlineParams{
			Symbol:   "BTCUSDT",
			Interval: bybit.Interval1Min,
			From:     time.Now().Add(-time.Hour).Second(),
			Limit:    3,
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSDT", response.Result[0].Symbol)
		assert.Equal(t, 3, len(response.Result))
	})

	t.Run("Get Symbol Info - Single Symbol - LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().GetSymbolInformation(context.Background(), &linear.GetSymbolInformationParams{
			Symbol: "BTCUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSDT", response.Result[0].Symbol)
		assert.Equal(t, 1, len(response.Result))
	})

	t.Run("Get Symbol Info - Multiple Symbols - LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().GetSymbolInformation(context.Background(), &linear.GetSymbolInformationParams{})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
		assert.GreaterOrEqual(t, len(response.Result), 0)
	})

	t.Run("Get Public Trading Records - LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().GetPublicTradingRecords(context.Background(), &linear.GetPublicTradingRecordsParams{
			Symbol: "BTCUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
		assert.GreaterOrEqual(t, len(response.Result), 0)
	})

	t.Run("Get Public Trading Records - Limit - LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().GetPublicTradingRecords(context.Background(), &linear.GetPublicTradingRecordsParams{
			Symbol: "BTCUSDT",
			Limit:  3,
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
		assert.Equal(t, 3, len(response.Result))
	})

	t.Run("Query Symbol -  LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().QuerySymbol(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Get Liquidated Orders -  LINEAR", func(t *testing.T) {
		t.Parallel()
		t.Skip("this endpoint is not available and sometimes goes offline")
		response, err := bybitClient.Market().GetLiquidatedOrders(context.Background(), &linear.GetLiquidatedOrdersParams{
			Symbol: "BTCUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.Equal(t, "OK", response.RetMsg)
		assert.NotNil(t, response)
	})

	t.Run("Query Mark Price Kline -  LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().QueryMarkPriceKline(context.Background(), &linear.QueryMarkPriceKlineParams{
			Symbol:   "BTCUSDT",
			Interval: bybit.Interval1Min,
			From:     time.Now().Add(-time.Hour).Second(),
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.Equal(t, "OK", response.RetMsg)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSDT", response.Result[0].Symbol)
	})

	t.Run("Query Mark Price Kline - Param validation error -  LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().QueryMarkPriceKline(context.Background(), &linear.QueryMarkPriceKlineParams{
			Symbol: "BTCUSDT",
			From:   time.Now().Add(-time.Hour).Second(),
		})
		assert.NoError(t, err)
		assert.Equal(t, 10001, response.RetCode)
		assert.NotEmpty(t, response)
		assert.Equal(t, "Param validation for 'interval' failed on the 'required' tag", response.RetMsg)
	})

	t.Run("Get Last Funding Rate -  LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().GetLastFundingRate(context.Background(), &linear.GetLastFundingRateParams{
			Symbol: "BTCUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.Equal(t, "OK", response.RetMsg)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSDT", response.Result.Symbol)
	})

	t.Run("Query Index Price Kline -  LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().QueryIndexPriceKline(context.Background(), &linear.QueryIndexPriceKlineParams{
			Symbol:   "BTCUSD",
			Interval: bybit.Interval1Min,
			From:     time.Now().Add(-time.Hour).Second(),
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.Equal(t, "OK", response.RetMsg)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSD", response.Result[0].Symbol)
	})

	t.Run("Query Index Price Kline - Param Validation fail -  LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().QueryIndexPriceKline(context.Background(), &linear.QueryIndexPriceKlineParams{
			Symbol:   "BTCUSDT",
			Interval: bybit.Interval1Min,
			From:     time.Now().Add(-time.Hour).Second(),
		})
		assert.NoError(t, err)
		assert.Equal(t, 10001, response.RetCode)
		assert.NotEmpty(t, response)
		assert.Equal(t, "Param validation for 'symbol' failed on the 'symbol' tag", response.RetMsg)
		assert.NotNil(t, response)
	})

	t.Run("Query Premium Index Kline -  LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().QueryPremiumIndexKline(context.Background(), &linear.QueryPremiumIndexKlineParams{
			Symbol:   "BTCUSD",
			Interval: bybit.Interval1Min,
			From:     time.Now().Add(-time.Hour).Second(),
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.Equal(t, "OK", response.RetMsg)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSD", response.Result[0].Symbol)
	})

	t.Run("Get Open Interest -  LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().GetOpenInterest(context.Background(), &linear.GetOpenInterestParams{
			Symbol: "BTCUSDT",
			Period: bybit.Period1hr,
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.Equal(t, "OK", response.RetMsg)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSDT", response.Result[0].Symbol)
	})

	t.Run("Get Latest Big Deal -  LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().GetLatestBigDeal(context.Background(), &linear.GetLatestBigDealParams{
			Symbol: "BTCUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.Equal(t, "OK", response.RetMsg)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSDT", response.Result[0].Symbol)
	})

	t.Run("Get Long Short Ratio -  LINEAR", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().GetLongShortRatio(context.Background(), &linear.GetLongShortRatioParams{
			Symbol: "BTCUSDT",
			Period: bybit.Period1hr,
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.Equal(t, "OK", response.RetMsg)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSDT", response.Result[0].Symbol)
	})
}
