package dto

type ResponseWrapper struct {
	Status uint32      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}
