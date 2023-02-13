package account

import "github.com/cksidharthan/go-bybit/rest/v5/domain"

type GetWalletBalanceParams struct {
	// Account type - Unified account: UNIFIED Or Normal account: CONTRACT
	AccountType string `json:"accountType"`
	// Coin name - If not passed, it returns non-zero asset info. You can pass multiple coins to query, separated by comma. USDT,USDC.
	Coin string `json:"coin,omitempty"`
}

type GetWalletBalanceResponse struct {
	domain.BaseResponse
	Result GetWalletBalanceResponseResult `json:"result"`
}

type GetWalletBalanceResponseResult struct {
	List []GetWalletBalanceResponseResultList `json:"list"`
}

type GetWalletBalanceResponseResultList struct {
	TotalEquity            float64                                  `json:"totalEquity"`
	AccountIMRate          float64                                  `json:"accountIMRate"`
	TotalMarginBalance     float64                                  `json:"totalMarginBalance"`
	TotalInitialMargin     float64                                  `json:"totalInitialMargin"`
	AccountType            string                                   `json:"accountType"`
	TotalAvailableBalance  float64                                  `json:"totalAvailableBalance"`
	AccountMMRate          float64                                  `json:"accountMMRate"`
	TotalPerpUPL           float64                                  `json:"totalPerpUPL"`
	TotalWalletBalance     float64                                  `json:"totalWalletBalance"`
	TotalMaintenanceMargin float64                                  `json:"totalMaintenanceMargin"`
	Coin                   []GetWalletBalanceResponseResultListCoin `json:"coin"`
}

type GetWalletBalanceResponseResultListCoin struct {
	AvailableToBorrow   float64 `json:"availableToBorrow"`
	AccruedInterest     float64 `json:"accruedInterest"`
	AvailableToWithdraw float64 `json:"availableToWithdraw"`
	TotalOrderIM        float64 `json:"totalOrderIM"`
	Equity              float64 `json:"equity"`
	TotalPositionMM     float64 `json:"totalPositionMM"`
	USDValue            float64 `json:"usdValue"`
	UnrealizedPNL       float64 `json:"unrealizedPnl"`
	BorrowAmount        float64 `json:"borrowAmount"`
	TotalPositionIM     float64 `json:"totalPositionIM"`
	WalletBalance       float64 `json:"walletBalance"`
	CumRealizedPNL      float64 `json:"cumRealizedPnl"`
	Coin                string  `json:"coin"`
}
