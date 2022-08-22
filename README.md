[![Code check](https://github.com/cksidharthan/go-bybit/actions/workflows/ci.yml/badge.svg)](https://github.com/cksidharthan/go-bybit/actions/workflows/ci.yml)

# go-bybit

[Bybit](https://bybit.com) API Client for [Golang (Go)](https://golang.org/).

This Project is under active development. Please feel free to contribute.

## Install

```bash
go get -u github.com/cksidharthan/go-bybit
```

# Enpoints completed

| Category          | SubCategory  |      Created       |     Testcases      |
|-------------------|--------------|:------------------:|:------------------:|
| SPOT              | Market Data  | :heavy_check_mark: | :heavy_check_mark: |
| SPOT              | Account Data | :heavy_check_mark: | :heavy_check_mark: |
| SPOT              | Wallet Data  | :heavy_check_mark: | :heavy_check_mark: |
| SPOT              | API Data     | :heavy_check_mark: | :heavy_check_mark: |
| USDT Perpetual    | Market Data  | :heavy_check_mark: | :heavy_check_mark: |
| USDT Perpetual    | Account Data | :heavy_check_mark: | :heavy_check_mark: |
| USDT Perpetual    | Wallet Data  | :heavy_check_mark: | :heavy_check_mark: |
| Inverse Perpetual | Market Data  | :heavy_check_mark: |   :construction:   |
| Inverse Perpetual | Account Data |   :construction:   |   :construction:   |
| Inverse Perpetual | Wallet Data  |   :construction:   |   :construction:   |

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
