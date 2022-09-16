package inverseperp

import (
	"github.com/cksidharthan/go-bybit/rest/domain"
)

type PlaceActiveOrderParams struct {
	Side           string  `url:"side" json:"side"`
	Symbol         string  `url:"symbol" json:"symbol"`
	OrderType      string  `url:"order_type" json:"order_type"`
	Qty            float64 `url:"qty" json:"qty"`
	Price          float64 `url:"price,omitempty" json:"price,omitempty"`
	TimeInForce    string  `url:"time_in_force" json:"time_in_force"`
	ReduceOnly     bool    `url:"reduce_only,omitempty" json:"reduce_only,omitempty"`
	CloseOnTrigger bool    `url:"close_on_trigger,omitempty" json:"close_on_trigger,omitempty"`
	OrderLinkID    string  `url:"order_link_id,omitempty" json:"order_link_id,omitempty"`
	TakeProfit     float64 `url:"take_profit,omitempty" json:"take_profit,omitempty"`
	StopLoss       float64 `url:"stop_loss,omitempty" json:"stop_loss,omitempty"`
	TpTriggerBy    string  `url:"tp_trigger_by,omitempty" json:"tp_trigger_by,omitempty"`
	SlTriggerBy    string  `url:"sl_trigger_by,omitempty" json:"sl_trigger_by,omitempty"`
}

type PlaceActiveOrderResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              PlaceActiveOrderResult `json:"result"`
}

type PlaceActiveOrderResult struct {
	UserID        int64  `url:"user_id" json:"user_id"`
	OrderID       string `url:"order_id" json:"order_id"`
	Symbol        string `url:"symbol" json:"symbol"`
	Side          string `url:"side" json:"side"`
	OrderType     string `url:"order_type" json:"order_type"`
	Price         string `url:"price" json:"price"`
	Qty           string `url:"qty" json:"qty"`
	TimeInForce   string `url:"time_in_force" json:"time_in_force"`
	OrderStatus   string `url:"order_status" json:"order_status"`
	LastExecTime  string `url:"last_exec_time" json:"last_exec_time"`
	LastExecPrice string `url:"last_exec_price" json:"last_exec_price"`
	LeavesQty     string `url:"leaves_qty" json:"leaves_qty"`
	CumExecQty    string `url:"cum_exec_qty" json:"cum_exec_qty"`
	CumExecValue  string `url:"cum_exec_value" json:"cum_exec_value"`
	CumExecFee    string `url:"cum_exec_fee" json:"cum_exec_fee"`
	RejectReason  string `url:"reject_reason" json:"reject_reason"`
	OrderLinkID   string `url:"order_link_id" json:"order_link_id"`
	CreatedAt     string `url:"created_at" json:"created_at"`
	UpdatedAt     string `url:"updated_at" json:"updated_at"`
	TakeProfit    string `url:"take_profit" json:"take_profit"`
	StopLoss      string `url:"stop_loss" json:"stop_loss"`
	TpTriggerBy   string `url:"tp_trigger_by" json:"tp_trigger_by"`
	SlTriggerBy   string `url:"sl_trigger_by" json:"sl_trigger_by"`
}

type GetActiveOrderParams struct {
	Symbol      string `url:"symbol" json:"symbol"`
	OrderStatus string `url:"order_status,omitempty" json:"order_status,omitempty"`
	Direction   string `url:"direction,omitempty" json:"direction,omitempty"`
	Limit       int64  `url:"limit,omitempty" json:"limit,omitempty"`
	Cursor      string `url:"cursor,omitempty" json:"cursor,omitempty"`
}

type GetActiveOrderResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              GetActiveOrderResult `json:"result"`
}

type GetActiveOrderResult struct {
	Data   []GetActiveOrderData `url:"data" json:"data"`
	Cursor string               `url:"cursor" json:"cursor"`
}

