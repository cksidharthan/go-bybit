package market

import (
	"context"

	"github.com/cksidharthan/go-bybit/rest/domain/spot/market"
	"github.com/cksidharthan/go-bybit/rest/filters"
)

type Interface interface {
	GetSymbols(ctx context.Context) (symbols *market.SymbolsResponse, err error)
	GetOrderBookDepth(ctx context.Context, filters *filters.OrderBookDepthFilter) (orderBook *market.OrderBookResponse, err error)
	GetMergedOrderBook(ctx context.Context, filters *filters.MergedOrderBookFilter) (mergedOrderBook *market.OrderBookResponse, err error)
	GetPublicTradeRecords(ctx context.Context, filters *filters.TradeRecordsFilter) (publicTradeRecords *market.TradeRecordsResponse, err error)
	GetKline(ctx context.Context, filters *filters.KlineFilter) (kline *market.KlineResponse, err error)
	GetTickerInfo24hr(ctx context.Context, symbol string) (tickerInfo24hr *market.TickerResponse, err error)
	GetLastTradedPrice(ctx context.Context, symbol string) (lastTradedPrice *market.PriceResponse, err error)
	GetBidAskPrice(ctx context.Context, symbol string) (bidAskPrice *market.PriceBidAskResponse, err error)
	GetServerTime(ctx context.Context) (serverTime *market.ServerTimeResponse, err error)
}
