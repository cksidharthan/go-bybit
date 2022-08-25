package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

// GetSymbols - returns a list of all symbols with information about each symbol. [/spot/v1/symbols - GET]
//
// docs - https://bybit-exchange.github.io/docs/spot/#t-spot_querysymbol
func (c *SpotMarketClient) GetSymbols(ctx context.Context) (symbols *spot.GetSymbolsResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotSymbolsPath, nil, &symbols)
	if err != nil {
		return
	}
	return
}

// GetOrderBookDepth - returns the orderbook for a given symbol. [/spot/quote/v1/depth - GET]
//
// docs  - https://bybit-exchange.github.io/docs/spot/#t-orderbook
func (c *SpotMarketClient) GetOrderBookDepth(ctx context.Context, params *spot.OrderBookDepthParams) (res *spot.OrderBookDepthResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteDepthPath, params, &res)
	if err != nil {
		return
	}
	return
}

// GetMergedOrderBook - returns the orderbook for a given symbol. [/spot/quote/v1/depth/merged - GET]
//
// docs  - https://bybit-exchange.github.io/docs/spot/#t-mergedorderbook
func (c *SpotMarketClient) GetMergedOrderBook(ctx context.Context, params *spot.MergedOrderBookParams) (res *spot.MergedOrderBookResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteDepthMergedPath, params, &res)
	if err != nil {
		return
	}
	return
}

// GetPublicTradeRecords - returns a list of all public trades for a given symbol. [/spot/quote/v1/trades - GET]
//
// docs  - https://bybit-exchange.github.io/docs/spot/#t-publictradingrecords
func (c *SpotMarketClient) GetPublicTradeRecords(ctx context.Context, params *spot.PublicTradeRecordsParams) (res *spot.PublicTradeRecordsResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteTradesPath, params, &res)
	if err != nil {
		return
	}
	return
}

// GetKline - returns a list of klines for a given symbol. [/spot/quote/v1/kline - GET]
//
// docs  - https://bybit-exchange.github.io/docs/spot/#t-querykline
func (c *SpotMarketClient) GetKline(ctx context.Context, params *spot.GetKlineParams) (res *spot.KlineResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteKlinePath, params, &res)
	if err != nil {
		return
	}
	return
}

// GetTickerInfo24hr - returns 24hr ticker info for a given symbol. [/spot/quote/v1/ticker/24hr - GET]
//
// docs  - https://bybit-exchange.github.io/docs/spot/#t-spot_latestsymbolinfo
func (c *SpotMarketClient) GetTickerInfo24hr(ctx context.Context, params *spot.GetTickerInfo24hrParams) (res *spot.GetTickerInfoResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteTicker24HrPath, params, &res)
	if err != nil {
		return
	}
	return
}

// GetLastTradedPrice - returns last trade price for a given symbol. [/spot/quote/v1/ticker/price - GET]
//
// docs  - https://bybit-exchange.github.io/docs/spot/#t-lasttradedprice
func (c *SpotMarketClient) GetLastTradedPrice(ctx context.Context, params *spot.GetLastTradedPriceParams) (res *spot.GetLastTradedPriceResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteTickerPrice, params, &res)
	if err != nil {
		return
	}
	return
}

// GetBidAskPrice - returns bid ask price for a given symbol. [/spot/quote/v1/ticker/book_ticker - GET]
//
// docs  - https://bybit-exchange.github.io/docs/spot/#t-bestbidask
func (c *SpotMarketClient) GetBidAskPrice(ctx context.Context, params *spot.GetBidAskPriceParams) (res *spot.BidAskPriceResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteTickerBookTickerPath, params, &res)
	if err != nil {
		return
	}
	return
}

// GetServerTime - returns server time. [/spot/v1/time - GET]
//
// docs  - https://bybit-exchange.github.io/docs/spot/#t-api
func (c *SpotMarketClient) GetServerTime(ctx context.Context) (res *spot.ServerTimeResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicServerTimePath, nil, &res)
	if err != nil {
		return
	}
	return
}
