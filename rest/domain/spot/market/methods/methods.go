package methods

import (
	"context"

	"github.com/cksidharthan/go-bybit/rest/domain/spot/market/filters"
	"github.com/cksidharthan/go-bybit/rest/domain/spot/market/types"
)

type Interface interface {
	GetSymbols(ctx context.Context) (symbols *types.SymbolsResponse, err error)
	GetOrderBookDepth(ctx context.Context, filters *filters.OrderBookDepthFilter) (orderBook *types.OrderBookResponse, err error)
	GetMergedOrderBook(ctx context.Context, filters *filters.MergedOrderBookFilter) (mergedOrderBook *types.OrderBookResponse, err error)
	GetPublicTradeRecords(ctx context.Context, filters *filters.TradeRecordsFilter) (publicTradeRecords *types.TradeRecordsResponse, err error)
	GetKline(ctx context.Context, filters *filters.KlineFilter) (kline *types.KlineResponse, err error)
	GetTickerInfo24hr(ctx context.Context, symbol string) (tickerInfo24hr *types.TickerResponse, err error)
	GetLastTradedPrice(ctx context.Context, symbol string) (lastTradedPrice *types.PriceResponse, err error)
	GetBidAskPrice(ctx context.Context, symbol string) (bidAskPrice *types.PriceBidAskResponse, err error)
	GetServerTime(ctx context.Context) (serverTime *types.ServerTimeResponse, err error)
}
