package helper

type Response struct {
	ResponseCode string `json:"rc"`
	Description  string `json:"desc"`
	Message      string `json:"message"`
	Data         interface{}
}

type ErrorStruct struct {
	Code        string `json:"rc"`
	HTTPCode    int    `json:"-"`
	Description string `json:"desc"`
	Message     string `json:"message"`
}