type GetActiveOrderData struct {
	UserID       int64  `url:"user_id" json:"user_id"`
	PositionIDx  int64  `url:"position_idx" json:"position_idx"`
	OrderStatus  string `url:"order_status" json:"order_status"`
	Symbol       string `url:"symbol" json:"symbol"`
	Side         string `url:"side" json:"side"`
	OrderType    string `url:"order_type" json:"order_type"`
	Price        string `url:"price" json:"price"`
	Qty          string `url:"qty" json:"qty"`
	TimeInForce  string `url:"time_in_force" json:"time_in_force"`
	OrderLinkID  string `url:"order_link_id" json:"order_link_id"`
	OrderID      string `url:"order_id" json:"order_id"`
	CreatedAt    string `url:"created_at" json:"created_at"`
	UpdatedAt    string `url:"updated_at" json:"updated_at"`
	LeavesQty    string `url:"leaves_qty" json:"leaves_qty"`
	LeavesValue  string `url:"leaves_value" json:"leaves_value"`
	CumExecQty   string `url:"cum_exec_qty" json:"cum_exec_qty"`
	CumExecValue string `url:"cum_exec_value" json:"cum_exec_value"`
	CumExecFee   string `url:"cum_exec_fee" json:"cum_exec_fee"`
	RejectReason string `url:"reject_reason" json:"reject_reason"`
	TakeProfit   string `url:"take_profit" json:"take_profit"`
	StopLoss     string `url:"stop_loss" json:"stop_loss"`
	TpTriggerBy  string `url:"tp_trigger_by" json:"tp_trigger_by"`
	SlTriggerBy  string `url:"sl_trigger_by" json:"sl_trigger_by"`
}

type CancelActiveOrderParams struct {
	Symbol      string `url:"symbol" json:"symbol"`
	OrderID     string `url:"order_id,omitempty" json:"order_id,omitempty"`
	OrderLinkID string `url:"order_link_id,omitempty" json:"order_link_id,omitempty"`
}

type CancelActiveOrderResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              CancelActiveOrderResult `json:"result"`
}

type CancelActiveOrderResult struct {
	UserID        int64  `url:"user_id" json:"user_id"`
	OrderID       string `url:"order_id" json:"order_id"`
	Symbol        string `url:"symbol" json:"symbol"`
	Side          string `url:"side" json:"side"`
	OrderType     string `url:"order_type" json:"order_type"`
	Price         string `url:"price" json:"price"`
	Qty           string `url:"qty" json:"qty"`
	TimeInForce   string `url:"time_in_force" json:"time_in_force"`
	OrderStatus   string `url:"order_status" json:"order_status"`
	LastExecTime  string `url:"last_exec_time" json:"last_exec_time"`
	LastExecPrice string `url:"last_exec_price" json:"last_exec_price"`
	LeavesQty     string `url:"leaves_qty" json:"leaves_qty"`
	CumExecQty    string `url:"cum_exec_qty" json:"cum_exec_qty"`
	CumExecValue  string `url:"cum_exec_value" json:"cum_exec_value"`
	CumExecFee    string `url:"cum_exec_fee" json:"cum_exec_fee"`
	RejectReason  string `url:"reject_reason" json:"reject_reason"`
	OrderLinkID   string `url:"order_link_id" json:"order_link_id"`
	CreatedAt     string `url:"created_at" json:"created_at"`
	UpdatedAt     string `url:"updated_at" json:"updated_at"`
	TakeProfit    string `url:"take_profit" json:"take_profit"`
	StopLoss      string `url:"stop_loss" json:"stop_loss"`
	TpTriggerBy   string `url:"tp_trigger_by" json:"tp_trigger_by"`
	SlTriggerBy   string `url:"sl_trigger_by" json:"sl_trigger_by"`
}

type CancelAllActiveOrdersParams struct {
	Symbol string `url:"symbol" json:"symbol"`
}

type CancelAllActiveOrdersResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []CancelAllActiveOrdersResult `json:"result"`
}

