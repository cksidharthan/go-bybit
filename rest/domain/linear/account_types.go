package linear

import (
	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain"
)

type PlaceActiveOrderParams struct {
	Side           bybit.Side        `url:"side" json:"side"`
	Symbol         string            `url:"symbol" json:"symbol"`
	OrderType      bybit.OrderType   `url:"order_type" json:"order_type"`
	Qty            float64           `url:"qty" json:"qty"`
	Price          float64           `url:"price,omitempty" json:"price,omitempty"`
	TimeInForce    bybit.TimeInForce `url:"time_in_force" json:"time_in_force"`
	ReduceOnly     bool              `url:"reduce_only" json:"reduce_only"`
	CloseOnTrigger bool              `url:"close_on_trigger" json:"close_on_trigger"`
	OrderLinkID    string            `url:"order_link_id,omitempty" json:"order_link_id,omitempty"`
	TakeProfit     float64           `url:"take_profit,omitempty" json:"take_profit,omitempty"`
	StopLoss       float64           `url:"stop_loss,omitempty" json:"stop_loss,omitempty"`
	TpTriggerBy    string            `url:"tp_trigger_by,omitempty" json:"tp_trigger_by,omitempty"`
	SlTriggerBy    string            `url:"sl_trigger_by,omitempty" json:"sl_trigger_by,omitempty"`
	PositionIDx    int               `url:"position_idx,omitempty" json:"position_idx,omitempty"`
}

type PlaceActiveOrderResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    PlaceActiveOrderResult `json:"result"`
}

type PlaceActiveOrderResult struct {
	UserID         int     `json:"user_id"`
	OrderID        string  `json:"order_id"`
	Symbol         string  `json:"symbol"`
	Side           string  `json:"side"`
	OrderType      string  `json:"order_type"`
	Price          float64 `json:"price"`
	Qty            float64 `json:"qty"`
	TimeInForce    string  `json:"time_in_force"`
	OrderStatus    string  `json:"order_status"`
	LastExecPrice  float64 `json:"last_exec_price"`
	CumExecQty     float64 `json:"cum_exec_qty"`
	CumExecValue   float64 `json:"cum_exec_value"`
	CumExecFee     float64 `json:"cum_exec_fee"`
	ReduceOnly     bool    `json:"reduce_only"`
	CloseOnTrigger bool    `json:"close_on_trigger"`
	OrderLinkID    string  `json:"order_link_id"`
	CreatedTime    string  `json:"created_time"`
	UpdatedTime    string  `json:"updated_time"`
	TakeProfit     float64 `json:"take_profit"`
	StopLoss       float64 `json:"stop_loss"`
	TpTriggerBy    string  `json:"tp_trigger_by"`
	SlTriggerBy    string  `json:"sl_trigger_by"`
	PositionIDx    int     `json:"position_idx"`
}

type GetActiveOrderParams struct {
	Symbol      string `url:"symbol" json:"symbol"`
	OrderID     string `url:"order_id,omitempty" json:"order_id,omitempty"`
	OrderLinkID string `url:"order_link_id,omitempty" json:"order_link_id,omitempty"`
	Order       string `url:"order,omitempty" json:"order,omitempty"`
	Page        int    `url:"page,omitempty" json:"page,omitempty"`
	Limit       int    `url:"limit,omitempty" json:"limit,omitempty"`
	OrderStatus string `url:"order_status,omitempty" json:"order_status,omitempty"`
}

type GetActiveOrderResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    GetActiveOrderResult `json:"result"`
}

type GetActiveOrderResult struct {
	CurrentPage int                        `json:"current_page"`
	Data        []GetActiveOrderResultData `json:"data"`
}

