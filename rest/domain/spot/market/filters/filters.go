package filters

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type OrderBookDepthFilter struct {
	Symbol string `json:"symbol,required"`
	Limit  string `json:"limit,omitempty"`
}

// ToQuery converts the current filters into querystring
func (o *OrderBookDepthFilter) ToQuery(u *url.URL) {
	// return as params here
	mydata := map[string]string{}
	data, _ := json.Marshal(u.Query())
	err := json.Unmarshal(data, &mydata)
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", mydata)
	queryString := u.Query()
	if o.Symbol != "" {
		queryString.Set("symbol", o.Symbol)
	}

	if o.Limit != "" {
		queryString.Set("limit", o.Limit)
	}
	u.RawQuery = queryString.Encode()
}

type MergedOrderBookFilter struct {
	Symbol string `json:"symbol,required"`
	Scale  string `json:"scale,omitempty"`
	Limit  string `json:"limit,omitempty"`
}

// ToQuery converts the current filters into querystring
func (m *MergedOrderBookFilter) ToQuery(u *url.URL) {
	queryString := u.Query()
	if m.Symbol != "" {
		queryString.Set("symbol", m.Symbol)
	}

	if m.Scale != "" {
		queryString.Set("scale", m.Scale)
	}

	if m.Limit != "" {
		queryString.Set("limit", m.Limit)
	}
	u.RawQuery = queryString.Encode()
}

type TradeRecordsFilter struct {
	Symbol string `json:"symbol,required"`
	Limit  string `json:"limit,omitempty"`
}

// ToQuery converts the current filters into querystring
func (t *TradeRecordsFilter) ToQuery(u *url.URL) {
	queryString := u.Query()
	if t.Symbol != "" {
		queryString.Set("symbol", t.Symbol)
	}

	if t.Limit != "" {
		queryString.Set("limit", t.Limit)
	}
	u.RawQuery = queryString.Encode()
}

type KlineFilter struct {
	Symbol    string `json:"symbol,required"`
	Interval  string `json:"interval,required"`
	Limit     string `json:"limit,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
}

// ToQuery converts the current filters into querystring
func (k *KlineFilter) ToQuery(u *url.URL) {
	queryString := u.Query()
	if k.Symbol != "" {
		queryString.Set("symbol", k.Symbol)
	}
	if k.Interval != "" {
		queryString.Set("interval", k.Interval)
	}
	if k.Limit != "" {
		queryString.Set("limit", k.Limit)
	}
	if k.StartTime != "" {
		queryString.Set("start_time", k.StartTime)
	}
	if k.EndTime != "" {
		queryString.Set("end_time", k.EndTime)
	}
	u.RawQuery = queryString.Encode()
}
