package helper

import "net/http"

const (
	RCSuccess      string = "00"
	RCGeneralError string = "99"
)

const (
	DescriptionSuccess string = "success"
	DescriptionFailed  string = "failed"
	DescErrorGeneral   string = "general error"
)

var (
	ErrBadRequest  = ErrorStruct{Code: "BAD01", HTTPCode: http.StatusBadRequest, Description: DescriptionFailed, Message: "bad request"}
	ErrNotFound    = ErrorStruct{Code: "DT01", HTTPCode: http.StatusNotFound, Description: DescriptionFailed, Message: "data not found"}
	ErrLoginFailed = ErrorStruct{Code: "AUTHXX", HTTPCode: http.StatusUnauthorized, Description: DescriptionFailed, Message: "gagal login"}
)
