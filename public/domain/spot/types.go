package spot

type BaseResponse struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code""`
	ExtInfo string `json:"ext_info""`
}

type Symbol struct {
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

type SymbolsResponse struct {
	BaseResponse `json:",inline"`
	Result       []Symbol `json:"result"`
}

type OrderBook struct {
	Time int64      `json:"time"`
	Bids [][]string `json:"bids"`
	Asks [][]string `json:"asks"`
}

type OrderBookResponse struct {
	BaseResponse `json:",inline"`
	Result       OrderBook `json:"result"`
}

type TradeRecord struct {
	Price        string `json:"price"`
	Time         int64  `json:"time"`
	Qty          string `json:"qty"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
}

type TradeRecordsResponse struct {
	BaseResponse `json:",inline"`
	Result       []TradeRecord `json:"result"`
}

type KlineResponse struct {
	BaseResponse `json:",inline"`
	Result       []interface{} `json:"result"`
}

type TickerData struct {
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

type TickerResponse struct {
	BaseResponse `json:",inline"`
	Result       TickerData `json:"result"`
}

type TickerPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type PriceResponse struct {
	BaseResponse `json:",inline"`
	Result       TickerPrice `json:"result"`
}

type PriceBidAsk struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
	Time     int64  `json:"time"`
}

type PriceBidAskResponse struct {
	BaseResponse `json:",inline"`
	Result       PriceBidAsk `json:"result"`
}
