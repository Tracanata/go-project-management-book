package helper

type Response struct {
	ResponseCode string `json:"rc"`
	Description  string `json:"desc"`
	Message      string `json:"msg"`
	Data         interface{}
}
