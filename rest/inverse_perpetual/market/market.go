package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	inverseperp "github.com/cksidharthan/go-bybit/rest/domain/inverse_perpetual"
)

func (c *InversePerpetualMarketClient) OrderBook(ctx context.Context, params *inverseperp.OrderBookParams) (response *inverseperp.OrderBookResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualOrderBookPath, params, &response)
	if err != nil {
		return
	}
	return
}

func (c *InversePerpetualMarketClient) QueryKline(ctx context.Context, params *inverseperp.QueryKlineParams) (response *inverseperp.QueryKlineResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualQueryKlinePath, params, &response)
	if err != nil {
		return
	}
	return
}

func (c *InversePerpetualMarketClient) GetSymbolInformation(ctx context.Context, params *inverseperp.GetSymbolInformationParams) (response *inverseperp.GetSymbolInformationResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualSymbolInformationPath, params, &response)
	if err != nil {
		return
	}
	return
}

func (c *InversePerpetualMarketClient) PublicTradingRecords(ctx context.Context, params *inverseperp.PublicTradingRecordsParams) (response *inverseperp.PublicTradingRecordsResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualPublicTradingRecordsPath, params, &response)
	if err != nil {
		return
	}
	return
}

func (c *InversePerpetualMarketClient) QuerySymbols(ctx context.Context) (response *inverseperp.QuerySymbolsResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualQuerySymbolPath, nil, &response)
	if err != nil {
		return
	}
	return
}

func (c *InversePerpetualMarketClient) QueryMarkPriceKline(ctx context.Context, params *inverseperp.QueryMarkPriceKlineParams) (response *inverseperp.QueryMarkPriceKlineResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualQueryMarkPriceKlinePath, params, &response)
	if err != nil {
		return
	}
	return
}

func (c *InversePerpetualMarketClient) QueryIndexPriceKline(ctx context.Context, params *inverseperp.QueryIndexPriceKlineParams) (response *inverseperp.QueryIndexPriceKlineResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualQueryIndexPriceKlinePath, params, &response)
	if err != nil {
		return
	}
	return
}

func (c *InversePerpetualMarketClient) QueryPremiumIndexKline(ctx context.Context, params *inverseperp.QueryPremiumIndexKlineParams) (response *inverseperp.QueryPremiumIndexKlineResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualQueryPremiumIndexKlinePath, params, &response)
	if err != nil {
		return
	}
	return
}

func (c *InversePerpetualMarketClient) OpenInterest(ctx context.Context, params *inverseperp.OpenInterestParams) (response *inverseperp.OpenInterestResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualOrderBookPath, params, &response)
	if err != nil {
		return
	}
	return
}

func (c *InversePerpetualMarketClient) LatestBigDeal(ctx context.Context, params *inverseperp.LatestBigDealParams) (response *inverseperp.LatestBigDealResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualLatestBigDealPath, params, &response)
	if err != nil {
		return
	}
	return
}

func (c *InversePerpetualMarketClient) LongShortRatio(ctx context.Context, params *inverseperp.LongShortRatioParams) (response *inverseperp.LongShortRatioResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualLongShortRatioPath, params, &response)
	if err != nil {
		return
	}
	return
}
