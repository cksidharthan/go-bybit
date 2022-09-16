package inverseperp

import "github.com/cksidharthan/go-bybit/rest/domain"

type WalletBalanceParams struct {
	Coin string `url:"coin" json:"coin,omitempty" description:"currency alias. Returns all wallet balances if not passed."`
}

type WalletBalanceResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              map[string]WalletBalanceResult `json:"result"`
}

type WalletBalanceResult struct {
	Equity           float64 `url:"equity" json:"equity"`
	AvailableBalance float64 `url:"available_balance" json:"available_balance"`
	UsedMargin       float64 `url:"used_margin" json:"used_margin"`
	OrderMargin      float64 `url:"order_margin" json:"order_margin"`
	PositionMargin   float64 `url:"position_margin" json:"position_margin"`
	OccClosingFee    float64 `url:"occ_closing_fee" json:"occ_closing_fee"`
	OccFundingFee    float64 `url:"occ_funding_fee" json:"occ_funding_fee"`
	WalletBalance    float64 `url:"wallet_balance" json:"wallet_balance"`
	RealisedPnl      float64 `url:"realised_pnl" json:"realised_pnl"`
	UnrealisedPnl    float64 `url:"unrealised_pnl" json:"unrealised_pnl"`
	CumRealisedPnl   float64 `url:"cum_realised_pnl" json:"cum_realised_pnl"`
	GivenCash        float64 `url:"given_cash" json:"given_cash"`
	ServiceCash      float64 `url:"service_cash" json:"service_cash"`
}

type WalletFundRecordsParams struct {
	StartDate      string `url:"start_date" json:"start_date,omitempty" description:"start date of the query"`
	EndDate        string `url:"end_date" json:"end_date,omitempty" description:"end date of the query"`
	Currency       string `url:"currency" json:"currency,omitempty" description:"currency alias"`
	Coin           string `url:"coin" json:"coin,omitempty" description:"currency alias"`
	WalletFundType string `url:"wallet_fund_type" json:"wallet_fund_type,omitempty" description:"wallet fund type"`
	Limit          int64  `url:"limit" json:"limit,omitempty" description:"number of records to return"`
	Page           int64  `url:"page" json:"page,omitempty" description:"page number"`
}

type WalletFundRecordsResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              WalletFundRecordsResult `json:"result"`
}

type WalletFundRecordsResult struct {
	Data []WalletFundRecordsData `json:"data"`
}

type WalletFundRecordsData struct {
	UserID        int64   `url:"user_id" json:"user_id"`
	ID            int64   `url:"id" json:"id"`
	Coin          string  `url:"coin" json:"coin"`
	WalletID      int64   `url:"wallet_id" json:"wallet_id"`
	Type          string  `url:"type" json:"type"`
	Amount        float64 `url:"amount" json:"amount"`
	TxID          string  `url:"tx_id" json:"tx_id"`
	Address       string  `url:"address" json:"address"`
	WalletBalance string  `url:"wallet_balance" json:"wallet_balance"`
	ExecTime      string  `url:"exec_time" json:"exec_time"`
	CrossSeq      int64   `url:"cross_seq" json:"cross_seq"`
}

type WithdrawRecordsParams struct {
	StartDate string `url:"start_date" json:"start_date,omitempty" description:"start date of the query"`
	EndDate   string `url:"end_date" json:"end_date,omitempty" description:"end date of the query"`
	Coin      string `url:"coin" json:"coin,omitempty" description:"currency alias"`
	Status    string `url:"status" json:"status,omitempty" description:"status of the withdrawal"`
	Limit     int64  `url:"limit" json:"limit,omitempty" description:"number of records to return"`
	Page      int64  `url:"page" json:"page,omitempty" description:"page number"`
}

type WithdrawRecordsResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              WithdrawRecordsResult `json:"result"`
}

type WithdrawRecordsResult struct {
	Data        []WithdrawRecordsData `json:"data"`
	CurrentPage int64                 `json:"current_page"`
	LastPage    int64                 `json:"last_page"`
}

type WithdrawRecordsData struct {
	ID          int64  `url:"id" json:"id"`
	UserID      int64  `url:"user_id" json:"user_id"`
	Coin        string `url:"coin" json:"coin"`
	Status      string `url:"status" json:"status"`
	Amount      string `url:"amount" json:"amount"`
	Fee         string `url:"fee" json:"fee"`
	Address     string `url:"address" json:"address"`
	TxID        string `url:"tx_id" json:"tx_id"`
	SubmittedAt string `url:"submitted_at" json:"submitted_at"`
	UpdatedAt   string `url:"updated_at" json:"updated_at"`
}

type AssetExchangeRecordsParams struct {
	Limit     int64  `url:"limit" json:"limit,omitempty" description:"number of records to return"`
	From      int64  `url:"from" json:"from,omitempty" description:"from id"`
	Direction string `url:"direction" json:"direction,omitempty" description:"direction of the query"`
}

type AssetExchangeRecordsResponse struct {
	domain.BaseResponse `json:",inline"`
	Result              []AssetExchangeRecordsResult `json:"result"`
}

type AssetExchangeRecordsResult struct {
	FromCoin     string  `url:"from_coin" json:"from_coin"`
	FromAmount   float64 `url:"from_amount" json:"from_amount"`
	ToCoin       string  `url:"to_coin" json:"to_coin"`
	ToAmount     float64 `url:"to_amount" json:"to_amount"`
	ExchangeRate float64 `url:"exchange_rate" json:"exchange_rate"`
	FromFee      float64 `url:"from_fee" json:"from_fee"`
	CreatedAt    string  `url:"created_at" json:"created_at"`
}
