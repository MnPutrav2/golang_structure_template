package model

type ResponseMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseBody struct {
	Status   string `json:"status"`
	Response any    `json:"message"`
}
