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
	Result              []QueryActiveOrderResult `json:"result"`
}

type QueryActiveOrderResult struct {
	UserID       int64              `url:"user_id" json:"user_id"`
	PositionIDx  int64              `url:"position_idx" json:"position_idx"`
	Symbol       string             `url:"symbol" json:"symbol"`
	Side         string             `url:"side" json:"side"`
	OrderType    string             `url:"order_type" json:"order_type"`
	Price        string             `url:"price" json:"price"`
	Qty          string             `url:"qty" json:"qty"`
	TimeInForce  string             `url:"time_in_force" json:"time_in_force"`
	OrderStatus  string             `url:"order_status" json:"order_status"`
	ExtFields    map[string]float64 `url:"ext_fields" json:"ext_fields"`
	LastExecTime string             `url:"last_exec_time" json:"last_exec_time"`
	LeavesQty    string             `url:"leaves_qty" json:"leaves_qty"`
	LeavesValue  string             `url:"leaves_value" json:"leaves_value"`
	CumExecQty   string             `url:"cum_exec_qty" json:"cum_exec_qty"`
	CumExecValue string             `url:"cum_exec_value" json:"cum_exec_value"`
	CumExecFee   string             `url:"cum_exec_fee" json:"cum_exec_fee"`
	RejectReason string             `url:"reject_reason" json:"reject_reason"`
	CancelType   string             `url:"cancel_type" json:"cancel_type"`
	OrderLinkID  string             `url:"order_link_id" json:"order_link_id"`
	CreatedAt    string             `url:"created_at" json:"created_at"`
	UpdatedAt    string             `url:"updated_at" json:"updated_at"`
	OrderID      string             `url:"order_id" json:"order_id"`
	TakeProfit   string             `url:"take_profit" json:"take_profit"`
	StopLoss     string             `url:"stop_loss" json:"stop_loss"`
	TpTriggerBy  string             `url:"tp_trigger_by" json:"tp_trigger_by"`
	SlTriggerBy  string             `url:"sl_trigger_by" json:"sl_trigger_by"`
}

type PlaceConditionalOrderParams struct {
	Side           string  `url:"side" json:"side"`
	Symbol         string  `url:"symbol" json:"symbol"`
	OrderType      string  `url:"order_type" json:"order_type"`
	Qty            string  `url:"qty" json:"qty"`
	Price          string  `url:"price,omitempty" json:"price,omitempty"`
	BasePrice      string  `url:"base_price" json:"base_price"`
	StopPx         string  `url:"stop_px" json:"stop_px"`
	TimeInForce    string  `url:"time_in_force" json:"time_in_force"`
	TriggerBy      string  `url:"trigger_by" json:"trigger_by"`
	CloseOnTrigger bool    `url:"close_on_trigger" json:"close_on_trigger"`
	OrderLinkID    string  `url:"order_link_id,omitempty" json:"order_link_id,omitempty"`
	TakeProfit     float64 `url:"take_profit,omitempty" json:"take_profit,omitempty"`
	StopLoss       float64 `url:"stop_loss,omitempty" json:"stop_loss,omitempty"`
	TpTriggerBy    string  `url:"tp_trigger_by,omitempty" json:"tp_trigger_by,omitempty"`
	SlTriggerBy    string  `url:"sl_trigger_by,omitempty" json:"sl_trigger_by,omitempty"`
}

type PlaceConditionalOrderResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              PlaceConditionalOrderResult `json:"result"`
}

