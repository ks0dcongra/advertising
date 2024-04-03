package responses

type TxnData_Statistics struct {
	Items []ItemsInfo `json:"items"`
}

type ItemsInfo struct {
	Title   string `json:"title"`
	EndAt   string `json:"end_at"`
}
