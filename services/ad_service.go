package services

import (
	"advertising/configs"
	"advertising/define"
	"advertising/models"
	"advertising/models/requests"
	"advertising/models/responses"
	"advertising/repositories"
	"advertising/repositories/redis"
	"encoding/json"
	"fmt"

	"regexp"
	"strings"
	"time"
)

type AdServiceInterface interface {
	CreateAd(requestData *requests.CreateAd) string
	GetAd(requestData *requests.ConditionInfoOfPage) ([]responses.AdsInfo, string)
}

type AdService struct {
	AdRepository repositories.AdRepositoryInterface
}

func NewAdService() AdServiceInterface {
	db := configs.DbConn
	adRepo := repositories.NewAdRepository(db)
	return &AdService{adRepo}
}

func (a *AdService) CreateAd(requestData *requests.CreateAd) string {
	nowDateTime := time.Now().UTC()
	nowDateTimeStr := nowDateTime.Format(define.YMD)
	nowDate, err := time.Parse(define.YMD, nowDateTimeStr)
	if err != nil {
		return define.TimeParsedError
	}
	nextDate := nowDate.Add(24 * time.Hour)
	// Dependency injection db connection
	db := configs.DbConn
	count, err := repositories.NewAdRepository(db).GetActiveAds(nowDateTime)
	if err != nil {
		return define.DbErr
	}
	if count >= 1000 {
		return define.AdAmountExceeded
	}

	count, err = repositories.NewAdRepository(db).GetTodayCreateAds(nowDate, nextDate)
	if err != nil {
		return define.DbErr
	}

	if count >= 3000 {
		return define.AdAmountExceeded
	}

	// Check request's time
	parsedStartAt, err := time.Parse(time.RFC3339, requestData.StartAt)
	if err != nil {
		return define.TimeParsedError
	}

	parsedEndAt, err := time.Parse(time.RFC3339, requestData.EndAt)
	if err != nil {
		return define.TimeParsedError
	}

	// Build SQL insert models
	adsModel := new(models.Ads)

	if requestData.Conditions[0].AgeStart != 0 {
		adsModel.AgeStart = requestData.Conditions[0].AgeStart
	}

	if requestData.Conditions[0].AgeEnd != 0 {
		adsModel.AgeEnd = requestData.Conditions[0].AgeEnd
	}

	// Use strings.Join to convert slices to strings, and separated by commas
	if len(requestData.Conditions[0].Gender) != 0 {
		genderStr := strings.Join(requestData.Conditions[0].Gender, ",")
		// Valid gender whether the string consists of one letters
		if match, err := regexp.MatchString(`[A-Z]{1},?\s?`, genderStr); !match || err != nil {
			return define.RegexParsedError
		}
		adsModel.Gender = requestData.Conditions[0].Gender
	}

	if len(requestData.Conditions[0].Country) != 0 {
		countryStr := strings.Join(requestData.Conditions[0].Country, ",")
		// Valid country whether the string consists of two letters connected

		if match, err := regexp.MatchString(`[A-Z]{2},?\s?`, countryStr); !match || err != nil {
			return define.RegexParsedError
		}

		adsModel.Country = requestData.Conditions[0].Country
	}

	if len(requestData.Conditions[0].Platform) != 0 {
		platformStr := strings.Join(requestData.Conditions[0].Platform, ",")
		// Valid platform whether the string consists of letters
		if match, err := regexp.MatchString(`[a-zA-Z]+,?\s?`, platformStr); !match || err != nil {
			return define.RegexParsedError
		}
		adsModel.Platform = requestData.Conditions[0].Platform
	}

	adsModel.Title = requestData.Title
	adsModel.StartAt = parsedStartAt
	adsModel.EndAt = parsedEndAt

	// Dependency injection db connection
	err = repositories.NewAdRepository(db).CreateAd(adsModel)
	if err != nil {
		return define.DbErr
	}

	return define.Success
}

func (a *AdService) GetAd(requestData *requests.ConditionInfoOfPage) ([]responses.AdsInfo, string) {
	nowDateTime := time.Now().UTC()
	query := "start_at <= ? AND end_at >= ?"
	args := []interface{}{nowDateTime, nowDateTime}

	if requestData.Age != 0 {
		query += " AND age_start <= ? AND age_end >= ?"
		args = append(args, requestData.Age, requestData.Age)
	}

	if requestData.Country != "" {
		query += " AND ? = ANY(country)"
		args = append(args, requestData.Country)
	}

	if requestData.Gender != "" {
		query += " AND ? = ANY(gender)"
		args = append(args, requestData.Gender)
	}

	if requestData.Platform != "" {
		query += " AND ? = ANY(platform)"
		args = append(args, requestData.Platform)
	}

	// SQL limit default is 5
	if requestData.AdLimit < 1 || requestData.AdLimit > 100 {
		requestData.AdLimit = 5
	}

	redisKey := fmt.Sprintf("{Age:%d,Gender:%s,Country:%s,Platform:%s,AdOffset:%d,AdLimit:%d}", requestData.Age, requestData.Gender, requestData.Country, requestData.Platform, requestData.AdOffset, requestData.AdLimit)

	var responseData []responses.AdsInfo
	// 如果抓取redis的過程有error,就去用postgres撈取資料並設置redis
	redisGetResult, err := redis.NewRedisRepositoryImpl().Get(define.AdsConditionPrefix + redisKey)
	if err != nil {
		// Dependency injection db connection
		db := configs.DbConn
		items := []string{"aid", "title", "end_at"}
		ads, err := repositories.NewAdRepository(db).GetAds(items, requestData, query, args...)
		if err != nil {
			return nil, define.DbErr
		}

		/* 
		Use marshal and unmarshal to choose partial column in return db result
		*/
		adsBytes, err := json.Marshal(ads)
		if err != nil {
			return nil, define.JsonMarshalError
		}
		err = json.Unmarshal(adsBytes, &responseData)
		if err != nil {
			return nil, define.JsonUnmarshalError
		}

		// Store db result of adsBytes in redis
		err = redis.NewRedisRepositoryImpl().Set(define.AdsConditionPrefix+redisKey, adsBytes, time.Second*2)
		if err != nil {
			return nil, define.RedisErr
		}

		return responseData, define.Success
	}

	err = json.Unmarshal(redisGetResult, &responseData)
	if err != nil {
		return nil, define.JsonUnmarshalError
	}
	return responseData, define.RedisSuccess
}
