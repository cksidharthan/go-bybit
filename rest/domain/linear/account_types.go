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

type PlaceConditionalOrderParams struct {
	Side           bybit.Side        `url:"side" json:"side"`
	Symbol         string            `url:"symbol" json:"symbol"`
	OrderType      bybit.OrderType   `url:"order_type" json:"order_type"`
	Qty            float64           `url:"qty" json:"qty"`
	Price          float64           `url:"price,omitempty" json:"price,omitempty"`
	BasePrice      float64           `url:"base_price" json:"base_price"`
	StopPx         float64           `url:"stop_px" json:"stop_px"`
	TimeInForce    bybit.TimeInForce `url:"time_in_force" json:"time_in_force"`
	TriggerBy      string            `url:"trigger_by" json:"trigger_by"`
	ReduceOnly     bool              `url:"reduce_only" json:"reduce_only"`
	CloseOnTrigger bool              `url:"close_on_trigger" json:"close_on_trigger"`
	OrderLinkID    string            `url:"order_link_id,omitempty" json:"order_link_id,omitempty"`
	TakeProfit     float64           `url:"take_profit,omitempty" json:"take_profit,omitempty"`
	StopLoss       float64           `url:"stop_loss,omitempty" json:"stop_loss,omitempty"`
	TpTriggerBy    string            `url:"tp_trigger_by,omitempty" json:"tp_trigger_by,omitempty"`
	SlTriggerBy    string            `url:"sl_trigger_by,omitempty" json:"sl_trigger_by,omitempty"`
	PositionIdx    int               `url:"position_idx,omitempty" json:"position_idx,omitempty"`
}

type PlaceConditionalOrderResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    PlaceConditionalOrderResult `json:"result"`
}

type PlaceConditionalOrderResult struct {
	StopOrderID    string  `json:"stop_order_id"`
	UserID         string  `json:"order_id"`
	Symbol         string  `json:"symbol"`
	Side           string  `json:"side"`
	OrderType      string  `json:"order_type"`
	Price          float64 `json:"price"`
	Qty            float64 `json:"qty"`
	TimeInForce    string  `json:"time_in_force"`
	OrderStatus    string  `json:"order_status"`
	TriggerPrice   float64 `json:"trigger_price"`
	OrderLinkID    string  `json:"order_link_id"`
	CreatedTime    string  `json:"created_time"`
	UpdatedTime    string  `json:"updated_time"`
	BasePrice      string  `json:"base_price"`
	TriggerBy      string  `json:"trigger_by"`
	TpTriggerBy    string  `json:"tp_trigger_by"`
	SlTriggerBy    string  `json:"sl_trigger_by"`
	TakeProfit     float64 `json:"take_profit"`
	StopLoss       float64 `json:"stop_loss"`
	ReduceOnly     bool    `json:"reduce_only"`
	CloseOnTrigger bool    `json:"close_on_trigger"`
	PositionIdx    int     `json:"position_idx"`
}

type GetConditionalOrderParams struct {
	StopOrderID     string `url:"stop_order_id,omitempty" json:"stop_order_id,omitempty"`
	OrderLinkID     string `url:"order_link_id,omitempty" json:"order_link_id,omitempty"`
	Symbol          string `url:"symbol,omitempty" json:"symbol,omitempty"`
	StopOrderStatus string `url:"stop_order_status,omitempty" json:"stop_order_status,omitempty"`
	Order           string `url:"order_status,omitempty" json:"order_status,omitempty"`
	Page            int    `url:"page,omitempty" json:"page,omitempty"`
	Limit           int    `url:"limit,omitempty" json:"limit,omitempty"`
}

type GetConditionalOrderResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    GetConditionalOrderResult `json:"result"`
}

type GetConditionalOrderResult struct {
	CurrentPage int                             `json:"current_page"`
	LastPage    int                             `json:"last_page"`
	Data        []GetConditionalOrderResultData `json:"data"`
}

type GetConditionalOrderResultData struct {
	StopOrderID    string  `json:"stop_order_id"`
	UserID         string  `json:"order_id"`
	Symbol         string  `json:"symbol"`
	Side           string  `json:"side"`
	OrderType      string  `json:"order_type"`
	Price          float64 `json:"price"`
	Qty            float64 `json:"qty"`
	TimeInForce    string  `json:"time_in_force"`
	OrderStatus    string  `json:"order_status"`
	TriggerPrice   float64 `json:"trigger_price"`
	OrderLinkID    string  `json:"order_link_id"`
	CreatedTime    string  `json:"created_time"`
	UpdatedTime    string  `json:"updated_time"`
	TakeProfit     float64 `json:"take_profit"`
	StopLoss       float64 `json:"stop_loss"`
	TriggerBy      string  `json:"trigger_by"`
	BasePrice      string  `json:"base_price"`
	TpTriggerBy    string  `json:"tp_trigger_by"`
	SlTriggerBy    string  `json:"sl_trigger_by"`
	ReduceOnly     bool    `json:"reduce_only"`
	CloseOnTrigger bool    `json:"close_on_trigger"`
}

