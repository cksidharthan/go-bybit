package spot

import (
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain"
)

type PlaceActiveOrderParams struct {
	Symbol bybit.Symbol        `url:"symbol" json:"symbol"`
	Qty    float64             `url:"qty" json:"qty"`
	Side   bybit.Side          `url:"side" json:"side"`
	Type   bybit.OrderTypeSpot `url:"type" json:"type"`

	TimeInForce bybit.TimeInForce `url:"timeInForce,omitempty" json:"timeInForce,omitempty"`
	Price       float64           `url:"price,omitempty" json:"price,omitempty"`
	OrderLinkID string            `url:"orderLinkId,omitempty" json:"orderLinkID,omitempty"`
}

type PlaceActiveOrderResult struct {
	OrderID      string                `json:"orderId"`
	OrderLinkID  string                `json:"orderLinkId"`
	Symbol       string                `json:"symbol"`
	TransactTime string                `json:"transactTime"`
	Price        string                `json:"price"`
	OrigQty      string                `json:"origQty"`
	Type         bybit.OrderTypeSpot   `json:"type"`
	Side         string                `json:"side"`
	Status       bybit.OrderStatusSpot `json:"status"`
	TimeInForce  bybit.TimeInForce     `json:"timeInForce"`
	AccountID    string                `json:"accountId"`
	SymbolName   string                `json:"symbolName"`
	ExecutedQty  string                `json:"executedQty"`
}

type PlaceActiveOrderResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              PlaceActiveOrderResult `json:"result"`
}

type GetActiveOrderParams struct {
	OrderID     string `url:"orderId,omitempty" json:"orderId,omitempty"`
	OrderLinkID string `url:"orderLinkId,omitempty" json:"orderLinkId,omitempty"`
}

type GetActiveOrderResult struct {
	AccountID           string                `json:"accountId"`
	ExchangeID          string                `json:"exchangeId"`
	Symbol              string                `json:"symbol"`
	SymbolName          string                `json:"symbolName"`
	OrderLinkID         string                `json:"orderLinkId"`
	OrderId             string                `json:"orderId"`
	Price               string                `json:"price"`
	OrigQty             string                `json:"origQty"`
	ExecutedQty         string                `json:"executedQty"`
	CummulativeQuoteQty string                `json:"cummulativeQuoteQty"`
	AvgPrice            string                `json:"avgPrice"`
	Status              bybit.OrderStatusSpot `json:"status"`
	TimeInForce         bybit.TimeInForce     `json:"timeInForce"`
	Type                bybit.OrderTypeSpot   `json:"type"`
	Side                bybit.Side            `json:"side"`
	StopPrice           string                `json:"stopPrice"`
	IcebergQty          string                `json:"icebergQty"`
	Time                string                `json:"time"`
	UpdateTime          string                `json:"updateTime"`
	IsWorking           bool                  `json:"isWorking"`
	Locked              string                `json:"locked"`
}

type GetActiveOrderResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              GetActiveOrderResult `json:"result"`
}
