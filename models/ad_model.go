package models

import (
	"time"

	"github.com/lib/pq"
)

type Ads struct {
	AID        int            `gorm:"column:aid;primaryKey;autoIncrement;"`
	Title      string         `gorm:"column:title;type:varchar(128);default:''"`
	StartAt    time.Time      `gorm:"column:start_at;default:null"`
	EndAt      time.Time      `gorm:"column:end_at;default:null"`
	AgeStart   int            `gorm:"column:age_start;type:integer;default:0"`
	AgeEnd     int            `gorm:"column:age_end;type:integer;default:0"`
	Gender     pq.StringArray `gorm:"column:gender;type:varchar(16)[];default:'{}'"`
	Country    pq.StringArray `gorm:"column:country;type:varchar(256)[];default:'{}'"`
	Platform   pq.StringArray `gorm:"column:platform;type:varchar(256)[];default:'{}'"`
	Updated_At time.Time      `gorm:"column:updated_at;autoUpdateTime;default:current_timestamp"`
}

func (Ads) TableName() string {
	return "advertising.ads"
}
