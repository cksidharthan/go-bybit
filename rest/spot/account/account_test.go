package account_test

import (
	"os"
	"testing"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/spot"
)

func TestClient_Account(t *testing.T) {
	t.Parallel()

	_ = spot.New(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))
}
