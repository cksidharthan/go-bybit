package linear

import (
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain"
)

type OrderBookParams struct {
	Symbol string `url:"symbol" json:"symbol"`
}

type OrderBookResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []OrderBookResult `json:"result"`
}

type OrderBookResult struct {
	Symbol string  `json:"symbol"`
	Price  string  `json:"price"`
	Size   float64 `json:"size"`
	Side   string  `json:"side"`
}

type QueryKlineParams struct {
	Symbol   string         `url:"symbol" json:"symbol"`
	Interval bybit.Interval `url:"interval" json:"interval"`
	From     int            `url:"from" json:"from"`
	Limit    int            `url:"limit,omitempty" json:"limit,omitempty"`
}

type QueryKlineResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []QueryKlineResult `json:"result"`
}

type QueryKlineResult struct {
	Symbol   string  `json:"symbol"`
	Period   string  `json:"period"`
	Interval string  `json:"interval"`
	StartAt  int     `json:"start_at"`
	OpenTime int     `json:"open_time"`
	Volume   float64 `json:"volume"`
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	Turnover float64 `json:"turnover"`
}

type GetSymbolInformationParams struct {
	Symbol string `url:"symbol,omitempty" json:"symbol,omitempty"`
}

type GetSymbolInformationResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []GetSymbolInformationResult `json:"result"`
}

type GetSymbolInformationResult struct {
	Symbol                 string  `json:"symbol"`
	BidPrice               string  `json:"bid_price"`
	AskPrice               string  `json:"ask_price"`
	LastPrice              string  `json:"last_price"`
	LastTickDirection      string  `json:"last_tick_direction"`
	PrevPrice24h           string  `json:"prev_price_24h"`
	Price24hPcnt           string  `json:"price_24h_pcnt"`
	HighPrice24h           string  `json:"high_price_24h"`
	LowPrice24h            string  `json:"low_price_24h"`
	PricePrive1h           string  `json:"price_prive_1h"`
	Price1HPcnt            string  `json:"price_1h_pcnt"`
	MarkPrice              string  `json:"mark_price"`
	IndexPrice             string  `json:"index_price"`
	OpenInterest           float64 `json:"open_interest"`
	OpenValue              string  `json:"open_value"`
	TotalTurnover          string  `json:"total_turnover"`
	Turnover24h            string  `json:"turnover_24h"`
	TotalVolume            float64 `json:"total_volume"`
	Volume24H              float64 `json:"volume_24h"`
	FundingRate            string  `json:"funding_rate"`
	PredictedFundingRate   string  `json:"predicted_funding_rate"`
	NextFundingTime        string  `json:"next_funding_time"`
	CountdownHour          float64 `json:"countdown_hour"`
	DeliveryFeeRate        string  `json:"delivery_fee_rate"`
	PredictedDeliveryPrice string  `json:"predicted_delivery_price"`
	DeliveryTime           string  `json:"delivery_time"`
}

type GetPublicTradingRecordsParams struct {
	Symbol string `url:"symbol" json:"symbol"`
	// Limit for data size, max size is 1000. Default size is 500
	Limit int `url:"limit,omitempty" json:"limit,omitempty"`
}

type GetPublicTradingRecordsResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []GetPublicTradingRecordsResult `json:"result"`
}

type GetPublicTradingRecordsResult struct {
	ID          string  `json:"id"`
	Symbol      string  `json:"symbol"`
	Price       float64 `json:"price"`
	Qty         float64 `json:"qty"`
	Side        string  `json:"side"`
	Time        string  `json:"time"`
	TradeTimeMs int     `json:"trade_time_ms"`
}

type GetLastFundingRateParams struct {
	Symbol string `url:"symbol" json:"symbol"`
}

type GetLastFundingRateResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              GetLastFundingRateResult `json:"result"`
}

type GetLastFundingRateResult struct {
	Symbol               string  `json:"symbol"`
	FundingRate          float64 `json:"funding_rate"`
	FundingRateTimestamp string  `json:"funding_rate_timestamp"`
}

type QuerySymbolResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []QuerySymbolResult `json:"result"`
}

type QuerySymbolResult struct {
	Name                  string  `json:"name"`
	Alias                 string  `json:"alias"`
	Status                string  `json:"status"`
	BaseCurrency          string  `json:"base_currency"`
	QuoteCurrency         string  `json:"quote_currency"`
	PriceScale            int     `json:"price_scale"`
	TakerFee              string  `json:"taker_fee"`
	MakerFee              string  `json:"maker_fee"`
	FundingInterval       int     `json:"funding_interval"`
	MinLeverage           float64 `json:"min_leverage"`
	MaxLeverage           float64 `json:"max_leverage"`
	LeverageStep          string  `json:"leverage_step"`
	MinPrice              string  `json:"min_price"`
	MaxPrice              string  `json:"max_price"`
	TickSize              string  `json:"tick_size"`
	MaxTradingQty         float64 `json:"max_trading_qty"`
	MinTradingQty         float64 `json:"min_trading_qty"`
	QtyStep               float64 `json:"qty_step"`
	PostOnlyMaxTradingQty string  `json:"post_only_max_trading_qty"`
}

type GetLiquidatedOrdersParams struct {
	Symbol    string `url:"symbol" json:"symbol"`
	From      int    `url:"from,omitempty" json:"from,omitempty"`
	Limit     int    `url:"limit,omitempty" json:"limit,omitempty"`
	StartTime int    `url:"start_time,omitempty" json:"start_time,omitempty"`
	EndTime   int    `url:"end_time,omitempty" json:"end_time,omitempty"`
}

type GetLiquidatedOrdersResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []GetLiquidatedOrdersResult `json:"result"`
}

type GetLiquidatedOrdersResult struct {
	ID     string  `json:"id"`
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Qty    float64 `json:"qty"`
	Side   string  `json:"side"`
	Time   string  `json:"time"`
}

type QueryMarkPriceKlineParams struct {
	Symbol   string         `url:"symbol" json:"symbol"`
	Interval bybit.Interval `url:"interval" json:"interval"`
	From     int            `url:"from,omitempty" json:"from,omitempty"`
	Limit    int            `url:"limit,omitempty" json:"limit,omitempty"`
}

type QueryMarkPriceKlineResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []QueryMarkPriceKlineResult `json:"result"`
}

type QueryMarkPriceKlineResult struct {
	Symbol  string  `json:"symbol"`
	Period  string  `json:"period"`
	Open    float64 `json:"open"`
	High    float64 `json:"high"`
	Low     float64 `json:"low"`
	Close   float64 `json:"close"`
	StartAt int     `json:"start_at"`
}