type CancelConditionalOrderParams struct {
	Symbol      string `url:"symbol" json:"symbol"`
	StopOrderID string `url:"stop_order_id,omitempty" json:"stop_order_id,omitempty"`
	OrderLinkID string `url:"order_link_id,omitempty" json:"order_link_id,omitempty"`
}

type CancelConditionalOrderResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    CancelConditionalOrderResult `json:"result"`
}

type CancelConditionalOrderResult struct {
	StopOrderID string `json:"stop_order_id"`
}

type CancelAllConditionalOrdersParams struct {
	Symbol string `url:"symbol" json:"symbol"`
}

type CancelAllConditionalOrdersResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    []string `json:"result"`
}

type ReplaceConditionalOrderParams struct {
	StopOrderID    string  `url:"stop_order_id,omitempty" json:"stop_order_id,omitempty"`
	OrderLinkID    string  `url:"order_link_id,omitempty" json:"order_link_id,omitempty"`
	Symbol         string  `url:"symbol,omitempty" json:"symbol,omitempty"`
	PRQty          float64 `url:"p_r_qty,omitempty" json:"p_r_qty,omitempty"`
	PRPrice        float64 `url:"p_r_price,omitempty" json:"p_r_price,omitempty"`
	PRTriggerPrice float64 `url:"p_r_trigger_price,omitempty" json:"p_r_trigger_price,omitempty"`
	TakeProfit     float64 `url:"take_profit,omitempty" json:"take_profit,omitempty"`
	StopLoss       float64 `url:"stop_loss,omitempty" json:"stop_loss,omitempty"`
	TpTriggerBy    string  `url:"tp_trigger_by,omitempty" json:"tp_trigger_by,omitempty"`
	SlTriggerBy    string  `url:"sl_trigger_by,omitempty" json:"sl_trigger_by,omitempty"`
}

type ReplaceConditionalOrderResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    ReplaceConditionalOrderResult `json:"result"`
}

type ReplaceConditionalOrderResult struct {
	StopOrderID string `json:"stop_order_id"`
}

type QueryConditionalOrderBySymbolParams struct {
	Symbol string `url:"symbol" json:"symbol"`
}

type QueryConditionalOrderBySymbolResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    []QueryConditionalOrderBySymbolResult `json:"result"`
}

type QueryConditionalOrderBySymbolResult struct {
	StopOrderID  string  `json:"stop_order_id"`
	UserID       int     `json:"user_id"`
	Symbol       string  `json:"symbol"`
	Side         string  `json:"side"`
	OrderType    string  `json:"order_type"`
	Price        float64 `json:"price"`
	Qty          float64 `json:"qty"`
	TimeInForce  string  `json:"time_in_force"`
	OrderStatus  string  `json:"order_status"`
	TriggerPrice float64 `json:"trigger_price"`
	OrderLinkID  string  `json:"order_link_id"`
	CreatedTime  string  `json:"created_time"`
	UpdatedTime  string  `json:"updated_time"`
	TakeProfit   float64 `json:"take_profit"`
	StopLoss     float64 `json:"stop_loss"`
	TpTriggerBy  string  `json:"tp_trigger_by"`
	SlTriggerBy  string  `json:"sl_trigger_by"`
	TriggerBy    string  `json:"trigger_by"`
}

type QueryConditionalOrderWithIDsParams struct {
	Symbol      string `url:"symbol" json:"symbol"`
	StopOrderID string `url:"stop_order_id,omitempty" json:"stop_order_id,omitempty"`
	OrderLinkID string `url:"order_link_id,omitempty" json:"order_link_id,omitempty"`
}

type QueryConditionalOrderWithIDsResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    QueryConditionalOrderWithIDsResult `json:"result"`
}

