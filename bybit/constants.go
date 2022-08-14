package bybit

const (
	BybitMainnetBaseURL  = "https://api.bybit.com"
	BytickMainnetBaseURL = "https://api.bytick.com"
	BybitTestnetBaseURL  = "https://api-testnet.bybit.com"

	/* Spot Market API */
	PublicServerTimePath                = "/spot/v1/time"
	PublicSpotSymbolsPath               = "/spot/v1/symbols"
	PublicSpotQuoteDepthPath            = "/spot/quote/v1/depth"
	PublicSpotQuoteDepthMergedPath      = "/spot/quote/v1/depth/merged"
	PublicSpotQuoteTradesPath           = "/spot/quote/v1/trades"
	PublicSpotQuoteKlinePath            = "/spot/quote/v1/kline"
	PublicSpotQuoteTicker24HrPath       = "/spot/quote/v1/ticker/24hr"
	PublicSpotQuoteTickerPrice          = "/spot/quote/v1/ticker/price"
	PublicSpotQuoteTickerBookTickerPath = "/spot/quote/v1/ticker/book_ticker"

	/* Spot Account API */
	PrivateSpotOrderPath                = "/spot/v1/order"
	PrivateSpotOrderFastPath            = "/spot/v1/order/fast"
	PrivateSpotOrderBatchCancelPath     = "/spot/order/batch-cancel"
	PrivateSpotOrderBatchFastCancelPath = "/spot/order/batch-fast-cancel"
	PrivateSpotBatchCancelByIdsPath     = "/spot/order/batch-cancel-by-ids"
	PrivateOpenOrdersPath               = "/spot/v1/open-orders"
	PrivateOrderHistoryPath             = "/spot/v1/history-orders"
	PrivateTradeHistoryPath             = "/spot/v1/myTrades"

	/* Spot Wallet API */
	PrivateWalletBalancePath = "/spot/v1/account"

	/* USDT Perpetual Market API */
	PublicLinearOrderBookPath                   = "v2/public/orderBook/L2"
	PublicLinearQueryKlinePath                  = "/public/linear/kline"
	PublicLinearGetSymbolInformationPath        = "/v2/public/tickers"
	PublicLinearGetPublicTradingRecordsPath     = "/public/linear/recent-trading-records"
	PublicLinearGetLastFundingRatePath          = "/public/linear/funding/prev-funding-rate"
	PublicLinearQuerySymbolPath                 = "/v2/public/symbols"
	PublicLinearLiquidatedOrdersPath            = "/v2/public/liq-records"
	PublicLinearQueryMarkPriceKline             = "/public/linear/mark-price-kline"
	PublicLinearQueryIndexPriceKlinePath        = "/v2/public/index-price-kline"
	PublicLinearQueryPremiumIndexPriceKlinePath = "/v2/public/premium-index-kline"
	PublicLinearGetOpenInterestPath             = "/v2/public/open-interest"
	PublicLinearGetLatestBigDealPath            = "/v2/public/big-deal"
	PublicLinearGetLongShortRatioPath           = "/v2/public/account-ratio"

	// ACCOUNT DATA ENDPOINTS
	// Inverse Perpetual
	PrivateInversePerpOrderCreatePath  = "/v2/private/order/create"
	PrivateInversePerpOrderCancelPath  = "/v2/private/order/cancel"
	PrivateInversePerpPositionListPath = "/v2/private/position/list"
	PrivateInversePerpLeverageSavePath = "/v2/private/position/leverage/save"

	// USDT Perpetual
	PrivateLinearPlaceOrderPath                  = "/private/linear/order/create"
	PrivateLinearGetActiveOrderPath              = "/private/linear/order/list"
	PrivateLinearOrderCancelPath                 = "/private/linear/order/cancel"
	PrivateLinearOrderCancelAllPath              = "/private/linear/order/cancel-all"
	PrivateLinearReplaceActiveOrderPath          = "/private/linear/order/replace"
	PrivateLinearQueryActiveOrderPath            = "/private/linear/order/search"
	PrivateLinearPlaceConditionalOrderPath       = "/private/linear/stop-order/create"
	PrivateLinearGetConditionalOrderPath         = "/private/linear/stop-order/list"
	PrivateLinearCancelConditionalOrderPath      = "/private/linear/stop-order/cancel"
	PrivateLinearCancelAllConditionalOrderPath   = "/private/linear/stop-order/cancel-all"
	PrivateLinearReplaceConditionalOrderPath     = "/private/linear/stop-order/replace"
	PrivateLinearQueryConditionalOrderPath       = "/private/linear/stop-order/search"
	PrivateLinearGetPositionPath                 = "/private/linear/position/list"
	PrivateLinearSetAutoAddMarginPath            = "/private/linear/position/set-auto-add-margin"
	PrivateLinearSwitchMarginPath                = "/private/linear/position/switch-isolated"
	PrivateLinearSwitchPositionModePath          = "/private/linear/position/switch-mode"
	PrivateLinearPositionTpSlSwitch              = "/private/linear/tpsl/switch-mode"
	PrivateLinearAddReduceMarginPath             = "/private/linear/position/add-margin"
	PrivateLinearSetLeveragePath                 = "/private/linear/position/set-leverage"
	PrivateLinearSetTradingStopPath              = "/private/linear/position/trading-stop"
	PrivateLinearGetUserTradeRecordsPath         = "/private/linear/trade/execution/list"
	PrivateLinearGetExtendedUserTradeRecordsPath = "/private/linear/trade/execution/history-list"
	PrivateLinearGetClosedProfitLossPath         = "/private/linear/trade/closed-pnl/list"
)
