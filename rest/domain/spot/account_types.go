package spot

import (
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain"
)

type ActiveOrderParams struct {
	Symbol bybit.Symbol        `url:"symbol" json:"symbol"`
	Qty    float64             `url:"qty" json:"qty"`
	Side   bybit.Side          `url:"side" json:"side"`
	Type   bybit.OrderTypeSpot `url:"type" json:"type"`

	TimeInForce bybit.TimeInForce `url:"timeInForce,omitempty" json:"timeInForce,omitempty"`
	Price       float64           `url:"price,omitempty" json:"price,omitempty"`
	OrderLinkID string            `url:"orderLinkId,omitempty" json:"orderLinkID,omitempty"`
}

type ActiveOrderResult struct {
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

type ActiveOrderResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              ActiveOrderResult `json:"result"`
}
