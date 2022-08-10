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
	PublicLinearOrderBookPath             = "v2/public/orderBook/L2"
	PublicQueryKlinePath                  = "/public/linear/kline"
	PublicGetSymbolInformationPath        = "/v2/public/tickers"
	PublicGetPublicTradingRecordsPath     = "/public/linear/recent-trading-records"
	PublicGetLastFundingRatePath          = "/public/linear/funding/prev-funding-rate"
	PublicQuerySymbolPath                 = "/v2/public/symbols"
	PublicLiquidatedOrdersPath            = "/v2/public/liq-records"
	PublicQueryMarkPriceKline             = "/public/linear/mark-price-kline"
	PublicQueryIndexPriceKlinePath        = "/v2/public/index-price-kline"
	PublicQueryPremiumIndexPriceKlinePath = "/v2/public/premium-index-kline"
	PublicGetOpenInterestPath             = "/v2/public/open-interest"
	PublicGetLatestBigDealPath            = "/v2/public/big-deal"
	PublicGetLongShortRatioPath           = "/v2/public/account-ratio"

	// ACCOUNT DATA ENDPOINTS
	// Inverse Perpetual
	PrivateInversePerpOrderCreatePath  = "/v2/private/order/create"
	PrivateInversePerpOrderCancelPath  = "/v2/private/order/cancel"
	PrivateInversePerpPositionListPath = "/v2/private/position/list"
	PrivateInversePerpLeverageSavePath = "/v2/private/position/leverage/save"

	// USDT Perpetual
	PrivateLinearPlaceOrderPath          = "/private/linear/order/create"
	PrivateLinearGetActiveOrderPath      = "/private/linear/order/list"
	PrivateLinearOrderCancelPath         = "/private/linear/order/cancel"
	PrivateLinearOrderCancelAllPath      = "/private/linear/order/cancel-all"
	PrivateLinearReplaceActiveOrderPath  = "/private/linear/order/replace"
	PrivateLinearQueryActiveOrderPath    = "/private/linear/order/search"
	PrivateLinearPositionListPath        = "/private/linear/position/list"
	PrivateLinearPositionSetLeveragePath = "/private/linear/position/set-leverage"
	PrivateLinearTradeExecutionListPath  = "/private/linear/trade/execution/list"
)
