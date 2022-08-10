package spot

import (
	"context"
)

type MarketInterface interface {
	GetSymbols(ctx context.Context) (symbols *GetSymbolsResponse, err error)
	GetOrderBookDepth(ctx context.Context, params *OrderBookDepthParams) (response *OrderBookDepthResponse, err error)
	GetMergedOrderBook(ctx context.Context, params *MergedOrderBookParams) (response *MergedOrderBookResponse, err error)
	GetPublicTradeRecords(ctx context.Context, params *PublicTradeRecordsParams) (response *PublicTradeRecordsResponse, err error)
	GetKline(ctx context.Context, params *GetKlineParams) (response *KlineResponse, err error)
	GetTickerInfo24hr(ctx context.Context, params *GetTickerInfo24hrParams) (response *GetTickerInfoResponse, err error)
	GetLastTradedPrice(ctx context.Context, params *GetLastTradedPriceParams) (response *GetLastTradedPriceResponse, err error)
	GetBidAskPrice(ctx context.Context, params *GetBidAskPriceParams) (response *BidAskPriceResponse, err error)
	GetServerTime(ctx context.Context) (response *ServerTimeResponse, err error)
}
