package bybit

type Symbol string

type Side string

const (
	SideBuy  Side = "Buy"
	SideSell Side = "Sell"
)

type OrderType string

const (
	OrderTypeLimit       OrderType = "Limit"
	OrderTypeMarket      OrderType = "Market"
	OrderTypeConditional OrderType = "Conditional"
)

type OrderTypeSpot string

const (
	OrderTypeSpotLimit      OrderTypeSpot = "LIMIT"
	OrderTypeSpotMarket     OrderTypeSpot = "MARKET"
	OrderTypeSpotLimitMaker OrderTypeSpot = "LIMIT_MAKER"
)

// TimeInForce :
type TimeInForce string

const (
	// TimeInForceGoodTillCancel :
	TimeInForceGoodTillCancel = TimeInForce("GoodTillCancel")
	// TimeInForceImmediateOrCancel :
	TimeInForceImmediateOrCancel = TimeInForce("ImmediateOrCancel")
	// TimeInForceFillOrKill :
	TimeInForceFillOrKill = TimeInForce("FillOrKill")
	// TimeInForcePostOnly :
	TimeInForcePostOnly = TimeInForce("PostOnly")
)

// TimeInForceSpot :
type TimeInForceSpot string

const (
	// TimeInForceSpotGTC :
	TimeInForceSpotGTC = TimeInForceSpot("GTC")
	// TimeInForceSpotFOK :
	TimeInForceSpotFOK = TimeInForceSpot("FOK")
	// TimeInForceSpotIOC :
	TimeInForceSpotIOC = TimeInForceSpot("IOC")
)

// OrderStatus :
type OrderStatus string

const (
	// OrderStatusCreated :
	OrderStatusCreated = OrderStatus("Created")
	// OrderStatusRejected :
	OrderStatusRejected = OrderStatus("Rejected")
	// OrderStatusNew :
	OrderStatusNew = OrderStatus("New")
	// OrderStatusPartiallyFilled :
	OrderStatusPartiallyFilled = OrderStatus("PartiallyFilled")
	// OrderStatusFilled :
	OrderStatusFilled = OrderStatus("Filled")
	// OrderStatusCancelled :
	OrderStatusCancelled = OrderStatus("Cancelled")
	// OrderStatusPendingCancel :
	OrderStatusPendingCancel = OrderStatus("PendingCancel")
)

// OrderStatusSpot :
type OrderStatusSpot string

const (
	// OrderStatusSpotNew :
	OrderStatusSpotNew = OrderStatusSpot("NEW")
	// OrderStatusSpotPartiallyFilled :
	OrderStatusSpotPartiallyFilled = OrderStatusSpot("PARTIALLY_FILLED")
	// OrderStatusSpotFilled :
	OrderStatusSpotFilled = OrderStatusSpot("FILLED")
	// OrderStatusSpotCanceled :
	OrderStatusSpotCanceled = OrderStatusSpot("CANCELED")
	// OrderStatusSpotPendingCancel :
	OrderStatusSpotPendingCancel = OrderStatusSpot("PENDING_CANCEL")
	// OrderStatusSpotPendingNew :
	OrderStatusSpotPendingNew = OrderStatusSpot("PENDING_NEW")
	// OrderStatusSpotRejected :
	OrderStatusSpotRejected = OrderStatusSpot("REJECTED")
)