type PlaceConditionalOrderResult struct {
	UserID       int64   `url:"user_id" json:"user_id"`
	Symbol       string  `url:"symbol" json:"symbol"`
	Side         string  `url:"side" json:"side"`
	OrderType    string  `url:"order_type" json:"order_type"`
	Price        string  `url:"price" json:"price"`
	Qty          string  `url:"qty" json:"qty"`
	TimeInForce  string  `url:"time_in_force" json:"time_in_force"`
	Remark       string  `url:"remark" json:"remark"`
	LeavesQty    float64 `url:"leaves_qty" json:"leaves_qty"`
	LeavesValue  float64 `url:"leaves_value" json:"leaves_value"`
	StopPx       string  `url:"stop_px" json:"stop_px"`
	RejectReason string  `url:"reject_reason" json:"reject_reason"`
	StopOrderID  string  `url:"stop_order_id" json:"stop_order_id"`
	OrderLinkID  string  `url:"order_link_id" json:"order_link_id"`
	TriggerBy    string  `url:"trigger_by" json:"trigger_by"`
	BasePrice    string  `url:"base_price" json:"base_price"`
	CreatedAt    string  `url:"created_at" json:"created_at"`
	UpdatedAt    string  `url:"updated_at" json:"updated_at"`
	TpTriggerBy  string  `url:"tp_trigger_by" json:"tp_trigger_by"`
	SlTriggerBy  string  `url:"sl_trigger_by" json:"sl_trigger_by"`
	TakeProfit   string  `url:"take_profit" json:"take_profit"`
	StopLoss     string  `url:"stop_loss" json:"stop_loss"`
}

type GetConditionalOrderParams struct {
	Symbol          string `url:"symbol" json:"symbol"`
	StopOrderStatus string `url:"stop_order_status,omitempty" json:"stop_order_status,omitempty"`
	Direction       string `url:"direction,omitempty" json:"direction,omitempty"`
	Limit           int64  `url:"limit,omitempty" json:"limit,omitempty"`
	Cursor          string `url:"cursor,omitempty" json:"cursor,omitempty"`
}

type GetConditionalOrderResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              GetConditionalOrderResult `json:"result"`
}

type GetConditionalOrderResult struct {
	Data   []GetConditionalOrderData `url:"data" json:"data"`
	Cursor string                    `url:"cursor" json:"cursor"`
}

type GetConditionalOrderData struct {
	UserID          int64  `url:"user_id" json:"user_id"`
	PositionIDx     int64  `url:"position_idx" json:"position_idx"`
	StopOrderStatus string `url:"stop_order_status" json:"stop_order_status"`
	Symbol          string `url:"symbol" json:"symbol"`
	Side            string `url:"side" json:"side"`
	OrderType       string `url:"order_type" json:"order_type"`
	Price           string `url:"price" json:"price"`
	Qty             string `url:"qty" json:"qty"`
	TimeInForce     string `url:"time_in_force" json:"time_in_force"`
	StopOrderType   string `url:"stop_order_type" json:"stop_order_type"`
	TriggerBy       string `url:"trigger_by" json:"trigger_by"`
	BasePrice       string `url:"base_price" json:"base_price"`
	OrderLinkID     string `url:"order_link_id" json:"order_link_id"`
	CreatedAt       string `url:"created_at" json:"created_at"`
	UpdatedAt       string `url:"updated_at" json:"updated_at"`
	StopPx          string `url:"stop_px" json:"stop_px"`
	StopOrderID     string `url:"stop_order_id" json:"stop_order_id"`
	TakeProfit      string `url:"take_profit" json:"take_profit"`
	StopLoss        string `url:"stop_loss" json:"stop_loss"`
	TpTriggerBy     string `url:"tp_trigger_by" json:"tp_trigger_by"`
	SlTriggerBy     string `url:"sl_trigger_by" json:"sl_trigger_by"`
}

type CancelConditionalOrderParams struct {
	Symbol      string `url:"symbol" json:"symbol"`
	StopOrderID string `url:"stop_order_id,omitempty" json:"stop_order_id,omitempty"`
	OrderLinkID string `url:"order_link_id,omitempty" json:"order_link_id,omitempty"`
}

type CancelConditionalOrderResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              CancelConditionalOrderResult `json:"result"`
}

