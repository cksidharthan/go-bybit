package account_test

import (
	"context"
	"os"
	"testing"

	"github.com/cksidharthan/go-bybit/bybit"
	inverseperp "github.com/cksidharthan/go-bybit/rest/domain/inverse_perpetual"
	"github.com/cksidharthan/go-bybit/rest/inverse_perpetual/account"
	"github.com/stretchr/testify/assert"
)

func TestClientInversePerpetual_AccountFunding(t *testing.T) {
	t.Parallel()

	bybitClient := account.NewInversePerpetualAccountClient(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	t.Run("Get Funding Rate - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.GetFundingRate(context.Background(), &inverseperp.GetFundingRateParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, "BTCUSD", response.Result.Symbol)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})

	t.Run("Get Last Funding Fee - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.GetLastFundingFee(context.Background(), &inverseperp.GetLastFundingFeeParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})

	t.Run("Get Predicted Funding - INVERSE PERPETUAL", func(t *testing.T) {
		t.Parallel()
		response, err := bybitClient.GetPredictedFunding(context.Background(), &inverseperp.GetPredictedFundingParams{
			Symbol: "BTCUSD",
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, response.RetCode)
		assert.Equal(t, "OK", response.RetMsg)
	})
}