type GetActiveOrderResultData struct {
	OrderID        string  `json:"order_id"`
	UserID         int     `json:"user_id"`
	Symbol         string  `json:"symbol"`
	Side           string  `json:"side"`
	OrderType      string  `json:"order_type"`
	Price          float64 `json:"price"`
	Qty            float64 `json:"qty"`
	TimeInForce    string  `json:"time_in_force"`
	OrderStatus    string  `json:"order_status"`
	LastExecPrice  float64 `json:"last_exec_price"`
	CumExecQty     float64 `json:"cum_exec_qty"`
	CumExecValue   float64 `json:"cum_exec_value"`
	CumExecFee     float64 `json:"cum_exec_fee"`
	OrderLinkID    string  `json:"order_link_id"`
	ReduceOnly     bool    `json:"reduce_only"`
	CloseOnTrigger bool    `json:"close_on_trigger"`
	CreatedTime    string  `json:"created_time"`
	UpdatedTime    string  `json:"updated_time"`
	TakeProfit     float64 `json:"take_profit"`
	StopLoss       float64 `json:"stop_loss"`
	TpTriggerBy    string  `json:"tp_trigger_by"`
	SlTriggerBy    string  `json:"sl_trigger_by"`
}

type CancelActiveOrderParams struct {
	Symbol      string `url:"symbol" json:"symbol"`
	OrderID     string `url:"order_id,omitempty" json:"order_id,omitempty"`
	OrderLinkID string `url:"order_link_id,omitempty" json:"order_link_id,omitempty"`
}

type CancelActiveOrderResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    CancelActiveOrderResult `json:"result"`
}

type CancelActiveOrderResult struct {
	OrderID string `json:"order_id"`
}

type CancelAllActiveOrdersParams struct {
	Symbol string `url:"symbol" json:"symbol"`
}

type CancelAllActiveOrdersResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    []string `json:"result"`
}

type ReplaceActiveOrderParams struct {
	Symbol      string  `url:"symbol" json:"symbol"`
	OrderID     string  `url:"order_id,omitempty" json:"order_id,omitempty"`
	OrderLinkID string  `url:"order_link_id,omitempty" json:"order_link_id,omitempty"`
	PRQty       float64 `url:"p_r_qty,omitempty" json:"p_r_qty,omitempty"`
	PRPrice     float64 `url:"p_r_price,omitempty" json:"p_r_price,omitempty"`
	TakeProfit  float64 `url:"take_profit,omitempty" json:"take_profit,omitempty"`
	StopLoss    float64 `url:"stop_loss,omitempty" json:"stop_loss,omitempty"`
	TpTriggerBy string  `url:"tp_trigger_by,omitempty" json:"tp_trigger_by,omitempty"`
	SlTriggerBy string  `url:"sl_trigger_by,omitempty" json:"sl_trigger_by,omitempty"`
}

type ReplaceActiveOrderResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    ReplaceActiveOrderResult `json:"result"`
}

type ReplaceActiveOrderResult struct {
	OrderID string `json:"order_id"`
}

type QueryActiveOrderParams struct {
	Symbol      string `url:"symbol" json:"symbol"`
	OrderID     string `url:"order_id,omitempty" json:"order_id,omitempty"`
	OrderLinkID string `url:"order_link_id,omitempty" json:"order_link_id,omitempty"`
}

type QueryActiveOrderResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    []QueryActiveOrderResult `json:"result"`
}

type QueryActiveOrderResult struct {
	OrderID        string  `json:"order_id"`
	UserID         int     `json:"user_id"`
	Symbol         string  `json:"symbol"`
	Side           string  `json:"side"`
	OrderType      string  `json:"order_type"`
	Price          float64 `json:"price"`
	Qty            float64 `json:"qty"`
	TimeInForce    string  `json:"time_in_force"`
	OrderStatus    string  `json:"order_status"`
	LastExecPrice  float64 `json:"last_exec_price"`
	CumExecQty     float64 `json:"cum_exec_qty"`
	CumExecValue   float64 `json:"cum_exec_value"`
	CumExecFee     float64 `json:"cum_exec_fee"`
	ReduceOnly     bool    `json:"reduce_only"`
	CloseOnTrigger bool    `json:"close_on_trigger"`
	OrderLinkID    string  `json:"order_link_id"`
	CreatedTime    string  `json:"created_time"`
	UpdatedTime    string  `json:"updated_time"`
	TakeProfit     float64 `json:"take_profit"`
	StopLoss       float64 `json:"stop_loss"`
	TpTriggerBy    string  `json:"tp_trigger_by"`
	SlTriggerBy    string  `json:"sl_trigger_by"`
}