type CancelConditionalOrderResult struct {
	StopOrderID string `url:"stop_order_id" json:"stop_order_id"`
}

type CancelAllConditionalOrdersParams struct {
	Symbol string `url:"symbol" json:"symbol"`
}

type CancelAllConditionalOrdersResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []CancelAllConditionalOrdersResult `json:"result"`
}

type CancelAllConditionalOrdersResult struct {
	ClOrdID           string  `url:"clOrdID" json:"clOrdID"`
	OrderLinkID       string  `url:"order_link_id" json:"order_link_id"`
	UserID            int64   `url:"user_id" json:"user_id"`
	Symbol            string  `url:"symbol" json:"symbol"`
	Side              string  `url:"side" json:"side"`
	OrderType         string  `url:"order_type" json:"order_type"`
	Price             string  `url:"price" json:"price"`
	Qty               float64 `url:"qty" json:"qty"`
	TimeInForce       string  `url:"time_in_force" json:"time_in_force"`
	CreateType        string  `url:"create_type" json:"create_type"`
	CancelType        string  `url:"cancel_type" json:"cancel_type"`
	OrderStatus       string  `url:"order_status" json:"order_status"`
	LeavesQty         float64 `url:"leaves_qty" json:"leaves_qty"`
	LeavesValue       float64 `url:"leaves_value" json:"leaves_value"`
	CreatedAt         string  `url:"created_at" json:"created_at"`
	UpdatedAt         string  `url:"updated_at" json:"updated_at"`
	CrossStatus       string  `url:"cross_status" json:"cross_status"`
	CrossSeq          int64   `url:"cross_seq" json:"cross_seq"`
	StopOrderType     string  `url:"stop_order_type" json:"stop_order_type"`
	TriggerBy         string  `url:"trigger_by" json:"trigger_by"`
	BasePrice         string  `url:"base_price" json:"base_price"`
	ExpectedDirection string  `url:"expected_direction" json:"expected_direction"`
}

type ReplaceConditionalOrderParams struct {
	StopOrderID    string  `url:"stop_order_id" json:"stop_order_id"`
	OrderLinkID    string  `url:"order_link_id" json:"order_link_id"`
	Symbol         string  `url:"symbol" json:"symbol"`
	PRQty          int64   `url:"p_r_qty" json:"p_r_qty"`
	PRPrice        string  `url:"p_r_price" json:"p_r_price"`
	PRTriggerPrice string  `url:"p_r_trigger_price" json:"p_r_trigger_price"`
	TakeProfit     float64 `url:"take_profit" json:"take_profit"`
	StopLoss       float64 `url:"stop_loss" json:"stop_loss"`
	TPTriggerBy    string  `url:"tp_trigger_by" json:"tp_trigger_by"`
	SLTriggerBy    string  `url:"sl_trigger_by" json:"sl_trigger_by"`
}

type ReplaceConditionalOrderResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              ReplaceConditionalOrderResult `json:"result"`
}

type ReplaceConditionalOrderResult struct {
	StopOrderID string `url:"stop_order_id" json:"stop_order_id"`
}

type QueryConditionalOrderParams struct {
	Symbol      string `url:"symbol" json:"symbol"`
	StopOrderID string `url:"stop_order_id,omitempty" json:"stop_order_id,omitempty"`
	OrderLinkID string `url:"order_link_id,omitempty" json:"order_link_id,omitempty"`
}

type QueryConditionalOrderResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []QueryConditionalOrderResult `json:"result"`
}