type CancelAllActiveOrdersResult struct {
	ClOrdID     string `url:"clOrdID" json:"clOrdID"`
	OrderLinkID string `url:"order_link_id" json:"order_link_id"`
	UserID      int64  `url:"user_id" json:"user_id"`
	Symbol      string `url:"symbol" json:"symbol"`
	Side        string `url:"side" json:"side"`
	OrderType   string `url:"order_type" json:"order_type"`
	Price       string `url:"price" json:"price"`
	Qty         string `url:"qty" json:"qty"`
	TimeInForce string `url:"time_in_force" json:"time_in_force"`
	CreateType  string `url:"create_type" json:"create_type"`
	CancelType  string `url:"cancel_type" json:"cancel_type"`
	OrderStatus string `url:"order_status" json:"order_status"`
	LeavesQty   string `url:"leaves_qty" json:"leaves_qty"`
	LeavesValue string `url:"leaves_value" json:"leaves_value"`
	CreatedAt   string `url:"created_at" json:"created_at"`
	UpdatedAt   string `url:"updated_at" json:"updated_at"`
	CrossStatus string `url:"cross_status" json:"cross_status"`
	CrossSeq    int64  `url:"cross_seq" json:"cross_seq"`
}

type ReplaceActiveOrderParams struct {
	OrderID     string `url:"order_id,omitempty" json:"order_id,omitempty"`
	OrderLinkID string `url:"order_link_id,omitempty" json:"order_link_id,omitempty"`
	Symbol      string `url:"symbol" json:"symbol"`
	PRQty       string `url:"p_r_qty,omitempty" json:"p_r_qty,omitempty"`
	PRPrice     string `url:"p_r_price,omitempty" json:"p_r_price,omitempty"`
	TakeProfit  string `url:"take_profit,omitempty" json:"take_profit,omitempty"`
	StopLoss    string `url:"stop_loss,omitempty" json:"stop_loss,omitempty"`
	TpTriggerBy string `url:"tp_trigger_by,omitempty" json:"tp_trigger_by,omitempty"`
	SlTriggerBy string `url:"sl_trigger_by,omitempty" json:"sl_trigger_by,omitempty"`
}

type ReplaceActiveOrderResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              ReplaceActiveOrderResult `json:"result"`
}

type ReplaceActiveOrderResult struct {
	OrderID string `url:"order_id" json:"order_id"`
}

type QueryActiveOrderParams struct {
	Symbol      string `url:"symbol" json:"symbol"`
	OrderID     string `url:"order_id,omitempty" json:"order_id,omitempty"`
	OrderLinkID string `url:"order_link_id,omitempty" json:"order_link_id,omitempty"`
}

type QueryActiveOrderResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              QueryActiveOrderResult `json:"result"`
}

type QueryActiveOrderResult struct {
	UserID       int64  `url:"user_id" json:"user_id"`
	PositionIDx  int64  `url:"position_idx" json:"position_idx"`
	Symbol       string `url:"symbol" json:"symbol"`
	Side         string `url:"side" json:"side"`
	OrderType    string `url:"order_type" json:"order_type"`
	Price        string `url:"price" json:"price"`
	Qty          string `url:"qty" json:"qty"`
	TimeInForce  string `url:"time_in_force" json:"time_in_force"`
	OrderStatus  string `url:"order_status" json:"order_status"`
	ExtFields    string `url:"ext_fields" json:"ext_fields"`
	LastExecTime string `url:"last_exec_time" json:"last_exec_time"`
	LeavesQty    string `url:"leaves_qty" json:"leaves_qty"`
	LeavesValue  string `url:"leaves_value" json:"leaves_value"`
	CumExecQty   string `url:"cum_exec_qty" json:"cum_exec_qty"`
	CumExecValue string `url:"cum_exec_value" json:"cum_exec_value"`
	CumExecFee   string `url:"cum_exec_fee" json:"cum_exec_fee"`
	RejectReason string `url:"reject_reason" json:"reject_reason"`
	CancelType   string `url:"cancel_type" json:"cancel_type"`
	OrderLinkID  string `url:"order_link_id" json:"order_link_id"`
	CreatedAt    string `url:"created_at" json:"created_at"`
	UpdatedAt    string `url:"updated_at" json:"updated_at"`
	OrderID      string `url:"order_id" json:"order_id"`
	TakeProfit   string `url:"take_profit" json:"take_profit"`
	StopLoss     string `url:"stop_loss" json:"stop_loss"`
	TpTriggerBy  string `url:"tp_trigger_by" json:"tp_trigger_by"`
	SlTriggerBy  string `url:"sl_trigger_by" json:"sl_trigger_by"`
}
