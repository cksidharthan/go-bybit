package linear

import "context"

type MarketInterface interface {
	GetOrderBook(ctx context.Context, params *OrderBookParams) (*OrderBookResponse, error)
	QueryKline(ctx context.Context, params *QueryKlineParams) (*QueryKlineResponse, error)
	GetSymbolInformation(ctx context.Context, params *GetSymbolInformationParams) (*GetSymbolInformationResponse, error)
	GetPublicTradingRecords(ctx context.Context, params *GetPublicTradingRecordsParams) (*GetPublicTradingRecordsResponse, error)
	QuerySymbol(ctx context.Context) (*QuerySymbolResponse, error)
	GetLiquidatedOrders(ctx context.Context, params *GetLiquidatedOrdersParams) (*GetLiquidatedOrdersResponse, error)
	GetLastFundingRate(ctx context.Context, params *GetLastFundingRateParams) (*GetLastFundingRateResponse, error)
	QueryMarkPriceKline(ctx context.Context, params *QueryMarkPriceKlineParams) (*QueryMarkPriceKlineResponse, error)
	QueryIndexPriceKline(ctx context.Context, params *QueryIndexPriceKlineParams) (*QueryIndexPriceKlineResponse, error)
	QueryPremiumIndexKline(ctx context.Context, params *QueryPremiumIndexKlineParams) (*QueryPremiumIndexKlineResponse, error)
	GetOpenInterest(ctx context.Context, params *GetOpenInterestParams) (*GetOpenInterestResponse, error)
	GetLatestBigDeal(ctx context.Context, params *GetLatestBigDealParams) (*GetLatestBigDealResponse, error)
	GetLongShortRatio(ctx context.Context, params *GetLongShortRatioParams) (*GetLongShortRatioResponse, error)
}
