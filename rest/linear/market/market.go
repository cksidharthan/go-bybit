package market

import (
	"context"
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
	"net/http"
)

func (c *Client) GetOrderBook(ctx context.Context, params *linear.OrderBookParams) (res *linear.OrderBookResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicLinearOrderBookPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) QueryKline(ctx context.Context, params *linear.QueryKlineParams) (res *linear.QueryKlineResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicQueryKlinePath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetSymbolInformation(ctx context.Context, params *linear.GetSymbolInformationParams) (res *linear.GetSymbolInformationResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicGetSymbolInformationPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetPublicTradingRecords(ctx context.Context, params *linear.GetPublicTradingRecordsParams) (res *linear.GetPublicTradingRecordsResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicGetPublicTradingRecordsPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) QuerySymbol(ctx context.Context) (res *linear.QuerySymbolResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicQuerySymbolPath, nil, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetLiquidatedOrders(ctx context.Context, params *linear.GetLiquidatedOrdersParams) (res *linear.GetLiquidatedOrdersResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicLiquidatedOrdersPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetLastFundingRate(ctx context.Context, params *linear.GetLastFundingRateParams) (res *linear.GetLastFundingRateResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicGetLastFundingRatePath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) QueryMarkPriceKline(ctx context.Context, params *linear.QueryMarkPriceKlineParams) (res *linear.QueryMarkPriceKlineResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicQueryMarkPriceKline, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) QueryIndexPriceKline(ctx context.Context, params *linear.QueryIndexPriceKlineParams) (res *linear.QueryIndexPriceKlineResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicQueryIndexPriceKlinePath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) QueryPremiumIndexKline(ctx context.Context, params *linear.QueryPremiumIndexKlineParams) (res *linear.QueryPremiumIndexKlineResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicQueryPremiumIndexPriceKlinePath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetOpenInterest(ctx context.Context, params *linear.GetOpenInterestParams) (res *linear.GetOpenInterestResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicGetOpenInterestPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetLatestBigDeal(ctx context.Context, params *linear.GetLatestBigDealParams) (res *linear.GetLatestBigDealResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicGetLatestBigDealPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetLongShortRatio(ctx context.Context, params *linear.GetLongShortRatioParams) (res *linear.GetLongShortRatioResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicGetLongShortRatioPath, params, &res)
	if err != nil {
		return
	}
	return
}
