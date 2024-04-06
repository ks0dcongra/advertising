package requests

type CreateAd struct {
	Title      string          `json:"title" binding:"required,lte=128"`
	StartAt    string          `json:"startAt" binding:"required,lte=32"`
	EndAt      string          `json:"endAt" binding:"required,lte=32"`
	Conditions []ConditionInfo `json:"conditions"`
}

type ConditionInfo struct {
	AgeStart int      `json:"ageStart" binding:"lte=3"`
	AgeEnd   int      `json:"ageEnd" binding:"lte=3"`
	Gender   []string `json:"gender" binding:"lte=16"`
	Country  []string `json:"country" binding:lte=256"`
	Platform []string `json:"platform" binding:"lte=256"`
}

type ConditionInfoOfPage struct {
	Age      int    `json:"age" binding:"lte=3"`
	Gender   string `json:"gender" binding:"lte=16"`
	Country  string `json:"country" binding:lte=256"`
	Platform string `json:"platform" binding:"lte=256"`
	AdOffset int    `json:"offset" binding:"required,lte=8"`
	AdLimit  int    `json:"limit" binding:"required,lte=8"`
}
