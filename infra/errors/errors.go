package errors

import (
	"strconv"
)

type ErrorCode struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

func (e *ErrorCode) Error() string {
	return "Code: " + strconv.Itoa(e.Code) + ", Detail " + e.Detail
}

var (
	ErrorRequest        = &ErrorCode{400, "http request fail"}
	ErrorResponse       = &ErrorCode{400, "http response fail"}
	ErrorStatusCode     = &ErrorCode{400, "http status code fail"}
	ErrorMaoYanRankList = &ErrorCode{400, "get maoyan data fail"}
)
