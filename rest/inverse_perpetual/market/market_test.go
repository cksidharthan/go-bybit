package market_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cksidharthan/go-bybit/bybit"
	domain "github.com/cksidharthan/go-bybit/rest/domain/inverse_perpetual"
	inverseperp "github.com/cksidharthan/go-bybit/rest/inverse_perpetual"
	"github.com/stretchr/testify/assert"
)

func TestClient_InversePerpetual_Market(t *testing.T) {
	t.Parallel()
	bybitClient := inverseperp.NewInversePerpetualClient(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	t.Run("Get Order Book - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().OrderBook(context.Background(), &domain.OrderBookParams{
			Symbol: "BTCUSDT",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSDT", response.Result[0].Symbol)
	})

	t.Run("Query Kline - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().QueryKline(context.Background(), &domain.QueryKlineParams{
			Symbol:   "BTCUSD",
			Interval: string(bybit.Interval1Min),
			From:     time.Now().Add(-4 * time.Hour).Unix(),
			Limit:    3,
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSD", response.Result[0].Symbol)
		assert.Equal(t, 3, len(response.Result))
	})

	t.Run("Get Symbol Info - Single Symbol - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().GetSymbolInformation(context.Background(), &domain.GetSymbolInformationParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSD", response.Result[0].Symbol)
		assert.Equal(t, 1, len(response.Result))
	})

	t.Run("Get Symbol Info - Multiple Symbols - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().GetSymbolInformation(context.Background(), &domain.GetSymbolInformationParams{})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Public Trading Records - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().PublicTradingRecords(context.Background(), &domain.PublicTradingRecordsParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSD", response.Result[0].Symbol)
	})

	t.Run("Query Symbols - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().QuerySymbols(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})

	t.Run("Query Mark Price Kline - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().QueryMarkPriceKline(context.Background(), &domain.QueryMarkPriceKlineParams{
			Symbol:   "BTCUSD",
			Interval: string(bybit.Interval1Min),
			From:     time.Now().Add(-4 * time.Hour).Unix(),
			Limit:    3,
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSD", response.Result[0].Symbol)
		assert.Equal(t, 3, len(response.Result))
	})

	t.Run("Query Index Price Kline - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().QueryIndexPriceKline(context.Background(), &domain.QueryIndexPriceKlineParams{
			Symbol:   "BTCUSD",
			Interval: string(bybit.Interval1Min),
			From:     time.Now().Add(-4 * time.Hour).Unix(),
			Limit:    3,
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSD", response.Result[0].Symbol)
		assert.Equal(t, 3, len(response.Result))
	})

	t.Run("Query Premium Index Price Kline - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().QueryPremiumIndexKline(context.Background(), &domain.QueryPremiumIndexKlineParams{
			Symbol:   "BTCUSD",
			Interval: string(bybit.Interval1Min),
			From:     time.Now().Add(-4 * time.Hour).Unix(),
			Limit:    3,
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSD", response.Result[0].Symbol)
		assert.Equal(t, 3, len(response.Result))
	})

	t.Run("Open Interest - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().OpenInterest(context.Background(), &domain.OpenInterestParams{
			Symbol: "BTCUSD",
			Limit:  50,
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSD", response.Result[0].Symbol)
		assert.Equal(t, 50, len(response.Result))
	})

	t.Run("Latest Big Deal - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().LatestBigDeal(context.Background(), &domain.LatestBigDealParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSD", response.Result[0].Symbol)
	})

	t.Run("Long Short Ratio - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.Market().LongShortRatio(context.Background(), &domain.LongShortRatioParams{
			Symbol: "BTCUSD",
			Period: "5min",
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSD", response.Result[0].Symbol)
	})
}
