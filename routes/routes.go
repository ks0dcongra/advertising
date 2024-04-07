package routes

import (
	"github.com/gin-gonic/gin"
	"advertising/controllers"
)

func ApiRoutes(router *gin.Engine) {
	adAPI := router.Group("api/v1")

	adAPI.POST("ad", controllers.NewAdController().CreateAd())
	adAPI.GET("ad", controllers.NewAdController().GetAds())
}
