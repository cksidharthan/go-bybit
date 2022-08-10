package bybit

type Symbol string

type Side string

const (
	SideBuy      Side = "Buy"
	SideSell     Side = "Sell"
	SideSpotBuy  Side = "BUY"
	SideSpotSell Side = "SELL"
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

// Interval :
type Interval string

const (
	// Interval1Min :
	Interval1Min = Interval("1")
	// Interval3Min :
	Interval3Min = Interval("3")
	// Interval5Min :
	Interval5Min = Interval("5")
	// Interval15Min :
	Interval15Min = Interval("15")
	// Interval30Min :
	Interval30Min = Interval("30")
	// Interval60Min :
	Interval60Min = Interval("60")
	// Interval120Min :
	Interval120Min = Interval("120")
	// Interval240Min :
	Interval240Min = Interval("240")
	// Interval360Min :
	Interval360Min = Interval("360")
	// Interval720Min :
	Interval720Min = Interval("720")
	// IntervalDay :
	IntervalDay = Interval("D")
	// IntervalWeek :
	IntervalWeek = Interval("W")
	// IntervalMonth :
	IntervalMonth = Interval("M")
)

type Period string

const (
	// Period5Min :
	Period5Min = Period("5m")
	// Period15Min :
	Period15Min = Period("15m")
	// Period30Min :
	Period30Min = Period("30m")
	// Period1hr :
	Period1hr = Period("1h")
	// Period4hr :
	Period4hr = Period("4h")
	// Period1day :
	Period1day = Period("1d")
)
