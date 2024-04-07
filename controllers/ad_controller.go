package controllers

import (
	"advertising/define"
	"advertising/models/requests"
	"advertising/models/responses"
	"advertising/services"
	"log"
	"net/http"
	"strconv"

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
			log.Println("Error:" + err.Error())
			c.JSON(http.StatusBadRequest, responses.Status(define.ParameterErr, nil))
			return
		}
		status := services.NewAdService().CreateAd(requestData)

		c.JSON(http.StatusBadRequest, responses.Status(status, nil))
	}
}

// ScoreSearch
func (a *AdController) GetAds() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
			1. Valid offset,limit, age that must be a number
		 	2. if get enpty string in int datatype column, set them to `0` 
		*/
		adOffset, err := strconv.Atoi(c.Query("offset"))
		if err != nil {
			if adOffset != 0 {
				c.JSON(http.StatusBadRequest, responses.Status(define.ParameterErr, nil))
				return
			}
		}

	
		adLimit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			if adOffset != 0 {
				c.JSON(http.StatusBadRequest, responses.Status(define.ParameterErr, nil))
				return
			}
		}

		age, err := strconv.Atoi(c.Query("age"))
		if err != nil {
			if adOffset != 0 {
				c.JSON(http.StatusBadRequest, responses.Status(define.ParameterErr, nil))
				return
			}
		}

		requestData := &requests.ConditionInfoOfPage{
			AdOffset: adOffset,
			AdLimit:  adLimit,
			Age:      age,
			Gender:   c.Query("gender"),
			Country:  c.Query("country"),
			Platform: c.Query("platform"),
		}

		ads, status := services.NewAdService().GetAds(requestData)

		if status == define.Success || status == define.RedisSuccess {
			c.JSON(http.StatusOK, responses.Status(status, ads))
			return
		}

		c.JSON(http.StatusBadRequest, responses.Status(status, nil))
	}
}
