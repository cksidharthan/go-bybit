package linear

import "github.com/cksidharthan/go-bybit/rest/domain"

type GetWalletBalanceParams struct {
	Coin string `json:"coin"`
}

type GetWalletBalanceResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    map[string]GetWalletBalanceData `json:"result"`
}

type GetWalletBalanceData struct {
	Equity           float64 `json:"equity"`
	AvailableBalance float64 `json:"available_balance"`
	UsedMargin       float64 `json:"used_margin"`
	OrderMargin      float64 `json:"order_margin"`
	PositionMargin   float64 `json:"position_margin"`
	OccClosingFee    float64 `json:"occ_closing_fee"`
	OccFundingFee    float64 `json:"occ_funding_fee"`
	WalletBalance    float64 `json:"wallet_balance"`
	RealisedPnl      float64 `json:"realised_pnl"`
	UnrealisedPnl    float64 `json:"unrealised_pnl"`
	CumRealisedPnl   float64 `json:"cum_realised_pnl"`
	GivenCash        float64 `json:"given_cash"`
	ServiceCash      float64 `json:"service_cash"`
}

type GetWalletFundRecordsParams struct {
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	Currency       string `json:"currency"`
	Coin           string `json:"coin"`
	WalletFundType string `json:"wallet_fund_type"`
	Page           int    `json:"page"`
	Limit          int    `json:"limit"`
}

type GetWalletFundRecordsResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    GetWalletFundRecordsResult `json:"result"`
}

type GetWalletFundRecordsResult struct {
	Data []GetWalletFundRecordsData `json:"data"`
}

type GetWalletFundRecordsData struct {
	ID            int     `json:"id"`
	UserID        int     `json:"user_id"`
	Coin          string  `json:"coin"`
	WalletID      int     `json:"wallet_id"`
	Type          string  `json:"type"`
	Amount        float64 `json:"amount"`
	TxID          string  `json:"tx_id"`
	Address       string  `json:"address"`
	WalletBalance string  `json:"wallet_balance"`
	ExecTime      string  `json:"exec_time"`
	CrossSeq      int     `json:"cross_seq"`
}

type GetWithdrawRecordsParams struct {
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	Coin      string `json:"coin,omitempty"`
	Status    string `json:"status,omitempty"`
	Page      int    `json:"page,omitempty"`
	Limit     int    `json:"limit,omitempty"`
}

type GetWithdrawRecordsResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    GetWithdrawRecordsResult `json:"result"`
}

type GetWithdrawRecordsResult struct {
	Data        []GetWithdrawRecordsData `json:"data"`
	CurrentPage int                      `json:"current_page"`
	LastPage    int                      `json:"last_page"`
}

type GetWithdrawRecordsData struct {
	ID          int     `json:"id"`
	UserID      int64   `json:"user_id"`
	Coin        string  `json:"coin"`
	Status      string  `json:"status"`
	Amount      float64 `json:"amount"`
	Fee         float64 `json:"fee"`
	Address     string  `json:"address"`
	TxID        string  `json:"tx_id"`
	SubmittedAt string  `json:"submitted_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type GetAssetExchangeRecordsParams struct {
	Limit     int    `json:"limit,omitempty"`
	From      int    `json:"from,omitempty"`
	Direction string `json:"direction,omitempty"`
}

type GetAssetExchangeRecordsResponse struct {
	domain.LinearBaseResponse `json:",inline"`
	Result                    []GetAssetExchangeRecordsResult `json:"result"`
}

type GetAssetExchangeRecordsResult struct {
	ID           int     `json:"id"`
	ExchangeRate float64 `json:"exchange_rate"`
	FromCoin     string  `json:"from_coin"`
	ToCoin       string  `json:"to_coin"`
	ToAmount     float64 `json:"to_amount"`
	FromFee      float64 `json:"from_fee"`
	FromAmount   float64 `json:"from_amount"`
	CreatedAt    string  `json:"created_at"`
}
