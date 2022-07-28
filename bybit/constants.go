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

	// MARKET DATA
	PublicMarketDataOrderBookL2Path       = "/v2/public/orderBook/L2"
	PublicMarketDataKlineListPath         = "/v2/public/kline/list"
	PublicMarketDataTickersPath           = "/v2/public/tickers"
	PublicMarketDataTradingRecordsPath    = "/v2/public/trading-records"
	PublicMarketDataSymbolsPath           = "/v2/public/symbols"
	PublicMarketDataMarkPriceKlinePath    = "/v2/public/mark-price-kline"
	PublicMarketDataIndexPriceKlinePath   = "/v2/public/index-price-kline"
	PublicMarketDataPremiumIndexKlinePath = "/v2/public/premium-index-kline"
	PublicMarketDataOpenInterestPath      = "/v2/public/open-interest"
	PublicMarketDataBigDealPath           = "/v2/public/big-deal"
	PublicMarketDataAccountRatioPath      = "/v2/public/account-ratio"

	// ACCOUNT DATA ENDPOINTS
	// Inverse Perpetual
	PrivateInversePerpOrderCreatePath  = "/v2/private/order/create"
	PrivateInversePerpOrderCancelPath  = "/v2/private/order/cancel"
	PrivateInversePerpPositionListPath = "/v2/private/position/list"
	PrivateInversePerpLeverageSavePath = "/v2/private/position/leverage/save"

	// USDT Perpetual
	PrivateUSDTPerpOrderCreatePath         = "/private/linear/order/create"
	PrivateUSDTPerpOrderCancelPath         = "/private/linear/order/cancel"
	PrivateUSDTPerpPositionListPath        = "/private/linear/position/list"
	PrivateUSDTPerpPositionSetLeveragePath = "/private/linear/position/set-leverage"
	PrivateUSDTPerpTradeExecutionListPath  = "/private/linear/trade/execution/list"

	// WALLET_DATA
	PrivateWalletBalancePath = "/v2/private/wallet/balance"
)
