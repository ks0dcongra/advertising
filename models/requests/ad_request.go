package requests

type CreateAd struct {
	Title      string          `json:"title" binding:"required,lte=128"`
	StartAt    string          `json:"startAt" binding:"required,lte=32"`
	EndAt      string          `json:"endAt" binding:"required,lte=32"`
	Conditions []ConditionInfo `json:"conditions"`
}

type ConditionInfo struct {
	AgeStart int      `json:"ageStart" binding:"required,gte=1,lte=100"`
	AgeEnd   int      `json:"ageEnd" binding:"required,gte=1,lte=100"`
	Gender   []string `json:"gender" binding:"required,lte=16"`
	Country  []string `json:"country" binding:"required,lte=256"`
	Platform []string `json:"platform" binding:"required,lte=256"`
}

type ConditionInfoOfPage struct {
	Age      int    `form:"age" binding:"gte=0,lte=100"`
	Gender   string `form:"gender" binding:"lte=16"`
	Country  string `form:"country" binding:"lte=256"`
	Platform string `form:"platform" binding:"lte=256"`
	AdOffset int    `form:"offset" binding:"gte=0"`
	AdLimit  int    `form:"limit" binding:"gte=0,lte=100"`
}