type QueryConditionalOrderResult struct {
	UserID       int64            `url:"user_id" json:"user_id"`
	PositionIDx  int64            `url:"position_idx" json:"position_idx"`
	Symbol       string           `url:"symbol" json:"symbol"`
	Side         string           `url:"side" json:"side"`
	OrderType    string           `url:"order_type" json:"order_type"`
	Price        string           `url:"price" json:"price"`
	Qty          float64          `url:"qty" json:"qty"`
	StopPx       string           `url:"stop_px" json:"stop_px"`
	BasePrice    string           `url:"base_price" json:"base_price"`
	TimeInForce  string           `url:"time_in_force" json:"time_in_force"`
	OrderStatus  string           `url:"order_status" json:"order_status"`
	ExtFields    map[string]int64 `url:"ext_fields" json:"ext_fields"`
	LeavesQty    float64          `url:"leaves_qty" json:"leaves_qty"`
	CumExecQty   float64          `url:"cum_exec_qty" json:"cum_exec_qty"`
	CumExecValue float64          `url:"cum_exec_value" json:"cum_exec_value"`
	CumExecFee   float64          `url:"cum_exec_fee" json:"cum_exec_fee"`
	RejectReason string           `url:"reject_reason" json:"reject_reason"`
	OrderLinkID  string           `url:"order_link_id" json:"order_link_id"`
	CreatedAt    string           `url:"created_at" json:"created:"`
	UpdatedAt    string           `url:"updated_at" json:"updated_at"`
	OrderID      string           `url:"order_id" json:"order_id"`
	TriggerBy    string           `url:"trigger_by" json:"trigger_by"`
	TakeProfit   float64          `url:"take_profit" json:"take_profit"`
	StopLoss     float64          `url:"stop_loss" json:"stop_loss"`
	TpTriggerBy  string           `url:"tp_trigger_by" json:"tp_trigger_by"`
	SlTriggerBy  string           `url:"sl_trigger_by" json:"sl_trigger_by"`
}

type GetPositionResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              GetPositionResult `json:"result"`
}

type GetPositionResult struct {
	ID                  int64   `url:"id" json:"id"`
	PositionIDx         int64   `url:"position_idx" json:"position_idx"`
	Mode                int64   `url:"mode" json:"mode"`
	UserID              int64   `url:"user_id" json:"user_id"`
	RiskID              int64   `url:"risk_id" json:"risk_id"`
	Symbol              string  `url:"symbol" json:"symbol"`
	Side                string  `url:"side" json:"side"`
	Size                float64 `url:"size" json:"size"`
	PositionValue       string  `url:"position_value" json:"position_value"`
	EntryPrice          string  `url:"entry_price" json:"entry_price"`
	IsIsolated          bool    `url:"is_isolated" json:"is_isolated"`
	AutoAddMargin       int64   `url:"auto_add_margin" json:"auto_add_margin"`
	Leverage            string  `url:"leverage" json:"leverage"`
	EffectiveLeverage   string  `url:"effective_leverage" json:"effective_leverage"`
	PositionMargin      string  `url:"position_margin" json:"position_margin"`
	LiqPrice            string  `url:"liq_price" json:"liq_price"`
	BustPrice           string  `url:"bust_price" json:"bust_price"`
	OccClosingFee       string  `url:"occ_closing_fee" json:"occ_closing_fee:"`
	TakeProfit          string  `url:"take_profit" json:"take_profit"`
	StopLoss            string  `url:"stop_loss" json:"stop_loss"`
	TrailingStop        string  `url:"trailing_stop" json:"trailing_stop"`
	PositionStatus      string  `url:"position_status" json:"position_status"`
	DeleverageIndicator int64   `url:"deleverage_indicator" json:"deleverage_indicator"`
	OcCalcData          string  `url:"oc_calc_data" json:"oc_calc_data"`
	OrderMargin         string  `url:"order_margin" json:"order_margin"`
	WalletBalance       string  `url:"wallet_balance" json:"wallet_balance"`
	RealisedPnl         string  `url:"realised_pnl" json:"realised_pnl"`
	UnrealisedPnl       string  `url:"unrealised_pnl" json:"unrealised_pnl"`
	CumRealisedPnl      string  `url:"cum_realised_pnl" json:"cum_realised_pnl"`
	CrossSeq            int64   `url:"cross_seq" json:"cross_seq"`
	PositionSeq         int64   `url:"position_seq" json:"position_seq"`
	CreatedAt           string  `url:"created_at" json:"created_at"`
	UpdatedAt           string  `url:"updated_at" json:"updated_at"`
	TpSlMode            string  `url:"tp_sl_mode" json:"tp_sl_mode"`
}

