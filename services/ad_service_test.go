package services_test

import (
	"advertising/define"
	"advertising/models"
	"advertising/models/requests"
	"advertising/services"
	"errors"
	"time"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAdRepository struct {
	mock.Mock
}

// a, _ := args.Get(0).(models.Ads)
func (m *MockAdRepository) GetAds(items []string, requestData *requests.ConditionInfoOfPage, query string, args ...interface{}) ([]models.Ads, error) {
	// 这里可以添加 GetAds 方法的模拟行为
	return nil, nil
}

func (m *MockAdRepository) CreateAd(adsModel *models.Ads) error {
	return nil
}

func (m *MockAdRepository) GetActiveAds(nowDateTime time.Time) (int64, error) {
	args := m.Called(nowDateTime)
	count, _ := args.Get(0).(int64)
	if count == 1001 {
		return count, nil
	}
	if args.Error(1) != nil {
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
		name        string
		args        args
		want        string
		mockDbErr   error
		mockDbCount int64
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
			mockDbErr:   nil,
			mockDbCount: 0,
			want:        define.Success,
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
			mockDbErr:   errors.New("DB error"),
			mockDbCount: 0,
			want:        define.DbErr,
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
			mockDbErr:   nil,
			mockDbCount: int64(1001),
			want:        define.AdAmountExceeded,
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
			mockDbErr:   nil,
			mockDbCount: 0,
			want:        define.TimeParsedError,
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
			mockDbErr:   nil,
			mockDbCount: 0,
			want:        define.RegexParsedError,
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
			mockDbErr:   nil,
			mockDbCount: 0,
			want:        define.RegexParsedError,
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
			mockDbErr:   nil,
			mockDbCount: 0,
			want:        define.RegexParsedError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adRepoMock := new(MockAdRepository)
			adRepoMock.On("CreateAd", mock.Anything).Return(tt.mockDbErr)
			adRepoMock.On("GetActiveAds", mock.Anything).Return(tt.mockDbCount, tt.mockDbErr)
			adRepoMock.On("GetTodayCreateAds", mock.Anything).Return(tt.mockDbCount, tt.mockDbErr)

			// Inject the mock AdRepository into AdService
			adService := services.AdService{
				AdRepository: adRepoMock,
			}

			// got := adService.CreateAd(tt.args.requestData)
			got := adService.CreateAd(tt.args.requestData)
			
			assert.Equal(t, tt.want, got)
		})
	}
}

// func TestAdService_GetAd(t *testing.T) {
// 	type args struct {
// 		requestData *requests.ConditionInfoOfPage
// 	}
// 	tests := []struct {
// 		name  string
// 		a     *AdService
// 		args  args
// 		want  []responses.AdsInfo
// 		want1 string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, got1 := tt.a.GetAd(tt.args.requestData)
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("AdService.GetAd() got = %v, want %v", got, tt.want)
// 			}
// 			if got1 != tt.want1 {
// 				t.Errorf("AdService.GetAd() got1 = %v, want %v", got1, tt.want1)
// 			}
// 		})
// 	}
// }
