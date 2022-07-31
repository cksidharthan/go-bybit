package spot

import "github.com/cksidharthan/go-bybit/rest/domain"

type GetWalletBalanceResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              Balances `json:"result"`
}

type Balances struct {
	Balances []GetWalletBalanceResult `json:"balances"`
}

type GetWalletBalanceResult struct {
	Coin     string `json:"coin"`
	CoinID   string `json:"coinId"`
	CoinName string `json:"coinName"`
	Total    string `json:"total"`
	Free     string `json:"free"`
	Locked   string `json:"locked"`
}
