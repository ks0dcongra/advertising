package controllers

import (
	"advertising/services"
)

type AdController struct {
	AdService services.AdServiceInterface
}

func NewAdController() *AdController {
	return &AdController{
		AdService: services.NewAdService(),
	}
}

// Create User
// func (h *UserController) CreateUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		requestData := new(model.Student)
// 		if err := c.ShouldBindJSON(&requestData); err != nil {
// 			fmt.Println("Error:" + err.Error())
// 			c.JSON(http.StatusNotAcceptable, responses.Status(responses.ParameterErr, nil))
// 			return
// 		}
// 		student_id, status := service.NewUserService().CreateUser(requestData)
// 		if status != responses.Success {
// 			c.JSON(http.StatusNotFound, responses.Status(responses.Error, nil))
// 			return
// 		}
// 		c.JSON(http.StatusOK, responses.Status(responses.Success, student_id))
// 	}
// }

// // ScoreSearch
// func (h *UserController) ScoreSearch() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		requestData := c.Param("id")
// 		if requestData == "0" || requestData == "" {
// 			c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil))
// 			return
// 		}

// 		// 創建 JwtFactory 實例
// 		JwtFactory := token.Newjwt()
// 		// [Token用]:先將uint轉換成int再運用strconv轉換成string。
// 		user_id, err := JwtFactory.ExtractTokenID(c)
// 		// [Token用]:Token出錯了!
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, responses.Status(responses.TokenErr, nil))
// 		}

// 		student, status := service.NewUserService().ScoreSearch(requestData, user_id)

// 		if status == responses.SuccessDb || status == responses.SuccessRedis {
// 			c.JSON(http.StatusOK, responses.Status(status, student))
// 		} else {
// 			c.JSON(http.StatusNotFound, responses.Status(status, student))
// 		}
// 	}
// }
