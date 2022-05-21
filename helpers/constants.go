package helpers

const (
	BYBIT_MAINNET_BASE_URL  = "https://api.bybit.com"
	BYTICK_MAINNET_BASE_URL = "https://api.bytick.com"
	BYBIT_TESTNET_BASE_URL  = "https://api-testnet.bybit.com"

	PUBLIC_SERVER_TIME_PATH = "/spot/v1/time"

	// MARKET DATA
	PUBLIC_MARKET_DATA_ORDER_BOOK_L2_PATH       = "/v2/public/orderBook/L2"
	PUBLIC_MARKET_DATA_KLINE_LIST_PATH          = "/v2/public/kline/list"
	PUBLIC_MARKET_DATA_TICKERS_PATH             = "/v2/public/tickers"
	PUBLIC_MARKET_DATA_TRADING_RECORDS_PATH     = "/v2/public/trading-records"
	PUBLIC_MARKET_DATA_SYMBOLS_PATH             = "/v2/public/symbols"
	PUBLIC_MARKET_DATA_MARK_PRICE_KLINE_PATH    = "/v2/public/mark-price-kline"
	PUBLIC_MARKET_DATA_INDEX_PRICE_KLINE_PATH   = "/v2/public/index-price-kline"
	PUBLIC_MARKET_DATA_PREMIUM_INDEX_KLINE_PATH = "/v2/public/premium-index-kline"
	PUBLIC_MARKET_DATA_OPEN_INTEREST_PATH       = "/v2/public/open-interest"
	PUBLIC_MARKET_DATA_BIG_DEAL_PATH            = "/v2/public/big-deal"
	PUBLIC_MARKET_DATA_ACCOUNT_RATIO_PATH       = "/v2/public/account-ratio"

	// SPOT
	PUBLIC_SPOT_SYMBOLS_PATH                  = "/spot/v1/symbols"
	PUBLIC_SPOT_QUOTE_DEPTH_PATH              = "/spot/quote/v1/depth"
	PUBLIC_SPOT_QUOTE_DEPTH_MERGED_PATH       = "/spot/quote/v1/depth/merged"
	PUBLIC_SPOT_QUOTE_TRADES_PATH             = "/spot/quote/v1/trades"
	PUBLIC_SPOT_QUOTE_KLINE_PATH              = "/spot/quote/v1/kline"
	PUBLIC_SPOT_QUOTE_TICKER_24HR_PATH        = "/spot/quote/v1/ticker/24hr"
	PUBLIC_SPOT_QUOTE_TICKER_PRICE            = "/spot/quote/v1/ticker/price"
	PUBLIC_SPOT_QUOTE_TICKER_BOOK_TICKER_PATH = "/spot/quote/v1/ticker/book_ticker"

	// ACCOUNT DATA ENDPOINTS
	// Inverse Perpetual
	PRIVATE_INVERSE_PERP_ORDER_CREATE_PATH  = "/v2/private/order/create"
	PRIVATE_INVERSE_PERP_ORDER_CANCEL_PATH  = "/v2/private/order/cancel"
	PRIVATE_INVERSE_PERP_POSITION_LIST_PATH = "/v2/private/position/list"
	PRIVATE_INVERSE_PERP_LEVERAGE_SAVE_PATH = "/v2/private/position/leverage/save"

	// USDT Perpetual
	PRIVATE_USDT_PERP_ORDER_CREATE_PATH          = "/private/linear/order/create"
	PRIVATE_USDT_PERP_ORDER_CANCEL_PATH          = "/private/linear/order/cancel"
	PRIVATE_USDT_PERP_POSITION_LIST_PATH         = "/private/linear/position/list"
	PRIVATE_USDT_PERP_POSITION_SET_LEVERAGE_PATH = "/private/linear/position/set-leverage"
	PRIVATE_USDT_PERP_TRADE_EXECUTION_LIST_PATH  = "/private/linear/trade/execution/list"

	// SPOT
	PRIVATE_SPOT_ORDER_PATH                  = "/spot/v1/order"
	PRIVATE_SPOT_ORDER_FAST_PATH             = "/spot/v1/order/fast"
	PRIVATE_SPOT_ORDER_BATCH_CANCEL_PATH     = "/spot/order/batch-cancel"
	PRIVATE_SPOT_ORDE_BATCH_FAST_CANCEL_PATH = "/spot/order/batch-fast-cancel"
	PRIVATE_SPOT_BATCH_CANCEL_BY_IDS_PATH    = "/spot/order/batch-cancel-by-ids"

	// WALLET_DATA
	PRIVATE_WALLET_BALANCE_PATH = "/v2/private/wallet/balance"
)
