package services

import (
	"advertising/repositories"
)

type AdServiceInterface interface {
}

type AdService struct {
	AdRepository repositories.AdRepositoryInterface
}

func NewAdService() *AdService {
	return &AdService{
		AdRepository: repositories.NewAdRepository(),
	}
}

// CreateUser
// func (h *UserService) CreateUser(data *model.Student) (student_id int, status string) {
// 	pwd := []byte(data.Password)
// 	hash := NewUserService().HashAndSalt(pwd)
// 	data.Password = hash
// 	student_id, db := repository.NewUserRepository().Create(data)
// 	if db.Error != nil {
// 		return -1, responses.Error
// 	}
// 	return student_id, responses.Success
// }

// // scoreSearch
// func (h *UserService) ScoreSearch(requestData string, user_id uint) (student []interface{}, status string) {
// 	str_user_id := strconv.Itoa(int(user_id))
// 	// [Token用]:限制只有本人能查詢分數，如果Token login時所暫存的user_id與傳入c的user_id不相符，則回傳只限本人查詢分數。
// 	if str_user_id != requestData {
// 		return nil, responses.ScoreTokenErr
// 	}
// 	redisKey := fmt.Sprintf("user_%s", requestData)

// 	// 如果抓取redis的過程有error就跑進service並重新設置redis
// 	dbData, err := NewUserService().GetRedisKey(redisKey)
// 	if err != nil {
// 		student := repository.NewUserRepository().ScoreSearch(requestData)
// 		// 加密成JSON檔，用ffjson比普通的json還快
// 		redisData, _ := ffjson.Marshal(student)
// 		err = NewUserService().SetRedisKey(redisKey, redisData)
// 		if err != nil {
// 			return student, responses.Error
// 		}
// 		return student, responses.SuccessDb
// 	} else {
// 		var studentRedis []interface{}
// 		// 將Byte解密映射到studentRedis上
// 		ffjson.Unmarshal(dbData, &studentRedis)
// 		return studentRedis, responses.SuccessRedis
// 	}
// }
