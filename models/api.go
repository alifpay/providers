package models

import (
	"time"

	"github.com/shopspring/decimal"
)

//ReqAPI - request of standart protocol for providers
type ReqAPI struct {
	ID      uint64                 `json:"id"`
	Action  string                 `json:"action"`
	SrvID   string                 `json:"srv_id"`
	Account string                 `json:"account"`
	Amount  decimal.Decimal        `json:"amount"`
	Info    map[string]interface{} `json:"info"`
	RegTime time.Time              `json:"time"`
}

//RespAPI - response of standart protocol for providers
type RespAPI struct {
	ID            uint64 `json:"id"`
	Code          uint64 `json:"code"`
	ResponseID    string `json:"response_id"`
	InfoForClient string `json:"info_for_client"`
}
