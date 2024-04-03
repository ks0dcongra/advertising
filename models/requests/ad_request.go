package requests

type CreateAd struct {
	Title     string         `json:"title" binding:"required,lte=60"`
	StartAt   string         `json:"start_at" binding:"required,lte=60"`
	EndAt     string         `json:"end_at" binding:"required,lte=60"`
	Condition *ConditionInfo `json:"condition"`
}

type ConditionInfo struct {
	AgeStart string `json:"age_start" binding:"required,lte=60"`
	AgeEnd   string `json:"age_end" binding:"required,lte=60"`
	Gender   string `json:"gender" binding:"required,lte=60"`
	Country  string `json:"country" binding:"required,lte=60"`
	Platform string `json:"platform" binding:"required,lte=60"`
}
