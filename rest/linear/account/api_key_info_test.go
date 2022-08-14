package account_test

import (
	"context"
	"os"
	"testing"

	"github.com/cksidharthan/go-bybit/bybit"
	linearRest "github.com/cksidharthan/go-bybit/rest/linear"
	"github.com/stretchr/testify/assert"
)

func TestClient_Linear_Account_ApiKey(t *testing.T) {
	bybitClient := linearRest.NewLinearClient(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	t.Run("Get API Key Info - LINEAR", func(t *testing.T) {
		response, err := bybitClient.Account().GetAPIKeyInfo(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 0, response.RetCode)
		assert.NotEmpty(t, response)
		assert.NotNil(t, response)
	})
}