type GetPositionWithSymbolParams struct {
	Symbol string `url:"symbol" json:"symbol"`
}

type GetPositionWithSymbolResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []GetPositionWithSymbolResult `json:"result"`
}

type GetPositionWithSymbolResult struct {
	IsValid bool              `url:"is_valid" json:"is_valid"`
	Data    GetPositionResult `url:"data" json:"data"`
}

type ChangeMarginParams struct {
	Symbol string `url:"symbol" json:"symbol"`
	Margin string `url:"margin" json:"margin"`
}

type ChangeMarginResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              float64 `json:"result"`
}
type SetTradingStopParams struct {
	Symbol            string  `url:"symbol" json:"symbol"`
	TakeProfit        float64 `url:"take_profit,omitempty" json:"take_profit,omitempty"`
	StopLoss          float64 `url:"stop_loss,omitempty" json:"stop_loss,omitempty"`
	TrailingStop      float64 `url:"trailing_stop,omitempty" json:"trailing_stop,omitempty"`
	TpTriggerBy       string  `url:"tp_trigger_by,omitempty" json:"tp_trigger_by,omitempty"`
	SlTriggerBy       string  `url:"sl_trigger_by,omitempty" json:"sl_trigger_by,omitempty"`
	NewTrailingActive float64 `url:"new_trailing_active,omitempty" json:"new_trailing_active,omitempty"`
	SlSize            float64 `url:"sl_size,omitempty" json:"sl_size,omitempty"`
	TpSize            float64 `url:"tp_size,omitempty" json:"tp_size,omitempty"`
}

type SetTradingStopResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              SetTradingStopResult `json:"result"`
}

type SetTradingStopResult struct {
	ID                  string             `url:"id" json:"id"`
	UserID              int64              `url:"user_id" json:"user_id"`
	Symbol              string             `url:"symbol" json:"symbol"`
	Side                string             `url:"side" json:"side"`
	Size                float64            `url:"size" json:"size"`
	PositionValue       float64            `url:"position_value" json:"position_value"`
	EntryPrice          float64            `url:"entry_price" json:"entry_price"`
	RiskID              int64              `url:"risk_id" json:"risk_id"`
	AutoAddMargin       int64              `url:"auto_add_margin" json:"auto_add_margin"`
	Leverage            float64            `url:"leverage" json:"leverage"`
	PositionMargin      float64            `url:"position_margin" json:"position_margin"`
	LiqPrice            float64            `url:"liq_price" json:"liq_price"`
	BustPrice           float64            `url:"bust_price" json:"bust_price"`
	OccClosingFee       float64            `url:"occ_closing_fee" json:"occ_closing_fee"`
	OccFundingFee       float64            `url:"occ_funding_fee" json:"occ_funding_fee"`
	TakeProfit          float64            `url:"take_profit" json:"take_profit"`
	StopLoss            float64            `url:"stop_loss" json:"stop_loss"`
	TrailingStop        float64            `url:"trailing_stop" json:"trailing_stop"`
	PositionStatus      string             `url:"position_status" json:"position_status"`
	DeleverageIndicator int64              `url:"deleverage_indicator" json:"deleverage_indicator"`
	OcCalcData          map[string]float64 `url:"oc_calc_data" json:"oc_calc_data"`
	OrderMargin         float64            `url:"order_margin" json:"order_margin"`
	WalletBalance       float64            `url:"wallet_balance" json:"wallet_balance"`
	RealisedPnl         float64            `url:"realised_pnl" json:"realised_pnl"`
	CumRealisedPnl      float64            `url:"cum_realised_pnl" json:"cum_realised_pnl"`
	CumCommission       float64            `url:"cum_commission" json:"cum_commission"`
	CrossSeq            int64              `url:"cross_seq" json:"cross_seq"`
	PositionSeq         int64              `url:"position_seq" json:"position_seq"`
	CreatedAt           string             `url:"created_at" json:"created_at"`
	UpdatedAt           string             `url:"updated_at" json:"updated_at"`
	ExtFields           ExtFields          `url:"ext_fields" json:"ext_fields"`
}

