package main

import (
	"context"
	"fmt"
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest"
	inverseperp "github.com/cksidharthan/go-bybit/rest/domain/inverse_perpetual"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
	"os"
)

func main() {
	client := rest.NewRestClient(bybit.BybitTestnetBaseURL, os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))

	// Get spot market client
	spotClient := client.Spot()
	// Get linear market client
	linearClient := client.Linear()
	// Get inverse perpetual market client
	inverseClient := client.InversePerpetual()

	spotResponse, err := spotClient.Market().GetSymbols(context.Background())
	if err != nil {
		_ = fmt.Errorf("error: %v", err)
	}

	linearResponse, err := linearClient.Market().GetSymbolInformation(context.Background(), &linear.GetSymbolInformationParams{
		Symbol: "BTCUSDT",
	})
	if err != nil {
		_ = fmt.Errorf("error: %v", err)
	}

	inverseResponse, err := inverseClient.Market().GetSymbolInformation(context.Background(), &inverseperp.GetSymbolInformationParams{
		Symbol: "BTCUSD",
	})
	if err != nil {
		_ = fmt.Errorf("error: %v", err)
	}

	fmt.Println(spotResponse)
	fmt.Println(linearResponse)
	fmt.Println(inverseResponse)
}
