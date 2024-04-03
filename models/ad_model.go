package models

import (
	"time"
)

type Ads struct {
	A_ID       int       `gorm:"column:title;primaryKey;autoIncrement;"`
	Title      string    `gorm:"column:title;type:varchar(50);default:''"`
	StartAt    string    `gorm:"column:start_at;type:varchar(50);default:''"`
	EndAt      string    `gorm:"column:end_at;type:varchar(50);default:''"`
	AgeStart   int       `gorm:"column:age_start;type:numeric(1);default:0"`
	AgeEnd     int       `gorm:"column:age_end;type:numeric(1);default:0"`
	Gender     string    `gorm:"column:gender;type:numeric(1);default:''	"`
	Country    string    `gorm:"column:country;type:varchar(60);default:''"`
	Platform   string    `gorm:"column:platform;type:varchar(60);default:''"`
	Updated_At time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Ads) TableName() string {
	return "advertising.ads"
}