type ExtFields struct {
	TrailingActive string `url:"trailing_active" json:"trailing_active"`
	SlTriggerBy    string `url:"sl_trigger_by" json:"sl_trigger_by"`
	TpTriggerBy    string `url:"tp_trigger_by" json:"tp_trigger_by"`
	V              int64  `url:"v" json:"v"`
	Mm             int64  `url:"mm" json:"mm"`
}

type SetLeverageParams struct {
	Symbol       string  `url:"symbol" json:"symbol"`
	Leverage     float64 `url:"leverage" json:"leverage"`
	LeverageOnly bool    `url:"leverage_only,omitempty" json:"leverage_only"`
}

type SetLeverageResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              float64 `json:"result"`
}

type GetUserTradeRecordsParams struct {
	OrderID   string `url:"order_id,omitempty" json:"order_id,omitempty"`
	Symbol    string `url:"symbol" json:"symbol"`
	StartTime string `url:"start_time,omitempty" json:"start_time,omitempty"`
	Page      int64  `url:"page,omitempty" json:"page,omitempty"`
	Limit     int64  `url:"limit,omitempty" json:"limit,omitempty"`
	Order     string `url:"order,omitempty" json:"order,omitempty"`
}

type GetUserTradeRecordsResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              UserTradeRecordsResult `json:"result"`
}

type UserTradeRecordsResult struct {
	OrderID   string      `url:"order_id" json:"order_id"`
	TradeList []TradeList `url:"trade_list" json:"trade_list"`
}

type TradeList struct {
	ClosedSize       float64 `url:"closed_size" json:"closed_size"`
	CrossSeq         int64   `url:"cross_seq" json:"cross_seq"`
	ExecFee          string  `url:"exec_fee" json:"exec_fee"`
	ExecID           string  `url:"exec_id" json:"exec_id"`
	ExecPrice        string  `url:"exec_price" json:"exec_price"`
	ExecQty          float64 `url:"exec_qty" json:"exec_qty"`
	ExecTime         int64   `url:"exec_time" json:"exec_time"`
	ExecType         string  `url:"exec_type" json:"exec_type"`
	ExecValue        string  `url:"exec_value" json:"exec_value"`
	FeeRate          string  `url:"fee_rate" json:"fee_rate"`
	LastLiquidityInd string  `url:"last_liquidity_ind" json:"last_liquidity_ind"`
	LeavesQty        float64 `url:"leaves_qty" json:"leaves_qty"`
	NthFill          int64   `url:"nth_fill" json:"nth_fill"`
	OrderID          string  `url:"order_id" json:"order_id"`
	OrderLinkID      string  `url:"order_link_id" json:"order_link_id"`
	OrderPrice       string  `url:"order_price" json:"order_price"`
	OrderQty         float64 `url:"order_qty" json:"order_qty"`
	OrderType        string  `url:"order_type" json:"order_type"`
	Side             string  `url:"side" json:"side"`
	Symbol           string  `url:"symbol" json:"symbol"`
	TradeTimeMs      int64   `url:"trade_time_ms" json:"trade_time_ms"`
	UserID           int64   `url:"user_id" json:"user_id"`
}

