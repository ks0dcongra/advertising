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
func (a *AdController) GetAd() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set offset default is 0, if get enpty string
		adOffset, err := strconv.Atoi(c.Query("offset"))
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.Status(define.ParameterErr, nil))
			return
		}

		adLimit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.Status(define.ParameterErr, nil))
			return
		}

		// Valid age that must be a number
		age, err := strconv.Atoi(c.Query("age"))
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.Status(define.ParameterErr, nil))
			return
		}

		requestData := &requests.ConditionInfoOfPage{
			AdOffset: adOffset,
			AdLimit:  adLimit,
			Age:      age,
			Gender:   c.Query("gender"),
			Country:  c.Query("country"),
			Platform: c.Query("platform"),
		}

		ads, status := services.NewAdService().GetAd(requestData)
		if status != define.Success  {
			c.JSON(http.StatusBadRequest, responses.Status(status, nil))		
			return
		}

		c.JSON(http.StatusOK, responses.Status(status, ads))
	}
}
