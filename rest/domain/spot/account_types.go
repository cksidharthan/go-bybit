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
	OrderID             string                `json:"orderId"`
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

type CancelActiveOrderParams struct {
	OrderID     string `url:"orderId,omitempty" json:"orderId,omitempty"`
	OrderLinkID string `url:"orderLinkId,omitempty" json:"orderLinkId,omitempty"`
}

type CancelActiveOrderResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              CancelAllActiveOrdersResult `json:"result"`
}

type CancelAllActiveOrdersResult struct {
	OrderID      string                `json:"orderId"`
	OrderLinkID  string                `json:"orderLinkId"`
	Symbol       string                `json:"symbol"`
	Status       bybit.OrderStatusSpot `json:"status"`
	AccountID    string                `json:"accountId"`
	TransactTime string                `json:"transactTime"`
	Price        string                `json:"price"`
	OrigQty      string                `json:"origQty"`
	ExecutedQty  string                `json:"executedQty"`
	TimeInForce  bybit.TimeInForce     `json:"timeInForce"`
	Type         bybit.OrderTypeSpot   `json:"type"`
	Side         bybit.Side            `json:"side"`
}

type FastCancelActiveOrderParams struct {
	SymbolID    string `url:"symbolId" json:"symbolId"`
	OrderID     string `url:"orderId,omitempty" json:"orderId,omitempty"`
	OrderLinkID string `url:"orderLinkId,omitempty" json:"orderLinkId,omitempty"`
}

type FastCancelActiveOrderResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              FastCancelActiveOrderResult `json:"result"`
}

type FastCancelActiveOrderResult struct {
	IsCancelled bool `json:"isCancelled"`
}

type BatchCancelActiveOrderParams struct {
	Symbol     bybit.Symbol `url:"symbol" json:"symbol"`
	Side       bybit.Side   `url:"side,omitempty" json:"side,omitempty"`
	OrderTypes string       `url:"orderTypes,omitempty" json:"orderTypes,omitempty"`
}

type BatchCancelActiveOrderResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              BatchCancelActiveOrderResult `json:"result"`
}

type BatchCancelActiveOrderResult struct {
	Success bool `json:"success"`
}

type BatchFastCancelActiveOrderParams struct {
	Symbol     string     `url:"symbol" json:"symbol"`
	Side       bybit.Side `url:"side,omitempty" json:"side,omitempty"`
	OrderTypes string     `url:"orderTypes,omitempty" json:"orderTypes,omitempty"`
}

type BatchFastCancelActiveOrderResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              BatchFastCancelActiveOrderResult `json:"result"`
}

type BatchFastCancelActiveOrderResult struct {
	Success bool `json:"success"`
}

type BatchCancelActiveOrdersByIDsParams struct {
	// Order ID, use commas to indicate multiple orderIds. Maximum of 100 ids.
	OrderIDs string `url:"orderIds" json:"orderIds"`
}

type BatchCancelActiveOrdersByIDsResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []BatchCancelActiveOrdersByIDsResult `json:"result"`
}

type BatchCancelActiveOrdersByIDsResult struct {
	OrderID int `json:"orderId"`
	Code    int `json:"code"`
}

type GetOpenOrdersParams struct {
	Symbol  bybit.Symbol `url:"symbol,omitempty" json:"symbol,omitempty"`
	OrderID string       `url:"orderId,omitempty" json:"orderId,omitempty"`
	Limit   int          `url:"limit,omitempty" json:"limit,omitempty"`
}

type GetOpenOrdersResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []GetOpenOrdersResult `json:"result"`
}

type GetOpenOrdersResult struct {
	AccountID           string                `json:"accountId"`
	ExchangeID          string                `json:"exchangeId"`
	Symbol              string                `json:"symbol"`
	SymbolName          string                `json:"symbolName"`
	OrderID             string                `json:"orderId"`
	OrderLinkID         string                `json:"orderLinkId"`
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
}

type GetOrderHistoryParams struct {
	Symbol    bybit.Symbol `url:"symbol,omitempty" json:"symbol,omitempty"`
	OrderID   string       `url:"orderId,omitempty" json:"orderId,omitempty"`
	Limit     int          `url:"limit,omitempty" json:"limit,omitempty"`
	StartTime int          `url:"startTime,omitempty" json:"startTime,omitempty"`
	EndTime   int          `url:"endTime,omitempty" json:"endTime,omitempty"`
}

type GetOrderHistoryResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []GetOrderHistoryResult `json:"result"`
}

type GetOrderHistoryResult struct {
	AccountID           string                `json:"accountId"`
	ExchangeID          string                `json:"exchangeId"`
	Symbol              string                `json:"symbol"`
	SymbolName          string                `json:"symbolName"`
	OrderID             string                `json:"orderId"`
	OrderLinkID         string                `json:"orderLinkId"`
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
}

type GetTradeHistoryParams struct {
	Symbol       bybit.Symbol `url:"symbol,omitempty" json:"symbol,omitempty"`
	Limit        int          `url:"limit,omitempty" json:"limit,omitempty"`
	FromTicketID int          `url:"fromTicketId,omitempty" json:"fromTicketId,omitempty"`
	ToTicketID   int          `url:"toTicketId,omitempty" json:"toTicketId,omitempty"`
	OrderID      int          `url:"orderId,omitempty" json:"orderId,omitempty"`
	StartTime    int          `url:"startTime,omitempty" json:"startTime,omitempty"`
	EndTime      int          `url:"endTime,omitempty" json:"endTime,omitempty"`
}

type GetTradeHistoryResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []GetTradeHistoryResult `json:"result"`
}

type GetTradeHistoryResult struct {
	Symbol          string `json:"symbol"`
	ID              string `json:"id"`
	OrderID         string `json:"orderId"`
	TicketID        string `json:"ticketId"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            string `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	SymbolName      string `json:"symbolName"`
	MatchOrderID    string `json:"matchOrderId"`
	Fee             Fee    `json:"fee"`
	FeeTokenID      string `json:"feeTokenId"`
	FeeAMount       string `json:"feeAmount"`
	MakerRebate     string `json:"makerRebate"`
	ExecutionTime   string `json:"executionTime"`
}

type Fee struct {
	FeeTokenID   string `json:"feeTokenId"`
	FeeTokenName string `json:"feeTokenName"`
	Fee          string `json:"fee"`
}