type ClosedProfitAndLossParams struct {
	Symbol    string `url:"symbol" json:"symbol"`
	StartTime int64  `url:"start_time,omitempty" json:"start_time,omitempty"`
	EndTime   int64  `url:"end_time,omitempty" json:"end_time,omitempty"`
	ExecType  string `url:"exec_type,omitempty" json:"exec_type,omitempty"`
	Page      int64  `url:"page,omitempty" json:"page,omitempty"`
	Limit     int64  `url:"limit,omitempty" json:"limit,omitempty"`
}

type ClosedProfitAndLossResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              ClosedProfitAndLossResult `json:"result"`
}

type ClosedProfitAndLossResult struct {
	CurrentPage int64                     `url:"current_page" json:"current_page"`
	Data        []ClosedProfitAndLossData `url:"data" json:"data"`
}

type ClosedProfitAndLossData struct {
	ID            int64   `url:"id" json:"id"`
	UserID        int64   `url:"user_id" json:"user_id"`
	Symbol        string  `url:"symbol" json:"symbol"`
	OrderID       string  `url:"order_id" json:"order_id"`
	Side          string  `url:"side" json:"side"`
	Qty           float64 `url:"qty" json:"qty"`
	OrderPrice    float64 `url:"order_price" json:"order_price"`
	OrderType     string  `url:"order_type" json:"order_type"`
	ExecType      string  `url:"exec_type" json:"exec_type"`
	ClosedSize    float64 `url:"closed_size" json:"closed_size"`
	CumEntryValue float64 `url:"cum_entry_value" json:"cum_entry_value"`
	AvgEntryPrice float64 `url:"avg_entry_price" json:"avg_entry_price"`
	CumExitValue  float64 `url:"cum_exit_value" json:"cum_exit_value"`
	AvgExitPrice  float64 `url:"avg_exit_price" json:"avg_exit_price"`
	ClosedPnL     float64 `url:"closed_pnl" json:"closed_pnl"`
	FillCount     int64   `url:"fill_count" json:"fill_count"`
	Leverage      float64 `url:"leverage" json:"leverage"`
	CreatedAt     int64   `url:"created_at" json:"created_at"`
}

type PositionTpSlSwitchParams struct {
	Symbol   string `url:"symbol" json:"symbol"`
	TpSlMode string `url:"tp_sl_mode" json:"tp_sl_mode"`
}

type PositionTpSlSwitchResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              PositionTpSlSwitchResult `json:"result"`
}

type PositionTpSlSwitchResult struct {
	TpSlMode string `url:"tp_sl_mode" json:"tp_sl_mode"`
}

type MarginSwitchParams struct {
	Symbol       string  `url:"symbol" json:"symbol"`
	IsIsolated   bool    `url:"is_isolated" json:"is_isolated"`
	BuyLeverage  float64 `url:"buy_leverage" json:"buy_leverage"`
	SellLeverage float64 `url:"sell_leverage" json:"sell_leverage"`
}

type MarginSwitchResponse struct {
	domain.BaseResponse `json:",inline"`
}

type QueryTradingFeeRateParams struct {
	Symbol string `url:"symbol" json:"symbol"`
}

type QueryTradingFeeRateResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              QueryTradingFeeRateResult `json:"result"`
}

type QueryTradingFeeRateResult struct {
	UserID       int64   `url:"user_id" json:"user_id"`
	TakerFeeRate float64 `url:"taker_fee_rate" json:"taker_fee_rate"`
	MakerFeeRate float64 `url:"maker_fee_rate" json:"maker_fee_rate"`
}

type GetRiskLimitParams struct {
	Symbol string `url:"symbol" json:"symbol"`
}

type GetRiskLimitResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []GetRiskLimitResult `json:"result"`
}

