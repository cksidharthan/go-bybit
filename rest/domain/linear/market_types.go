package linear

import (
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain"
)

type OrderBookParams struct {
	Symbol   string         `url:"symbol" json:"symbol"`
	Interval bybit.Interval `url:"interval" json:"interval"`
	From     int            `url:"from" json:"from"`
	Limit    int            `url:"limit,omitempty" json:"limit,omitempty"`
}

type OrderBookResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []OrderBookResult `json:"result"`
}

type OrderBookResult struct {
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
