package domain

type BaseResponse struct {
	ReturnCode    int    `json:"retCode"`
	ReturnMsg     string `json:"retMsg"`
	ReturnExtInfo string `json:"extInfo"`
	Time          int64  `json:"time_now"`
}
