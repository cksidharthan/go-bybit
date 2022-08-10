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
	domain.BaseResponse `json:",inline"`
	Result              PlaceActiveOrderResult `json:"result"`
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
