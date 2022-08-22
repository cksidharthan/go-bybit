package inverseperp

import (
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
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
	Size   int64  `json:"size"`
	Side   string `json:"side"`
}

type QueryKlineParams struct {
	Symbol   string `url:"symbol" json:"symbol"`
	Interval string `url:"interval" json:"interval"`
	From     int64  `url:"from" json:"from"`
	Limit    int64  `url:"limit,omitempty" json:"limit,omitempty"`
}

type QueryKlineResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []QueryKlineResult `json:"result"`
}

type QueryKlineResult struct {
	Symbol   string `json:"symbol"`
	Interval string `json:"interval"`
	OpenTime int64  `json:"open_time"`
	Open     string `json:"open"`
	High     string `json:"high"`
	Low      string `json:"low"`
	Close    string `json:"close"`
	Volume   string `json:"volume"`
	Turnover string `json:"turnover"`
}

type GetSymbolInformationParams struct {
	Symbol string `url:"symbol" json:"symbol"`
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
	PrevPrice1h            string  `json:"prev_price_1h"`
	Price1hPcnt            string  `json:"price_1h_pcnt"`
	MarkPrice              string  `json:"mark_price"`
	IndexPrice             string  `json:"index_price"`
	OpenInterest           float64 `json:"open_interest"`
	OpenValue              string  `json:"open_value"`
	TotalTurnover          string  `json:"total_turnover"`
	Turnover24h            string  `json:"turnover_24h"`
	TotalVolume            float64 `json:"total_volume"`
	Volume24h              float64 `json:"volume_24h"`
	FundingRate            string  `json:"funding_rate"`
	PredictedFundingRate   string  `json:"predicted_funding_rate"`
	NextFundingTime        string  `json:"next_funding_time"`
	CountdownHour          float64 `json:"countdown_hour"`
	DeliveryFeeRate        string  `json:"delivery_fee_rate"`
	PredictedDeliveryPrice string  `json:"predicted_delivery_price"`
	DeliveryTime           string  `json:"delivery_time"`
}

type PublicTradingRecordsParams struct {
	Symbol string `url:"symbol" json:"symbol"`
	Limit  int64  `url:"limit,omitempty" json:"limit,omitempty"`
}

type PublicTradingRecordsResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []PublicTradingRecordsResult `json:"result"`
}

type PublicTradingRecordsResult struct {
	ID     float64 `json:"id"`
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Qty    float64 `json:"qty"`
	Side   string  `json:"side"`
	Time   string  `json:"time"`
}

type QuerySymbolsResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []QuerySymbolsResult `json:"result"`
	Time                string               `json:"time"`
}

type QuerySymbolsResult struct {
	Name            string         `json:"name"`
	Alias           string         `json:"alias"`
	Status          string         `json:"status"`
	BaseCurrency    string         `json:"base_currency"`
	QuoteCurrency   string         `json:"quote_currency"`
	PriceScale      int64          `json:"price_scale"`
	TakerFee        string         `json:"taker_fee"`
	MakerFee        string         `json:"maker_fee"`
	FundingInterval float64        `json:"funding_interval"`
	LeverageFilter  LeverageFilter `json:"leverage_filter"`
	PriceFilter     PriceFilter    `json:"price_filter"`
	LotSizeFilter   LotSizeFilter  `json:"lot_size_filter"`
}

type LeverageFilter struct {
	MinLeverage float64 `json:"min_leverage"`
	MaxLeverage float64 `json:"max_leverage"`
	TickSize    float64 `json:"tick_size"`
}

type PriceFilter struct {
	MinPrice float64 `json:"min_price"`
	MaxPrice float64 `json:"max_price"`
	TickSize float64 `json:"tick_size"`
}

type LotSizeFilter struct {
	MaxTradingQty         float64 `json:"max_trading_qty"`
	MinTradingQty         float64 `json:"min_trading_qty"`
	QtyStep               float64 `json:"qty_step"`
	PostOnlyMaxTradingQty float64 `json:"post_only_max_trading_qty"`
}