type GetRiskLimitResult struct {
	ID             int64   `url:"id" json:"id"`
	Symbol         string  `url:"symbol" json:"symbol"`
	Limit          int64   `url:"limit" json:"limit"`
	MaintainMargin float64 `url:"maintain_margin" json:"maintain_margin"`
	StartingMargin float64 `url:"starting_margin" json:"starting_margin"`
	Section        string  `url:"section" json:"section"`
	IsLowestRisk   int64   `url:"is_lowest_risk" json:"is_lowest_risk"`
	CreatedAt      string  `url:"created_at" json:"created_at"`
	UpdatedAt      string  `url:"updated_at" json:"updated_at"`
	MaxLeverage    float64 `url:"max_leverage" json:"max_leverage"`
}

type SetRiskLimitParams struct {
	Symbol string `url:"symbol" json:"symbol"`
	RiskID int64  `url:"risk_id" json:"risk_id"`
}

type SetRiskLimitResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              SetRiskLimitResult `json:"result"`
}

type SetRiskLimitResult struct {
	RiskID int64 `url:"risk_id" json:"risk_id"`
}

type GetFundingRateParams struct {
	Symbol string `url:"symbol" json:"symbol"`
}

type GetFundingRateResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              GetFundingRateResult `json:"result"`
}

type GetFundingRateResult struct {
	Symbol               string `url:"symbol" json:"symbol"`
	FundingRate          string `url:"funding_rate" json:"funding_rate"`
	FundingRateTimestamp int64  `url:"funding_rate_timestamp" json:"funding_rate_timestamp"`
}

type GetLastFundingFeeParams struct {
	Symbol string `url:"symbol" json:"symbol"`
}

type GetLastFundingFeeResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              GetLastFundingFeeResult `json:"result"`
}

type GetLastFundingFeeResult struct {
	Symbol        string  `url:"symbol" json:"symbol"`
	Side          string  `url:"side" json:"side"`
	Size          float64 `url:"size" json:"size"`
	FundingRate   float64 `url:"funding_rate" json:"funding_rate"`
	ExecFee       float64 `url:"exec_fee" json:"exec_fee"`
	ExecTimestamp int64   `url:"exec_timestamp" json:"exec_timestamp"`
}

type GetPredictedFundingParams struct {
	Symbol string `url:"symbol" json:"symbol"`
}

type GetPredictedFundingResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              GetPredictedFundingResult `json:"result"`
}

type GetPredictedFundingResult struct {
	PredictedFundingRate float64 `url:"predicted_funding_rate" json:"predicted_funding_rate"`
	PredictedFundingFee  float64 `url:"predicted_funding_fee" json:"predicted_funding_fee"`
}

type GetAPIKeyInfoResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []GetAPIKeyInfoResult `json:"result"`
}

type GetAPIKeyInfoResult struct {
	APIKey        string   `url:"api_key" json:"api_key"`
	Type          string   `url:"type" json:"type"`
	UserID        int64    `url:"user_id" json:"user_id"`
	InviterID     int64    `url:"inviter_id" json:"inviter_id"`
	IPs           []string `url:"ips" json:"ips"`
	Note          string   `url:"note" json:"note"`
	Permissions   []string `url:"permissions" json:"permissions"`
	CreatedAt     string   `url:"created_at" json:"created_at"`
	ExpiredAt     string   `url:"expired_at" json:"expired_at"`
	ReadOnly      bool     `url:"read_only" json:"read_only"`
	VIPLevel      string   `url:"vip_level" json:"vip_level"`
	MktMakerLevel string   `url:"mkt_maker_level" json:"mkt_maker_level"`
}

type GetLCPInfoParams struct {
	Symbol string `url:"symbol" json:"symbol"`
}

type GetLCPInfoResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              GetLCPInfoResult `json:"result"`
}

type GetLCPInfoResult struct {
	LCPList []GetLCPInfoList `url:"lcp_list" json:"lcp_list"`
}

type GetLCPInfoList struct {
	Date          string  `url:"date" json:"date"`
	SelfRatio     float64 `url:"self_ratio" json:"self_ratio"`
	PlatformRatio float64 `url:"platform_ratio" json:"platform_ratio"`
	Score         float64 `url:"score" json:"score"`
}
