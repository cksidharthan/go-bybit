package spot

import (
	"context"
	"encoding/json"
	"github.com/cksidharthan/go-bybit/helpers"
	"github.com/cksidharthan/go-bybit/public/domain/spot"
	"github.com/cksidharthan/go-bybit/public/spot/filters"
	"github.com/cksidharthan/go-bybit/transport"
	httpTransport "github.com/cksidharthan/go-bybit/transport/http"
	"net/http"
	"net/url"
)

type SpotClient struct {
	transporter transport.Transporter
}

func New(url string) SpotInterface {
	transporter := httpTransport.New(url, "", "")
	return &SpotClient{
		transporter: transporter,
	}
}

func (c *SpotClient) GetSymbols(ctx context.Context) (symbols *spot.SymbolsResponse, err error) {
	apiPath, err := url.Parse(helpers.PUBLIC_SPOT_SYMBOLS_PATH)
	if err != nil {
		return
	}

	payload, err := c.transporter.Call(ctx, apiPath, http.MethodGet, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &symbols)
	if err != nil {
		return
	}
	return
}

func (c *SpotClient) GetOrderBookDepth(ctx context.Context, filters *filters.OrderBookDepthFilter) (orderBook *spot.OrderBookResponse, err error) {
	apiPath, err := url.Parse(helpers.PUBLIC_SPOT_QUOTE_DEPTH_PATH)
	if err != nil {
		return
	}
	if filters != nil {
		filters.ToQuery(apiPath)
	}

	payload, err := c.transporter.Call(ctx, apiPath, http.MethodGet, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &orderBook)
	if err != nil {
		return
	}
	return
}

func (c *SpotClient) GetMergedOrderBook(ctx context.Context, filters *filters.MergedOrderBookFilter) (orderBook *spot.OrderBookResponse, err error) {
	apiPath, err := url.Parse(helpers.PUBLIC_SPOT_QUOTE_DEPTH_MERGED_PATH)
	if err != nil {
		return
	}
	if filters != nil {
		filters.ToQuery(apiPath)
	}

	payload, err := c.transporter.Call(ctx, apiPath, http.MethodGet, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &orderBook)
	if err != nil {
		return
	}
	return
}

func (c *SpotClient) GetPublicTradeRecords(ctx context.Context, filters *filters.TradeRecordsFilter) (publicTradeRecords *spot.TradeRecordsResponse, err error) {
	apiPath, err := url.Parse(helpers.PUBLIC_SPOT_QUOTE_TRADES_PATH)
	if err != nil {
		return
	}
	if filters != nil {
		filters.ToQuery(apiPath)
	}

	payload, err := c.transporter.Call(ctx, apiPath, http.MethodGet, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &publicTradeRecords)
	if err != nil {
		return
	}
	return
}

func (c *SpotClient) GetKline(ctx context.Context, filters *filters.KlineFilter) (kline *spot.KlineResponse, err error) {
	apiPath, err := url.Parse(helpers.PUBLIC_SPOT_QUOTE_KLINE_PATH)
	if err != nil {
		return
	}
	if filters != nil {
		filters.ToQuery(apiPath)
	}

	payload, err := c.transporter.Call(ctx, apiPath, http.MethodGet, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &kline)
	if err != nil {
		return
	}
	return
}

func (c *SpotClient) GetTickerInfo24hr(ctx context.Context, symbol string) (tickerInfo24hr *spot.TickerResponse, err error) {
	path := helpers.PUBLIC_SPOT_QUOTE_TICKER_24HR_PATH + "?symbol=" + symbol
	apiPath, err := url.Parse(path)
	if err != nil {
		return
	}
	payload, err := c.transporter.Call(ctx, apiPath, http.MethodGet, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &tickerInfo24hr)
	if err != nil {
		return
	}
	return
}

func (c *SpotClient) GetLastTradedPrice(ctx context.Context, symbol string) (lastTradedPrice *spot.PriceResponse, err error) {
	path := helpers.PUBLIC_SPOT_QUOTE_TICKER_PRICE + "?symbol=" + symbol
	apiPath, err := url.Parse(path)
	if err != nil {
		return
	}
	payload, err := c.transporter.Call(ctx, apiPath, http.MethodGet, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &lastTradedPrice)
	if err != nil {
		return
	}
	return
}

func (c *SpotClient) GetBidAskPrice(ctx context.Context, symbol string) (bidAskPrice *spot.PriceBidAskResponse, err error) {
	path := helpers.PUBLIC_SPOT_QUOTE_TICKER_BOOK_TICKER_PATH + "?symbol=" + symbol
	apiPath, err := url.Parse(path)
	if err != nil {
		return
	}
	payload, err := c.transporter.Call(ctx, apiPath, http.MethodGet, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &bidAskPrice)
	if err != nil {
		return
	}
	return
}
