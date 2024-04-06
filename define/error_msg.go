package define

var MsgText = map[string]string{
	Success:            "Success",
	AppError:           "Application has some problem",
	ParameterErr:       "Parameter error, please check your field",
	DbErr:              "SQL not found",
	TimeParsedError:    "Time parsed from RFC3339 failed",
	RegexParsedError:   "Regex detected request has some problem",
	JsonMarshalError:   "JSON marshal failed",
	JsonUnmarshalError: "JSON unmarshal failed",
	AdLimitExceeded:    "Amount of ads already reached maximum",
}
