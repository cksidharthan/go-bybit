package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
)

func (c *LinearMarketClient) GetOrderBook(ctx context.Context, params *linear.OrderBookParams) (res *linear.OrderBookResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicLinearOrderBookPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearMarketClient) QueryKline(ctx context.Context, params *linear.QueryKlineParams) (res *linear.QueryKlineResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicLinearQueryKlinePath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearMarketClient) GetSymbolInformation(ctx context.Context, params *linear.GetSymbolInformationParams) (res *linear.GetSymbolInformationResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicLinearGetSymbolInformationPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearMarketClient) GetPublicTradingRecords(ctx context.Context, params *linear.GetPublicTradingRecordsParams) (res *linear.GetPublicTradingRecordsResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicLinearGetPublicTradingRecordsPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearMarketClient) QuerySymbol(ctx context.Context) (res *linear.QuerySymbolResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicLinearQuerySymbolPath, nil, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearMarketClient) GetLiquidatedOrders(ctx context.Context, params *linear.GetLiquidatedOrdersParams) (res *linear.GetLiquidatedOrdersResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicLinearLiquidatedOrdersPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearMarketClient) GetLastFundingRate(ctx context.Context, params *linear.GetLastFundingRateParams) (res *linear.GetLastFundingRateResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicLinearGetLastFundingRatePath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearMarketClient) QueryMarkPriceKline(ctx context.Context, params *linear.QueryMarkPriceKlineParams) (res *linear.QueryMarkPriceKlineResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicLinearQueryMarkPriceKline, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearMarketClient) QueryIndexPriceKline(ctx context.Context, params *linear.QueryIndexPriceKlineParams) (res *linear.QueryIndexPriceKlineResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicLinearQueryIndexPriceKlinePath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearMarketClient) QueryPremiumIndexKline(ctx context.Context, params *linear.QueryPremiumIndexKlineParams) (res *linear.QueryPremiumIndexKlineResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicLinearQueryPremiumIndexPriceKlinePath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearMarketClient) GetOpenInterest(ctx context.Context, params *linear.GetOpenInterestParams) (res *linear.GetOpenInterestResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicLinearGetOpenInterestPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearMarketClient) GetLatestBigDeal(ctx context.Context, params *linear.GetLatestBigDealParams) (res *linear.GetLatestBigDealResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicLinearGetLatestBigDealPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *LinearMarketClient) GetLongShortRatio(ctx context.Context, params *linear.GetLongShortRatioParams) (res *linear.GetLongShortRatioResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicLinearGetLongShortRatioPath, params, &res)
	if err != nil {
		return
	}
	return
}
