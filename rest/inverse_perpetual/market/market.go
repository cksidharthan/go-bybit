package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	inverseperp "github.com/cksidharthan/go-bybit/rest/domain/inverse_perpetual"
)

// OrderBook - get order book. [ /v2/public/orderBook/L2  - GET ]
//
// Get the orderbook. Each side has a depth of 25.
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-orderbook
func (c *InversePerpetualMarketClient) OrderBook(ctx context.Context, params *inverseperp.OrderBookParams) (response *inverseperp.OrderBookResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualOrderBookPath, params, &response)
	if err != nil {
		return
	}
	return
}

// QueryKline - get kline data. [ /v2/public/kline/list  - GET ]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-querykline
func (c *InversePerpetualMarketClient) QueryKline(ctx context.Context, params *inverseperp.QueryKlineParams) (response *inverseperp.QueryKlineResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualQueryKlinePath, params, &response)
	if err != nil {
		return
	}
	return
}

// GetSymbolInformation - get latest symbol information. [ /v2/public/tickers  - GET ]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-latestsymbolinfo
func (c *InversePerpetualMarketClient) GetSymbolInformation(ctx context.Context, params *inverseperp.GetSymbolInformationParams) (response *inverseperp.GetSymbolInformationResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualSymbolInformationPath, params, &response)
	if err != nil {
		return
	}
	return
}

// PublicTradingRecords - get public trading records. [ /v2/public/trading-records  - GET ]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-publictradingrecords
func (c *InversePerpetualMarketClient) PublicTradingRecords(ctx context.Context, params *inverseperp.PublicTradingRecordsParams) (response *inverseperp.PublicTradingRecordsResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualPublicTradingRecordsPath, params, &response)
	if err != nil {
		return
	}
	return
}

// QuerySymbols - get symbols. [ /v2/public/symbols  - GET ]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-querysymbol
func (c *InversePerpetualMarketClient) QuerySymbols(ctx context.Context) (response *inverseperp.QuerySymbolsResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualQuerySymbolPath, nil, &response)
	if err != nil {
		return
	}
	return
}

// QueryMarkPriceKline - Query mark price kline (like Query Kline but for mark price). [ /v2/public/mark-price-kline - GET ]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-markpricekline
func (c *InversePerpetualMarketClient) QueryMarkPriceKline(ctx context.Context, params *inverseperp.QueryMarkPriceKlineParams) (response *inverseperp.QueryMarkPriceKlineResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualQueryMarkPriceKlinePath, params, &response)
	if err != nil {
		return
	}
	return
}

// QueryIndexPriceKline - Index price kline. Tracks BTC spot prices, with a frequency of every second [ /v2/public/index-price-kline - GET ]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-queryindexpricekline
func (c *InversePerpetualMarketClient) QueryIndexPriceKline(ctx context.Context, params *inverseperp.QueryIndexPriceKlineParams) (response *inverseperp.QueryIndexPriceKlineResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualQueryIndexPriceKlinePath, params, &response)
	if err != nil {
		return
	}
	return
}

// QueryPremiumIndexKline - Query premium index kline. [ /v2/public/premium-index-kline - GET ]
//
// Premium index kline. Tracks the premium / discount of BTC perpetual contracts relative to the mark price per minute
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-querypremiumindexkline
func (c *InversePerpetualMarketClient) QueryPremiumIndexKline(ctx context.Context, params *inverseperp.QueryPremiumIndexKlineParams) (response *inverseperp.QueryPremiumIndexKlineResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualQueryPremiumIndexKlinePath, params, &response)
	if err != nil {
		return
	}
	return
}

// OpenInterest - get open interest. [ /v2/public/open-interest  - GET ]
//
// Gets the total amount of unsettled contracts. In other words, the total number of contracts held in open positions.
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-marketopeninterest
func (c *InversePerpetualMarketClient) OpenInterest(ctx context.Context, params *inverseperp.OpenInterestParams) (response *inverseperp.OpenInterestResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualOpenInterestPath, params, &response)
	if err != nil {
		return
	}
	return
}

// LatestBigDeal - get latest big deal. [ /v2/public/big-deal  - GET ]
//
// Obtain filled orders worth more than 500,000 USD within the last 24h.
//
// This endpoint may return orders which are over the maximum order qty for the symbol you call. For instance, the maximum order qty for BTCUSD is 1 million contracts, but in the event of the liquidation of a position larger than 1 million this endpoint returns this "impossible" order size.
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-marketbigdeal
func (c *InversePerpetualMarketClient) LatestBigDeal(ctx context.Context, params *inverseperp.LatestBigDealParams) (response *inverseperp.LatestBigDealResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualLatestBigDealPath, params, &response)
	if err != nil {
		return
	}
	return
}

// LongShortRatio - get long short ratio. [ /v2/public/account-ratio  - GET ]
//
// docs - https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-marketaccountratio
func (c *InversePerpetualMarketClient) LongShortRatio(ctx context.Context, params *inverseperp.LongShortRatioParams) (response *inverseperp.LongShortRatioResponse, err error) {
	err = c.Transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicInversePerpetualLongShortRatioPath, params, &response)
	if err != nil {
		return
	}
	return
}
