package domain

type BaseResponse struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode int    `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
}
