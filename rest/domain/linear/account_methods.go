package linear

import "context"

type AccountInterface interface {
	PlaceActiveOrder(ctx context.Context, params *PlaceActiveOrderParams) (*PlaceActiveOrderResponse, error)
	GetActiveOrder(ctx context.Context, params *GetActiveOrderParams) (*GetActiveOrderResponse, error)
	CancelActiveOrder(ctx context.Context, params *CancelActiveOrderParams) (*CancelActiveOrderResponse, error)
	CancelAllActiveOrders(ctx context.Context, params *CancelAllActiveOrdersParams) (*CancelAllActiveOrdersResponse, error)
	ReplaceActiveOrder(ctx context.Context, params *ReplaceActiveOrderParams) (*ReplaceActiveOrderResponse, error)
	QueryActiveOrder(ctx context.Context, params *QueryActiveOrderParams) (*QueryActiveOrderResponse, error)

	PlaceConditionalOrder(ctx context.Context, params *PlaceConditionalOrderParams) (*PlaceConditionalOrderResponse, error)
	GetConditionalOrder(ctx context.Context, params *GetConditionalOrderParams) (*GetConditionalOrderResponse, error)
	CancelConditionalOrder(ctx context.Context, params *CancelConditionalOrderParams) (*CancelConditionalOrderResponse, error)
	CancelAllConditionalOrders(ctx context.Context, params *CancelAllConditionalOrdersParams) (*CancelAllConditionalOrdersResponse, error)
	ReplaceConditionalOrder(ctx context.Context, params *ReplaceConditionalOrderParams) (*ReplaceConditionalOrderResponse, error)
	QueryConditionalOrderBySymbol(ctx context.Context, params *QueryConditionalOrderBySymbolParams) (*QueryConditionalOrderBySymbolResponse, error)
	QueryConditionalOrderWithIDs(ctx context.Context, params *QueryConditionalOrderWithIDsParams) (*QueryConditionalOrderWithIDsResponse, error)

	GetPositionsBySymbol(ctx context.Context, params *GetPositionsBySymbolParams) (*GetPositionsBySymbolResponse, error)
	GetPositions(ctx context.Context) (*GetPositionsResponse, error)
	SetAutoAddMargin(ctx context.Context, params *SetAutoAddMarginParams) (*SetAutoAddMarginResponse, error)
	SwitchMargin(ctx context.Context, params *SwitchMarginParams) (*SwitchMarginResponse, error)
	SwitchPositionMode(ctx context.Context, params *SwitchPositionModeParams) (*SwitchPositionModeResponse, error)
	PositionTpSlSwitch(ctx context.Context, params *PositionTpSlSwitchParams) (*PositionTpSlSwitchResponse, error)
	AddReduceMargin(ctx context.Context, params *AddReduceMarginParams) (*AddReduceMarginResponse, error)
	SetLeverage(ctx context.Context, params *SetLeverageParams) (*SetLeverageResponse, error)
	SetTradingStop(ctx context.Context, params *SetTradingStopParams) (*SetTradingStopResponse, error)
	GetUserTradeRecords(ctx context.Context, params *GetUserTradeRecordsParams) (*GetUserTradeRecordsResponse, error)
	GetExtendedUserTradeRecords(ctx context.Context, params *GetExtendedUserTradeRecordsParams) (*GetExtendedUserTradeRecordsResponse, error)
	GetClosedProfitLoss(ctx context.Context, params *GetClosedProfitLossParams) (*GetClosedProfitLossResponse, error)
}
