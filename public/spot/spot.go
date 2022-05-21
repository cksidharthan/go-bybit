package spot

import (
	"context"
	"github.com/cksidharthan/go-bybit/public/domain/spot"
	"github.com/cksidharthan/go-bybit/public/spot/filters"
)

type SpotInterface interface {
	GetSymbols(ctx context.Context) (symbols *spot.SymbolsResponse, err error)
	GetOrderBookDepth(ctx context.Context, filters *filters.OrderBookDepthFilter) (orderBook *spot.OrderBookResponse, err error)
	GetMergedOrderBook(ctx context.Context, filters *filters.MergedOrderBookFilter) (mergedOrderBook *spot.OrderBookResponse, err error)
	GetPublicTradeRecords(ctx context.Context, filters *filters.TradeRecordsFilter) (publicTradeRecords *spot.TradeRecordsResponse, err error)
	GetKline(ctx context.Context, filters *filters.KlineFilter) (kline *spot.KlineResponse, err error)
	GetTickerInfo24hr(ctx context.Context, symbol string) (tickerInfo24hr *spot.TickerResponse, err error)
	GetLastTradedPrice(ctx context.Context, symbol string) (lastTradedPrice *spot.PriceResponse, err error)
	GetBidAskPrice(ctx context.Context, symbol string) (bidAskPrice *spot.PriceBidAskResponse, err error)
}
