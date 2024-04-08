package services_test

import (
	"advertising/define"
	"advertising/helpers"
	"advertising/models"
	"advertising/models/requests"
	"advertising/models/responses"
	"advertising/services"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAdRepository struct {
	mock.Mock
}

type MockRedisRepository struct {
	mock.Mock
}

func (m *MockRedisRepository) Set(keyName string, value []byte, expiration time.Duration) error {
	mockArgs := m.Called()
	return mockArgs.Error(0)
}

func (m *MockRedisRepository) Get(keyName string) ([]byte, error) {
	mockArgs := m.Called()
	mockRedisResult, _ := mockArgs.Get(0).([]byte)
	if mockRedisResult != nil {
		return mockRedisResult, nil
	}
	if mockArgs.Error(1) != nil {
		return nil, errors.New("Redis error")
	}
	return nil, nil
}

// a, _ := args.Get(0).(models.Ads)
func (m *MockAdRepository) GetAds(items []string, requestData *requests.ConditionInfoOfPage, query string, args ...interface{}) ([]models.Ads, error) {
	mockArgs := m.Called()
	mockDbResult, _ := mockArgs.Get(0).([]models.Ads)
	if mockDbResult != nil {
		return mockDbResult, nil
	}
	if mockArgs.Error(1) != nil {
		return nil, errors.New("DB error")
	}
	return nil, nil
}

func (m *MockAdRepository) CreateAd(adsModel *models.Ads) error {
	return nil
}

func (m *MockAdRepository) GetActiveAds(nowDateTime time.Time) (int64, error) {
	mockArgs := m.Called()
	count, _ := mockArgs.Get(0).(int64)
	if count == 1001 {
		return count, nil
	}
	if mockArgs.Error(1) != nil {
		return 0, errors.New("DB error")
	}
	return 0, nil
}

func (m *MockAdRepository) GetTodayCreateAds(nowDate, nextDate time.Time) (int64, error) {
	// 这里可以添加 GetTodayCreateAds 方法的模拟行为
	return 0, nil
}

func TestAdService_CreateAd(t *testing.T) {
	type args struct {
		requestData *requests.CreateAd
	}
	tests := []struct {
		name              string
		args              args
		want              string
		mockDbErr         error
		mockDbResultCount int64
	}{
		{
			name: "Success",
			args: args{
				requestData: &requests.CreateAd{
					Title:   "AD 55",
					StartAt: "2023-12-10T03:00:00.000Z",
					EndAt:   "2024-12-31T16:00:00.000Z",
					Conditions: []requests.ConditionInfo{
						{
							AgeStart: 20,
							AgeEnd:   30,
							Gender:   []string{"F"},
							Country:  []string{"TW", "JP"},
							Platform: []string{"android", "ios"},
						},
					},
				},
			},
			mockDbErr:         nil,
			mockDbResultCount: 0,
			want:              define.Success,
		},
		{
			name: "GetActiveAds has DbErr",
			args: args{
				requestData: &requests.CreateAd{
					Title:   "AD 56",
					StartAt: "2023-12-10T03:00:00.000Z",
					EndAt:   "2024-12-31T16:00:00.000Z",
					Conditions: []requests.ConditionInfo{
						{
							AgeStart: 20,
							AgeEnd:   30,
							Gender:   []string{"F"},
							Country:  []string{"TW", "JP"},
							Platform: []string{"android", "ios"},
						},
					},
				},
			},
			mockDbErr:         errors.New("DB error"),
			mockDbResultCount: 0,
			want:              define.DbErr,
		},
		{
			name: "GetActiveAds function's return count that greated than 1000",
			args: args{
				requestData: &requests.CreateAd{
					Title:   "AD 56",
					StartAt: "2023-12-10T03:00:00.000Z",
					EndAt:   "2024-12-31T16:00:00.000Z",
					Conditions: []requests.ConditionInfo{
						{
							AgeStart: 20,
							AgeEnd:   30,
							Gender:   []string{"F"},
							Country:  []string{"TW", "JP"},
							Platform: []string{"android", "ios"},
						},
					},
				},
			},
			mockDbErr:         nil,
			mockDbResultCount: int64(1001),
			want:              define.AdAmountExceeded,
		},
		{
			name: "requestData's column 'StartAt' Wrong",
			args: args{
				requestData: &requests.CreateAd{
					Title:   "AD 57",
					StartAt: "2023-12-10T03:00:00.000",
					EndAt:   "2024-12-31T16:00:00.000Z",
					Conditions: []requests.ConditionInfo{
						{
							AgeStart: 20,
							AgeEnd:   30,
							Gender:   []string{"F"},
							Country:  []string{"TW", "JP"},
							Platform: []string{"android", "ios"},
						},
					},
				},
			},
			mockDbErr:         nil,
			mockDbResultCount: 0,
			want:              define.TimeParsedError,
		},
		{
			name: "requestData's column 'Gender' Wrong",
			args: args{
				requestData: &requests.CreateAd{
					Title:   "AD 58",
					StartAt: "2023-12-10T03:00:00.000Z",
					EndAt:   "2024-12-31T16:00:00.000Z",
					Conditions: []requests.ConditionInfo{
						{
							AgeStart: 20,
							AgeEnd:   30,
							Gender:   []string{"FM", "F"},
							Country:  []string{"TW", "JP"},
							Platform: []string{"android", "ios"},
						},
					},
				},
			},
			mockDbErr:         nil,
			mockDbResultCount: 0,
			want:              define.RegexParsedError,
		},
		{
			name: "requestData's column 'Country' Wrong",
			args: args{
				requestData: &requests.CreateAd{
					Title:   "AD 59",
					StartAt: "2023-12-10T03:00:00.000Z",
					EndAt:   "2024-12-31T16:00:00.000Z",
					Conditions: []requests.ConditionInfo{
						{
							AgeStart: 20,
							AgeEnd:   30,
							Gender:   []string{"F"},
							Country:  []string{"TWD", "JP"},
							Platform: []string{"android", "ios"},
						},
					},
				},
			},
			mockDbErr:         nil,
			mockDbResultCount: 0,
			want:              define.RegexParsedError,
		},
		{
			name: "requestData's column 'Platform' Wrong",
			args: args{
				requestData: &requests.CreateAd{
					Title:   "AD 60",
					StartAt: "2023-12-10T03:00:00.000Z",
					EndAt:   "2024-12-31T16:00:00.000Z",
					Conditions: []requests.ConditionInfo{
						{
							AgeStart: 20,
							AgeEnd:   30,
							Gender:   []string{"F"},
							Country:  []string{"TW", "JP"},
							Platform: []string{"123"},
						},
					},
				},
			},
			mockDbErr:         nil,
			mockDbResultCount: 0,
			want:              define.RegexParsedError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adRepoMock := new(MockAdRepository)
			adRepoMock.On("CreateAd", mock.Anything).Return(tt.mockDbErr)
			adRepoMock.On("GetActiveAds", mock.Anything).Return(tt.mockDbResultCount, tt.mockDbErr)
			adRepoMock.On("GetTodayCreateAds", mock.Anything).Return(tt.mockDbResultCount, tt.mockDbErr)

			// Inject the mock AdRepository into AdService
			adService := services.AdService{
				AdRepository: adRepoMock,
			}

			got := adService.CreateAd(tt.args.requestData)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAdService_GetAd(t *testing.T) {
	helpers.SetNow(time.Date(2024, time.March, 27, 17, 06, 29, 0, time.FixedZone("CST", 8*3600)))

	type args struct {
		requestData *requests.ConditionInfoOfPage
	}
	tests := []struct {
		name            string
		args            args
		mockDbResult    []models.Ads
		mockRedisResult []byte
		mockDbErr       error
		mockRedisGetErr error
		mockRedisSetErr error
		want            []responses.AdsInfo
		want1           string
	}{
		{
			name: "Success from Redis",
			args: args{
				requestData: &requests.ConditionInfoOfPage{
					Age:      30,
					Gender:   "F",
					Country:  "TW",
					Platform: "ios",
					AdOffset: 0,
					AdLimit:  0,
				},
			},
			mockDbResult: nil,
			mockDbErr:       nil,
			mockRedisResult: helpers.ToByteSlice(&[]responses.AdsInfo{
				{
					AID:   1,
					Title: "AD 55",
					EndAt: helpers.Now().UTC(),
				},
			}),
			mockRedisGetErr: nil,
			mockRedisSetErr: nil,
			want: []responses.AdsInfo{
				{
					AID:   1,
					Title: "AD 55",
					EndAt: helpers.Now().UTC(),
				},
			},
			want1: define.RedisSuccess,
		},
		{
			name: "Success from DB",
			args: args{
				requestData: &requests.ConditionInfoOfPage{
					Age:      30,
					Gender:   "F",
					Country:  "TW",
					Platform: "ios",
					AdOffset: 0,
					AdLimit:  0,
				},
			},
			mockDbResult: []models.Ads{
				{
					AID:   1,
					Title: "AD 55",
					EndAt: helpers.Now().UTC(),
				},
			},
			mockDbErr:       nil,
			mockRedisResult: nil,
			mockRedisGetErr: errors.New("Redis get error"),
			mockRedisSetErr: nil,
			want: []responses.AdsInfo{
				{
					AID:   1,
					Title: "AD 55",
					EndAt: helpers.Now().UTC(),
				},
			},
			want1: define.Success,
		},
		{
			name: "DB Error",
			args: args{
				requestData: &requests.ConditionInfoOfPage{
					Age:      30,
					Gender:   "F",
					Country:  "TW",
					Platform: "ios",
					AdOffset: 0,
					AdLimit:  0,
				},
			},
			mockDbResult: nil,
			mockDbErr:       errors.New("DB get error"),
			mockRedisResult: nil,
			mockRedisGetErr: errors.New("Redis get error"),
			mockRedisSetErr: nil,
			want: nil,
			want1: define.DbErr,
		},
		{
			name: "Set Redis Failed",
			args: args{
				requestData: &requests.ConditionInfoOfPage{
					Age:      30,
					Gender:   "F",
					Country:  "TW",
					Platform: "ios",
					AdOffset: 0,
					AdLimit:  0,
				},
			},
			mockDbResult: []models.Ads{
				{
					AID:   1,
					Title: "AD 55",
					EndAt: helpers.Now().UTC(),
				},
			},
			mockDbErr:       nil,
			mockRedisResult: nil,
			mockRedisGetErr: errors.New("Redis get error"),
			mockRedisSetErr: errors.New("Redis set error"),
			want: nil,
			want1: define.RedisErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adRepoMock := new(MockAdRepository)
			adRepoMock.On("GetAds", mock.Anything).Return(tt.mockDbResult, tt.mockDbErr)

			redisRepoMock := new(MockRedisRepository)
			redisRepoMock.On("Get", mock.Anything).Return(tt.mockRedisResult, tt.mockRedisGetErr)
			redisRepoMock.On("Set", mock.Anything).Return(tt.mockRedisSetErr)

			// Inject the mock AdRepository into AdService
			adService := services.AdService{
				AdRepository:    adRepoMock,
				RedisRepository: redisRepoMock,
			}
			got, got1 := adService.GetAds(tt.args.requestData)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
