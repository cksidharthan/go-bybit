package market

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

func (c *Client) GetSymbols(ctx context.Context) (symbols *spot.GetSymbolsResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotSymbolsPath, nil, &symbols)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetOrderBookDepth(ctx context.Context, params *spot.OrderBookDepthParams) (res *spot.OrderBookDepthResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteDepthPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetMergedOrderBook(ctx context.Context, params *spot.MergedOrderBookParams) (res *spot.MergedOrderBookResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteDepthMergedPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetPublicTradeRecords(ctx context.Context, params *spot.PublicTradeRecordsParams) (res *spot.PublicTradeRecordsResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteTradesPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetKline(ctx context.Context, params *spot.GetKlineParams) (res *spot.KlineResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteKlinePath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetTickerInfo24hr(ctx context.Context, params *spot.GetTickerInfo24hrParams) (res *spot.GetTickerInfoResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteTicker24HrPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetLastTradedPrice(ctx context.Context, params *spot.GetLastTradedPriceParams) (res *spot.GetLastTradedPriceResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteTickerPrice, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetBidAskPrice(ctx context.Context, params *spot.GetBidAskPriceParams) (res *spot.BidAskPriceResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicSpotQuoteTickerBookTickerPath, params, &res)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetServerTime(ctx context.Context) (res *spot.ServerTimeResponse, err error) {
	err = c.transporter.UnsignedRequest(ctx, http.MethodGet, bybit.PublicServerTimePath, nil, &res)
	if err != nil {
		return
	}
	return
}
