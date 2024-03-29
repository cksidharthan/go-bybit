[![Code check](https://github.com/cksidharthan/go-bybit/actions/workflows/ci.yml/badge.svg)](https://github.com/cksidharthan/go-bybit/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/cksidharthan/go-bybit)](https://goreportcard.com/report/github.com/cksidharthan/go-bybit)

# go-bybit

[Bybit](https://bybit.com) API Client SDK for [Golang (Go)](https://golang.org/).
This library is a wrapper SDK for the [Bybit](https://bybit.com) API written in Golang(Go).

Bybit API Docs: [https://bybit.com/docs](https://bybit-exchange.github.io/docs/spot/#t-introduction)

This Project is under active development. Please feel free to contribute.

## Install

```bash
go get -u github.com/cksidharthan/go-bybit
```

## Example
```go
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
		fmt.Errorf("error: %v", err)
	}

	linearResponse, err := linearClient.Market().GetSymbolInformation(context.Background(), &linear.GetSymbolInformationParams{
		Symbol: "BTCUSDT",
	})
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	inverseResponse, err := inverseClient.Market().GetSymbolInformation(context.Background(), &inverseperp.GetSymbolInformationParams{
		Symbol: "BTCUSD",
	})
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	fmt.Println(spotResponse)
	fmt.Println(linearResponse)
	fmt.Println(inverseResponse)
}
```

All the tests in this repository interact with the Bybit API directly. Please refer them for examples, when using this library. :) 

In the meantime, I'll be adding more examples to `./examples` folder.

# Endpoints completed

| Category          | SubCategory                                                                            |      Created       |     Testcases      |
|-------------------|----------------------------------------------------------------------------------------|:------------------:|:------------------:|
| SPOT              | [Market Data](https://bybit-exchange.github.io/docs/spot/#t-marketdata)                | :heavy_check_mark: | :heavy_check_mark: |
| SPOT              | [Account Data](https://bybit-exchange.github.io/docs/spot/#t-accountdata)              | :heavy_check_mark: | :heavy_check_mark: |
| SPOT              | [Wallet Data](https://bybit-exchange.github.io/docs/spot/#t-wallet)                    | :heavy_check_mark: | :heavy_check_mark: |
| SPOT              | [API Data](https://bybit-exchange.github.io/docs/spot/#t-api)                          | :heavy_check_mark: | :heavy_check_mark: |
| USDT Perpetual    | [Market Data](https://bybit-exchange.github.io/docs/futuresV2/linear/#t-marketdata)    | :heavy_check_mark: | :heavy_check_mark: |
| USDT Perpetual    | [Account Data](https://bybit-exchange.github.io/docs/futuresV2/linear/#t-accountdata)  | :heavy_check_mark: | :heavy_check_mark: |
| USDT Perpetual    | [Wallet Data](https://bybit-exchange.github.io/docs/futuresV2/linear/#t-wallet)        | :heavy_check_mark: | :heavy_check_mark: |
| Inverse Perpetual | [Market Data](https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-marketdata)   | :heavy_check_mark: | :heavy_check_mark: |
| Inverse Perpetual | [Account Data](https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-accountdata) | :heavy_check_mark: | :heavy_check_mark: |
| Inverse Perpetual | [Wallet Data](https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-wallet)       | :heavy_check_mark: | :heavy_check_mark: |

### Contributing

Pull requests are welcome.

Please make sure to add tests when adding new methods.

## Local Development
Code Formatting and linting.

```bash
make fmt
make lint
```

## Unit Testing

```bash
make test
```
![visitors](https://visitor-badge.glitch.me/badge?page_id=cksidharthan.go-bybit)