type QueryConditionalOrderWithIDsResult struct {
	StopOrderID    string  `json:"stop_order_id"`
	UserID         int     `json:"user_id"`
	Symbol         string  `json:"symbol"`
	Side           string  `json:"side"`
	OrderType      string  `json:"order_type"`
	Price          float64 `json:"price"`
	Qty            float64 `json:"qty"`
	TimeInForce    string  `json:"time_in_force"`
	OrderStatus    string  `json:"order_status"`
	TriggerPrice   float64 `json:"trigger_price"`
	BasePrice      string  `json:"base_price"`
	OrderLinkID    string  `json:"order_link_id"`
	CreatedTime    string  `json:"created_time"`
	UpdatedTime    string  `json:"updated_time"`
	TakeProfit     float64 `json:"take_profit"`
	StopLoss       float64 `json:"stop_loss"`
	TpTriggerBy    string  `json:"tp_trigger_by"`
	SlTriggerBy    string  `json:"sl_trigger_by"`
	TriggerBy      string  `json:"trigger_by"`
	ReduceOnly     bool    `json:"reduce_only"`
	CloseOnTrigger bool    `json:"close_on_trigger"`
}

type GetPositionsBySymbolParams struct {
	Symbol string `url:"symbol" json:"symbol"`
}

type GetPositionsBySymbolResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    []PositionsResult `json:"result"`
}

type PositionsResult struct {
	UserID              int     `json:"user_id"`
	Symbol              string  `json:"symbol"`
	Side                string  `json:"side"`
	Size                float64 `json:"size"`
	PositionValue       float64 `json:"position_value"`
	EntryPrice          float64 `json:"entry_price"`
	LiqPrice            float64 `json:"liq_price"`
	BustPrice           float64 `json:"bust_price"`
	Leverage            float64 `json:"leverage"`
	AutoAddMargin       int     `json:"auto_add_margin"`
	IsIsolated          bool    `json:"is_isolated"`
	PositionMargin      float64 `json:"position_margin"`
	OccClosingFee       float64 `json:"occ_closing_fee"`
	RealisedPnl         float64 `json:"realised_pnl"`
	CumRealisedPnl      float64 `json:"cum_realised_pnl"`
	FreeQty             float64 `json:"free_qty"`
	TpSlMode            string  `json:"tp_sl_mode"`
	UnrealisedPnl       float64 `json:"unrealised_pnl"`
	DeleverageIndicator int     `json:"deleverage_indicator"`
	RiskID              int     `json:"risk_id"`
	StopLoss            float64 `json:"stop_loss"`
	TakeProfit          float64 `json:"take_profit"`
	TrailingStop        float64 `json:"trailing_stop"`
	PositionIDx         int     `json:"position_idx"`
	Mode                string  `json:"mode"`
}

type GetPositionsResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    []PositionResultData `json:"result"`
}

type PositionResultData struct {
	Data    PositionsResult `json:"data"`
	IsValid bool            `json:"is_valid"`
}

type SetAutoAddMarginParams struct {
	Symbol        string     `url:"symbol" json:"symbol"`
	Side          bybit.Side `url:"side" json:"side"`
	AutoAddMargin bool       `url:"auto_add_margin" json:"auto_add_margin"`
	PositionIDx   int        `url:"position_idx,omitempty" json:"position_idx,omitempty"`
}

type SetAutoAddMarginResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    string `json:"result"`
}

type SwitchMarginParams struct {
	Symbol       string `url:"symbol" json:"symbol"`
	IsIsolated   bool   `url:"is_isolated" json:"is_isolated"`
	BuyLeverage  int    `url:"buy_leverage" json:"buy_leverage"`
	SellLeverage int    `url:"sell_leverage" json:"sell_leverage"`
}

type SwitchMarginResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    string `json:"result"`
}

type SwitchPositionModeParams struct {
	Symbol string `url:"symbol,omitempty" json:"symbol,omitempty"`
	Coin   string `url:"coin,omitempty" json:"coin,omitempty"`
	Mode   string `url:"mode" json:"mode"`
}

type SwitchPositionModeResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    string `json:"result"`
}

type PositionTpSlSwitchParams struct {
	Symbol   string `url:"symbol" json:"symbol"`
	TpSlMode string `url:"tp_sl_mode" json:"tp_sl_mode"`
}

type PositionTpSlSwitchResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    PositionTpSlSwitchResult `json:"result"`
}

type PositionTpSlSwitchResult struct {
	TpSlMode string `json:"tp_sl_mode"`
}

type AddReduceMarginParams struct {
	Symbol      string     `url:"symbol" json:"symbol"`
	Side        bybit.Side `url:"side" json:"side"`
	Margin      float64    `url:"margin" json:"margin"`
	PositionIDx int        `url:"position_idx,omitempty" json:"position_idx,omitempty"`
}

type AddReduceMarginResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    string `json:"result"`
}