type QueryMarkPriceKlineParams struct {
	Symbol   string `url:"symbol" json:"symbol"`
	Interval string `url:"interval" json:"interval"`
	From     int64  `url:"from" json:"from"`
	Limit    int64  `url:"limit,omitempty" json:"limit,omitempty"`
}

type QueryMarkPriceKlineResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []QueryMarkPriceKlineResult `json:"result"`
	TimeNow             string                      `json:"time_now"`
}

type QueryMarkPriceKlineResult struct {
	ID      float64 `json:"id"`
	Symbol  string  `json:"symbol"`
	Period  string  `json:"period"`
	StartAt int64   `json:"start_at"`
	Open    int64   `json:"open"`
	High    int64   `json:"high"`
	Low     int64   `json:"low"`
	Close   int64   `json:"close"`
}

type QueryIndexPriceKlineParams struct {
	Symbol   string `url:"symbol" json:"symbol"`
	Interval string `url:"interval" json:"interval"`
	From     int64  `url:"from" json:"from"`
	Limit    int64  `url:"limit,omitempty" json:"limit,omitempty"`
}

type QueryIndexPriceKlineResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []QueryIndexPriceKlineResult `json:"result"`
	TimeNow             string                       `json:"time_now"`
}

type QueryIndexPriceKlineResult struct {
	Symbol   string `json:"symbol"`
	Period   string `json:"period"`
	OpenTime int64  `json:"open_time"`
	Open     int64  `json:"open"`
	High     int64  `json:"high"`
	Low      int64  `json:"low"`
	Close    int64  `json:"close"`
}

type QueryPremiumIndexKlineParams struct {
	Symbol   string `url:"symbol" json:"symbol"`
	Interval string `url:"interval" json:"interval"`
	From     int64  `url:"from" json:"from"`
	Limit    int64  `url:"limit,omitempty" json:"limit,omitempty"`
}

type QueryPremiumIndexKlineResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []QueryPremiumIndexKlineResult `json:"result"`
	TimeNow             string                         `json:"time_now"`
}

type QueryPremiumIndexKlineResult struct {
	Symbol   string `json:"symbol"`
	Period   string `json:"period"`
	OpenTime int64  `json:"open_time"`
	Open     int64  `json:"open"`
	High     int64  `json:"high"`
	Low      int64  `json:"low"`
	Close    int64  `json:"close"`
}

type OpenInterestParams struct {
	Symbol string `url:"symbol" json:"symbol"`
	Period string `url:"period" json:"period"`
	Limit  int64  `url:"limit,omitempty" json:"limit,omitempty"`
}

type OpenInterestResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []OpenInterestResult `json:"result"`
	TimeNow             string               `json:"time_now"`
}

type OpenInterestResult struct {
	OpenInterest float64 `json:"open_interest"`
	Timestamp    int64   `json:"timestamp"`
	Symbol       string  `json:"symbol"`
}

type LatestBigDealParams struct {
	Symbol string `url:"symbol" json:"symbol"`
	Limit  int64  `url:"limit,omitempty" json:"limit,omitempty"`
}

type LatestBigDealResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []LatestBigDealResult `json:"result"`
	TimeNow             string                `json:"time_now"`
}

type LatestBigDealResult struct {
	Symbol    string  `json:"symbol"`
	Side      string  `json:"side"`
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

type LongShortRatioParams struct {
	Symbol string `url:"symbol" json:"symbol"`
	Period string `url:"period" json:"period"`
	Limit  int64  `url:"limit,omitempty" json:"limit,omitempty"`
}

type LongShortRatioResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []LongShortRatioResult `json:"result"`
	TimeNow             string                 `json:"time_now"`
}

type LongShortRatioResult struct {
	Symbol    string  `json:"symbol"`
	BuyRatio  float64 `json:"buy_ratio"`
	SellRatio float64 `json:"sell_ratio"`
	Timestamp int64   `json:"timestamp"`
}
