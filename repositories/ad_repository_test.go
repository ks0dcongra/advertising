package repositories_test

// import (
// 	"advertising/models"
// 	"advertising/models/requests"
// 	"reflect"
// 	"testing"
// 	"time"
	
// )

// func TestAdRepository_CreateAd(t *testing.T) {
// 	type args struct {
// 		adsModel *models.Ads
// 	}
// 	tests := []struct {
// 		name    string
// 		a       *AdRepository
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.a.CreateAd(tt.args.adsModel); (err != nil) != tt.wantErr {
// 				t.Errorf("AdRepository.CreateAd() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestAdRepository_GetAds(t *testing.T) {
// 	type args struct {
// 		items       []string
// 		requestData *requests.ConditionInfoOfPage
// 		query       string
// 		args        []interface{}
// 	}
// 	tests := []struct {
// 		name    string
// 		a       *AdRepository
// 		args    args
// 		want    []models.Ads
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.a.GetAds(tt.args.items, tt.args.requestData, tt.args.query, tt.args.args...)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("AdRepository.GetAds() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("AdRepository.GetAds() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestAdRepository_GetActiveAds(t *testing.T) {
// 	type args struct {
// 		nowDateTime time.Time
// 	}
// 	tests := []struct {
// 		name    string
// 		a       *AdRepository
// 		args    args
// 		want    int64
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.a.GetActiveAds(tt.args.nowDateTime)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("AdRepository.GetActiveAds() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("AdRepository.GetActiveAds() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestAdRepository_GetTodayCreateAds(t *testing.T) {
// 	type args struct {
// 		nowDate  time.Time
// 		nextDate time.Time
// 	}
// 	tests := []struct {
// 		name    string
// 		a       *AdRepository
// 		args    args
// 		want    int64
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.a.GetTodayCreateAds(tt.args.nowDate, tt.args.nextDate)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("AdRepository.GetTodayCreateAds() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("AdRepository.GetTodayCreateAds() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
