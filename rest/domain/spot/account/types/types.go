package types

import (
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain"
)

type PlaceActiveOrderParams struct {
	Symbol bybit.Symbol    `json:"symbol"`
	Qty    float64         `json:"qty"`
	Side   bybit.Side      `json:"side"`
	Type   bybit.OrderType `json:"type"`

	TimeInForce bybit.TimeInForce `json:"timeInForce,omitempty"`
	Price       float64           `json:"price,omitempty"`
	OrderLinkID string            `json:"orderLinkId,omitempty"`
}

type PlaceActiveOrderResponse struct {
	domain.BaseResponse `json:",inline"`
	OrderID             int                   `json:"orderId"`
	OrderLinkID         string                `json:"orderLinkId"`
	Symbol              string                `json:"symbol"`
	TransactTime        int                   `json:"transactTime"`
	Price               float64               `json:"price"`
	OrigQty             float64               `json:"origQty"`
	Type                bybit.OrderTypeSpot   `json:"type"`
	Side                string                `json:"side"`
	Status              bybit.OrderStatusSpot `json:"status"`
	TimeInForce         bybit.TimeInForce     `json:"timeInForce"`
	AccountID           int64                 `json:"accountId"`
	SymbolName          string                `json:"symbolName"`
	ExecutedQty         string                `json:"executedQty"`
}
