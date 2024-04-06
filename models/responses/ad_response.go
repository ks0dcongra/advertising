package responses

import "time"

type AdsInfo struct {
	AID   int       `gorm:"column:aid;primaryKey;autoIncrement;"`
	Title string    `gorm:"column:title;type:varchar(128);default:''"`
	EndAt time.Time `gorm:"column:end_at;default:null"`
}
