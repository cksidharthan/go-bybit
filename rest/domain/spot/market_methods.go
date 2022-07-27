package spot

import (
	"context"
)

type MarketInterface interface {
	GetSymbols(ctx context.Context) (symbols *SymbolsResponse, err error)
	GetOrderBookDepth(ctx context.Context, filters *OrderBookDepthFilter) (orderBook *OrderBookResponse, err error)
	GetMergedOrderBook(ctx context.Context, filters *MergedOrderBookFilter) (mergedOrderBook *OrderBookResponse, err error)
	GetPublicTradeRecords(ctx context.Context, filters *TradeRecordsFilter) (publicTradeRecords *TradeRecordsResponse, err error)
	GetKline(ctx context.Context, filters *KlineFilter) (kline *KlineResponse, err error)
	GetTickerInfo24hr(ctx context.Context, symbol string) (tickerInfo24hr *TickerResponse, err error)
	GetLastTradedPrice(ctx context.Context, symbol string) (lastTradedPrice *PriceResponse, err error)
	GetBidAskPrice(ctx context.Context, symbol string) (bidAskPrice *PriceBidAskResponse, err error)
	GetServerTime(ctx context.Context) (serverTime *ServerTimeResponse, err error)
}
