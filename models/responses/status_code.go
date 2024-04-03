package responses

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// 定義系統錯誤與相關回傳訊息
const (
	Success       = "0"
	Error         = "1"
)

var MsgText = map[string]string{
	Success:       "Success",
	Error:         "Has some problem",
}

func Status(code string, data interface{}) Response {
	return Response{
		Status:  code,
		Data:    data,
		Message: MsgText[code],
	}
}
