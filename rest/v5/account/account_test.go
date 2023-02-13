package account_test

import (
	"context"
	goHttp "net/http"
	"testing"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/http"
	"github.com/cksidharthan/go-bybit/rest/v5/account"
	domain "github.com/cksidharthan/go-bybit/rest/v5/domain/account"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetWalletBalance(t *testing.T) {
	t.Parallel()

	t.Run("should return wallet balance", func(t *testing.T) {
		t.Parallel()

		accountClient := account.New(&http.ClientOpts{
			BaseURL:    bybit.APITestnetBaseURL,
			UserAgent:  "integration-test/go-bybit",
			HTTPClient: goHttp.DefaultClient,
		})

		response, err := accountClient.GetWalletBalance(context.Background(), &domain.GetWalletBalanceParams{
			AccountType: "spot",
			Coin:        "BTC",
		})

		assert.NoError(t, err)
		assert.NotNil(t, response)
	})
}
