package controllers

import (
	"advertising/define"
	"advertising/models/requests"
	"advertising/models/responses"
	"advertising/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdController struct {
	AdService services.AdServiceInterface
}

func NewAdController() *AdController {
	return &AdController{
		AdService: services.NewAdService(),
	}
}

func (a *AdController) CreateAd() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestData := new(requests.CreateAd)
		if err := c.ShouldBindJSON(&requestData); err != nil {
			log.Println("Error1:" + err.Error())
			c.JSON(http.StatusBadRequest, responses.Status(define.ParameterErr, nil))
			return
		}
		status := a.AdService.CreateAd(requestData)

		c.JSON(http.StatusBadRequest, responses.Status(status, nil))
	}
}

func (a *AdController) GetAds() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestData := new(requests.ConditionInfoOfPage)
		if err := c.ShouldBindQuery(&requestData); err != nil {
			log.Println("Error3:" + err.Error())
			c.JSON(http.StatusBadRequest, responses.Status(define.ParameterErr, nil))
			return
		}

		ads, status := a.AdService.GetAds(requestData)

		if status == define.Success || status == define.RedisSuccess {
			c.JSON(http.StatusOK, responses.Status(status, ads))
			return
		}

		c.JSON(http.StatusBadRequest, responses.Status(status, nil))
	}
}
