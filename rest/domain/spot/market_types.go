package spot

import "github.com/cksidharthan/go-bybit/rest/domain"

type GetSymbolsResult struct {
	Name              string `json:"name"`
	Alias             string `json:"alias"`
	BaseCurrency      string `json:"baseCurrency"`
	QuoteCurrency     string `json:"quoteCurrency"`
	BasePrecision     string `json:"basePrecision"`
	QuotePrecision    string `json:"quotePrecision"`
	MinTradeQuantity  string `json:"minTradeQuantity"`
	MinTradeAmount    string `json:"minTradeAmount"`
	MaxTradeQuantity  string `json:"maxTradeQuantity"`
	MaxTradeAmount    string `json:"maxTradeAmount"`
	MinPricePrecision string `json:"minPricePrecision"`
	Category          int    `json:"category"`
	ShowStatus        bool   `json:"showStatus"`
	Innovation        bool   `json:"innovation"`
}

type GetSymbolsResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []GetSymbolsResult `json:"result"`
}

type OrderBookDepthParams struct {
	Symbol string `url:"symbol" json:"symbol"`
	Limit  int    `url:"limit,omitempty" json:"limit,omitempty"`
}

type OrderBookDepthResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              OrderBookResult `json:"result"`
}

type OrderBookResult struct {
	Time int64      `json:"time"`
	Bids [][]string `json:"bids"`
	Asks [][]string `json:"asks"`
}

type MergedOrderBookParams struct {
	Symbol string `url:"symbol" json:"symbol"`
	// Precision of the merged orderbook, 1 means 1 digit
	Scale int `url:"scale,omitempty" json:"scale,omitempty"`
	Limit int `url:"limit,omitempty" json:"limit,omitempty"`
}

type MergedOrderBookResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              OrderBookResult `json:"result"`
}

type PublicTradeRecordsParams struct {
	Symbol string `url:"symbol" json:"symbol"`
	Limit  int    `url:"limit,omitempty" json:"limit,omitempty"`
}

type PublicTradeRecordsResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []TradeRecordResult `json:"result"`
}

type TradeRecordResult struct {
	Price        string `json:"price"`
	Time         int64  `json:"time"`
	Qty          string `json:"qty"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
}

type GetKlineParams struct {
	Symbol string `url:"symbol" json:"symbol"`
	// Interval of the kline data
	Interval  string `url:"interval" json:"interval"`
	Limit     int    `url:"limit,omitempty" json:"limit,omitempty"`
	StartTime int64  `url:"startTime,omitempty" json:"startTime,omitempty"`
	EndTime   int64  `url:"endTime,omitempty" json:"endTime,omitempty"`
}

type KlineResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []interface{} `json:"result"`
}

type GetTickerInfo24hrParams struct {
	Symbol string `url:"symbol,omitempty" json:"symbol,omitempty"`
}

type GetTickerInfoResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              TickerDataResult `json:"result"`
}

type TickerDataResult struct {
	Time         int64  `json:"time"`
	Symbol       string `json:"symbol"`
	BestBidPrice string `json:"bestBidPrice"`
	BestAskPrice string `json:"bestAskPrice"`
	Volume       string `json:"volume"`
	QuoteVolume  string `json:"quoteVolume"`
	LastPrice    string `json:"lastPrice"`
	HighPrice    string `json:"highPrice"`
	LowPrice     string `json:"lowPrice"`
	OpenPrice    string `json:"openPrice"`
}

type TickerPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type GetLastTradedPriceParams struct {
	Symbol string `url:"symbol,omitempty" json:"symbol,omitempty"`
}

type GetLastTradedPriceResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              TickerPrice `json:"result"`
}

type GetBidAskPriceParams struct {
	Symbol string `url:"symbol,omitempty" json:"symbol,omitempty"`
}

type BidAskPriceResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              BidAskPriceResult `json:"result"`
}

type BidAskPriceResult struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
	Time     int64  `json:"time"`
}

type ServerTimeResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              ServerTimeResult `json:"result"`
}

type ServerTimeResult struct {
	ServerTime int64 `json:"serverTime"`
}