type AddReduceMarginResult struct {
	PositionListResult PositionsResult `json:"PositionListResult"`
	WalletBalance      float64         `json:"wallet_balance"`
	AvailableBalance   float64         `json:"available_balance"`
}

type SetLeverageParams struct {
	Symbol       string `url:"symbol" json:"symbol"`
	BuyLeverage  int    `url:"buy_leverage" json:"buy_leverage"`
	SellLeverage int    `url:"sell_leverage" json:"sell_leverage"`
}

type SetLeverageResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    string `json:"result"`
}

type SetTradingStopParams struct {
	Symbol       string     `url:"symbol" json:"symbol"`
	Side         bybit.Side `url:"side" json:"side"`
	TakeProfit   float64    `url:"take_profit,omitempty" json:"take_profit,omitempty"`
	StopLoss     float64    `url:"stop_loss,omitempty" json:"stop_loss,omitempty"`
	TrailingStop float64    `url:"trailing_stop,omitempty" json:"trailing_stop,omitempty"`
	TpTriggerBy  string     `url:"tp_trigger_by,omitempty" json:"tp_trigger_by,omitempty"`
	SlTriggerBy  string     `url:"sl_trigger_by,omitempty" json:"sl_trigger_by,omitempty"`
	SlSize       float64    `url:"sl_size,omitempty" json:"sl_size,omitempty"`
	TpSize       float64    `url:"tp_size,omitempty" json:"tp_size,omitempty"`
	PositionIDx  int        `url:"position_idx,omitempty" json:"position_idx,omitempty"`
}

type SetTradingStopResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    string `json:"result"`
}

type GetUserTradeRecordsParams struct {
	Symbol    string `url:"symbol" json:"symbol"`
	StartTime int64  `url:"start_time,omitempty" json:"start_time,omitempty"`
	EndTime   int64  `url:"end_time,omitempty" json:"end_time,omitempty"`
	ExecType  string `url:"exec_type,omitempty" json:"exec_type,omitempty"`
	Page      int    `url:"page,omitempty" json:"page,omitempty"`
	Limit     int    `url:"limit,omitempty" json:"limit,omitempty"`
}

type GetUserTradeRecordsResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    UserTradeRecordResult `json:"result"`
}

type UserTradeRecordResult struct {
	CurrentPage int               `json:"current_page"`
	Data        []UserTradeRecord `json:"data"`
}

type GetExtendedUserTradeRecordsParams struct {
	Symbol    string `url:"symbol" json:"symbol"`
	StartTime int64  `url:"start_time,omitempty" json:"start_time,omitempty"`
	EndTime   int64  `url:"end_time,omitempty" json:"end_time,omitempty"`
	ExecType  string `url:"exec_type,omitempty" json:"exec_type,omitempty"`
	PageToken string `url:"page_token,omitempty" json:"page_token,omitempty"`
	Limit     int    `url:"limit,omitempty" json:"limit,omitempty"`
}

type GetExtendedUserTradeRecordsResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    GetExtendedUserTradeRecordsResult `json:"result"`
}

type GetExtendedUserTradeRecordsResult struct {
	CurrentPage int               `json:"current_page"`
	Data        []UserTradeRecord `json:"data"`
}

type UserTradeRecord struct {
	OrderID          string  `json:"order_id"`
	OrderLinkID      string  `json:"order_link_id"`
	Symbol           string  `json:"symbol"`
	Side             string  `json:"side"`
	ExecID           string  `json:"exec_id"`
	Price            float64 `json:"price"`
	OrderPrice       float64 `json:"order_price"`
	OrderQty         float64 `json:"order_qty"`
	OrderType        string  `json:"order_type"`
	FeeRate          float64 `json:"fee_rate"`
	ExecPrice        float64 `json:"exec_price"`
	ExecType         string  `json:"exec_type"`
	ExecQty          float64 `json:"exec_qty"`
	ExecFee          float64 `json:"exec_fee"`
	ExecValue        float64 `json:"exec_value"`
	LeavesQty        float64 `json:"leaves_qty"`
	ClosedSize       float64 `json:"closed_size"`
	LastLiquidityInd string  `json:"last_liquidity_ind"`
	TradeTime        int64   `json:"trade_time"`
	TradeTimeMs      int64   `json:"trade_time_ms"`
}

type GetClosedProfitLossParams struct {
	Symbol    string `url:"symbol" json:"symbol"`
	StartTime int64  `url:"start_time,omitempty" json:"start_time,omitempty"`
	EndTime   int64  `url:"end_time,omitempty" json:"end_time,omitempty"`
	ExecType  string `url:"exec_type,omitempty" json:"exec_type,omitempty"`
	Page      int    `url:"page,omitempty" json:"page,omitempty"`
	Limit     int    `url:"limit,omitempty" json:"limit,omitempty"`
}

type GetClosedProfitLossResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    ProfitAndLossResult `json:"result"`
}

type ProfitAndLossResult struct {
	CurrentPage int                       `json:"current_page"`
	Data        []ProfitAndLossResultData `json:"data"`
}

type ProfitAndLossResultData struct {
	ID                int64   `json:"id"`
	UserID            int64   `json:"user_id"`
	Symbol            string  `json:"symbol"`
	Side              string  `json:"side"`
	Qty               float64 `json:"qty"`
	OrderID           string  `json:"order_id"`
	OrderPrice        float64 `json:"order_price"`
	OrdrType          string  `json:"order_type"`
	ExecType          string  `json:"exec_type"`
	ClosedSize        float64 `json:"closed_size"`
	CumEntryValue     float64 `json:"cum_entry_value"`
	AverageEntryPrice float64 `json:"average_entry_price"`
	CumExitValue      float64 `json:"cum_exit_value"`
	AvgExitValue      float64 `json:"avg_exit_value"`
	ClosedPnl         float64 `json:"closed_pnl"`
	FillCount         int     `json:"fill_count"`
	Leverage          float64 `json:"leverage"`
	CreatedAt         int64   `json:"created_at"`
}

type GetRiskLimitParams struct {
	Symbol string `url:"symbol" json:"symbol"`
}

type GetRiskLimitResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    []GetRiskLimitResult `json:"result"`
}

type GetRiskLimitResult struct {
	ID             int64    `json:"id"`
	Symbol         string   `json:"symbol"`
	Limit          float64  `json:"limit"`
	MaintainMargin float64  `json:"maintain_margin"`
	StartingMargin float64  `json:"starting_margin"`
	Section        []string `json:"section"`
	IsLowestRisk   int      `json:"is_lowest_risk"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
	MaxLeverage    float64  `json:"max_leverage"`
}

type SetRiskLimitParams struct {
	Symbol      string     `url:"symbol" json:"symbol"`
	Side        bybit.Side `url:"side" json:"side"`
	RiskID      int64      `url:"risk_id" json:"risk_id"`
	PositionIDx int64      `url:"position_idx,omitempty" json:"position_idx,omitempty"`
}

type SetRiskLimitResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    SetRiskLimitResult `json:"result"`
}

type SetRiskLimitResult struct {
	RiskID int64 `json:"risk_id"`
}

type PredictedFundingRateParams struct {
	Symbol string `url:"symbol" json:"symbol"`
}

type PredictedFundingRateResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    PredictedFundingRateResult `json:"result"`
}

type PredictedFundingRateResult struct {
	PredictedFundingRate float64 `json:"predicted_funding_rate"`
	PredictedFundingFee  float64 `json:"predicted_funding_fee"`
}

type GetLastFundingFeeParams struct {
	Symbol string `url:"symbol" json:"symbol"`
}

type GetLastFundingFeeResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    GetLastFundingFeeResult `json:"result"`
}

type GetLastFundingFeeResult struct {
	Symbol      string  `json:"symbol"`
	Side        string  `json:"side"`
	Size        float64 `json:"size"`
	FundingRate float64 `json:"funding_rate"`
	ExecFee     float64 `json:"exec_fee"`
	ExecTime    string  `json:"exec_time"`
}

type GetAPIKeyInfoResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    []GetAPIKeyInfoResult `json:"result"`
}

type GetAPIKeyInfoResult struct {
	APIKey        string   `json:"api_key"`
	Type          string   `json:"type"`
	UserID        int64    `json:"user_id"`
	InviterID     int64    `json:"inviter_id"`
	IPs           []string `json:"ips"`
	Note          string   `json:"note"`
	Permissions   []string `json:"permissions"`
	CreatedAt     string   `json:"created_at"`
	ExpiredAt     string   `json:"expired_at"`
	ReadOnly      bool     `json:"read_only"`
	VipLevel      string   `json:"vip_level"`
	MktMakerLevel string   `json:"mkt_maker_level"`
}

type GetLCPInfoParams struct {
	Symbol string `url:"symbol" json:"symbol"`
}

type GetLCPInfoResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    GetLCPInfoResult `json:"result"`
}

type GetLCPInfoResult struct {
	LCPList []LCPData `json:"lcp_list"`
}

type LCPData struct {
	Date          string  `json:"date"`
	SelfRatio     float64 `json:"self_ratio"`
	PlatformRatio float64 `json:"platform_ratio"`
	Score         float64 `json:"score"`
}
