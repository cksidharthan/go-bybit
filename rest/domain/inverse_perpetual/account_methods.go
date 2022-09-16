package inverseperp

import "context"

type AccountInterface interface {
	// Active Orders
	PlaceActiveOrder(ctx context.Context, params *PlaceActiveOrderParams) (*PlaceActiveOrderResponse, error)
	GetActiveOrder(ctx context.Context, params *GetActiveOrderParams) (*GetActiveOrderResponse, error)
	CancelActiveOrder(ctx context.Context, params *CancelActiveOrderParams) (*CancelActiveOrderResponse, error)
	CancelAllActiveOrders(ctx context.Context, params *CancelAllActiveOrdersParams) (*CancelAllActiveOrdersResponse, error)
	ReplaceActiveOrder(ctx context.Context, params *ReplaceActiveOrderParams) (*ReplaceActiveOrderResponse, error)
	QueryActiveOrder(ctx context.Context, params *QueryActiveOrderParams) (*QueryActiveOrderResponse, error)

	// Conditional Orders
	PlaceConditionalOrder(ctx context.Context, params *PlaceConditionalOrderParams) (*PlaceConditionalOrderResponse, error)
	GetConditionalOrder(ctx context.Context, params *GetConditionalOrderParams) (*GetConditionalOrderResponse, error)
	CancelConditionalOrder(ctx context.Context, params *CancelConditionalOrderParams) (*CancelConditionalOrderResponse, error)
	CancelAllConditionalOrders(ctx context.Context, params *CancelAllConditionalOrdersParams) (*CancelAllConditionalOrdersResponse, error)
	ReplaceConditionalOrder(ctx context.Context, params *ReplaceConditionalOrderParams) (*ReplaceConditionalOrderResponse, error)
	QueryConditionalOrder(ctx context.Context, params *QueryConditionalOrderParams) (*QueryConditionalOrderResponse, error)

	// Position
	GetPosition(ctx context.Context, params *GetPositionParams) (*GetPositionResponse, error)
	ChangeMargin(ctx context.Context, params *ChangeMarginParams) (*ChangeMarginResponse, error)
	SetTradingStop(ctx context.Context, params *SetTradingStopParams) (*SetTradingStopResponse, error)
	SetLeverage(ctx context.Context, params *SetLeverageParams) (*SetLeverageResponse, error)
	GetUserTradeRecords(ctx context.Context, params *GetUserTradeRecordsParams) (*GetUserTradeRecordsResponse, error)
	ClosedProfitAndLoss(ctx context.Context, params *ClosedProfitAndLossParams) (*ClosedProfitAndLossResponse, error)
	PositionTpSlSwitch(ctx context.Context, params *PositionTpSlSwitchParams) (*PositionTpSlSwitchResponse, error)
	MarginSwitch(ctx context.Context, params *MarginSwitchParams) (*MarginSwitchResponse, error)
	QueryTradingFeeRate(ctx context.Context, params *QueryTradingFeeRateParams) (*QueryTradingFeeRateResponse, error)

	// Risk Limit
	GetRiskLimit(ctx context.Context, params *GetRiskLimitParams) (*GetRiskLimitResponse, error)
	SetRiskLimit(ctx context.Context, params *SetRiskLimitParams) (*SetRiskLimitResponse, error)

	// Funding
	GetFundingRate(ctx context.Context, params *GetFundingRateParams) (*GetFundingRateResponse, error)
	GetLastFundingFee(ctx context.Context, params *GetLastFundingFeeParams) (*GetLastFundingFeeResponse, error)
	GetPredictedFunding(ctx context.Context, params *GetPredictedFundingParams) (*GetPredictedFundingResponse, error)

	GetAPIKeyInfo(ctx context.Context, params *GetAPIKeyInfoParams) (*GetAPIKeyInfoResponse, error)
	GetLCPInfo(ctx context.Context, params *GetLCPInfoParams) (*GetLCPInfoResponse, error)
}
