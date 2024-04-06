package responses

import "advertising/define"

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Status(code string, data interface{}) Response {
	return Response{
		Status:  code,
		Message: define.MsgText[code],
		Data:    data,
	}
}
