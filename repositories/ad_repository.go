package repositories

import (
	"advertising/models"
	"advertising/models/requests"
	"errors"
	"time"

	"gorm.io/gorm"
)

type AdRepositoryInterface interface {
	CreateAd(adsModel *models.Ads) error
	GetAds(items []string, requestData *requests.ConditionInfoOfPage, query string, args ...interface{}) ([]models.Ads, error)
	GetActiveAds(nowDateTime time.Time) (int64, error)
	GetTodayCreateAds(nowDate time.Time, nextDate time.Time) (int64, error)
}

type AdRepository struct {
	db *gorm.DB
}

func NewAdRepository(db *gorm.DB) AdRepositoryInterface {
	return &AdRepository{db}
}

func (a *AdRepository) CreateAd(adsModel *models.Ads) error {
	result := a.db.Create(&adsModel)
	return result.Error
}

func (a *AdRepository) GetAds(items []string, requestData *requests.ConditionInfoOfPage, query string, args ...interface{}) ([]models.Ads, error) {
	ads := make([]models.Ads, 0)

	results := a.db.Select(items).Where(query, args...).Order("end_at ASC").Limit(requestData.AdLimit).Offset(requestData.AdOffset).Debug().Find(&ads)
	if results.Error != nil {
		return nil, results.Error
	}
	if results.RowsAffected == 0 {
		return nil, errors.New("result not found")
	}

	return ads, nil
}

func (a *AdRepository) GetActiveAds(nowDateTime time.Time) (int64, error) {
	var count int64
	result := a.db.Model(&models.Ads{}).Where("start_at <= ? AND end_at >= ?", nowDateTime,nowDateTime).Count(&count)
    if result.Error != nil {
        return count, result.Error
    }

	return count, nil
}

func (a *AdRepository) GetTodayCreateAds(nowDate time.Time, nextDate time.Time) (int64, error) {
	var count int64
	result := a.db.Model(&models.Ads{}).Where("updated_at >= ? AND updated_at < ?", nowDate, nextDate).Count(&count)
    if result.Error != nil {
        return count, result.Error
    }

	return count, nil
}

