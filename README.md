[![Code check](https://github.com/cksidharthan/go-bybit/actions/workflows/ci.yml/badge.svg)](https://github.com/cksidharthan/go-bybit/actions/workflows/ci.yml)

# go-bybit

Unofficial [Bybit](https://bybit.com) API Client  for [Go](https://golang.org/).

This Project is under active development and is not production ready.

## Install

```bash
go get -u github.com/cksidharthan/go-bybit
```

# Enpoints completed

| Category          | SubCategory  |      Created       |       Tested       |
|-------------------|--------------|:------------------:|:------------------:|
| SPOT              | Market Data  | :heavy_check_mark: | :heavy_check_mark: |
| SPOT              | Account Data | :heavy_check_mark: | :heavy_check_mark: |
| SPOT              | Wallet Data  | :heavy_check_mark: | :heavy_check_mark: |
| SPOT              | API Data     | :heavy_check_mark: | :heavy_check_mark: |
| USDT Perpetual    | Market Data  |        :x:         |        :x:         |
| USDT Perpetual    | Account Data |        :x:         |        :x:         |
| USDT Perpetual    | Wallet Data  |        :x:         |        :x:         |
| USDT Perpetual    | API Data     |        :x:         |        :x:         |
| Inverse Perpetual | Market Data  |        :x:         |        :x:         |
| Inverse Perpetual | Account Data |        :x:         |        :x:         |
| Inverse Perpetual | Wallet Data  |        :x:         |        :x:         |
| Inverse Perpetual | API Data     |        :x:         |        :x:         |

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