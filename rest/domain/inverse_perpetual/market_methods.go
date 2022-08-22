package inverseperp

import (
	"context"
)

type MarketInterface interface {
	OrderBook(ctx context.Context, params *OrderBookParams) (*OrderBookResponse, error)
	QueryKline(ctx context.Context, params *QueryKlineParams) (*QueryKlineResponse, error)
	GetSymbolInformation(ctx context.Context, params *GetSymbolInformationParams) (*GetSymbolInformationResponse, error)
	PublicTradingRecords(ctx context.Context, params *PublicTradingRecordsParams) (*PublicTradingRecordsResponse, error)
	QuerySymbols(ctx context.Context) (*QuerySymbolsResponse, error)
	QueryMarkPriceKline(ctx context.Context, params *QueryMarkPriceKlineParams) (*QueryMarkPriceKlineResponse, error)
	QueryIndexPriceKline(ctx context.Context, params *QueryIndexPriceKlineParams) (*QueryIndexPriceKlineResponse, error)
	QueryPremiumIndexKline(ctx context.Context, params *QueryPremiumIndexKlineParams) (*QueryPremiumIndexKlineResponse, error)
	OpenInterest(ctx context.Context, params *OpenInterestParams) (*OpenInterestResponse, error)
	LatestBigDeal(ctx context.Context, params *LatestBigDealParams) (*LatestBigDealResponse, error)
	LongShortRatio(ctx context.Context, params *LongShortRatioParams) (*LongShortRatioResponse, error)
}
