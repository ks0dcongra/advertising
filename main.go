package main

import (
	"advertising/configs"
	"advertising/routes"
	"fmt"
	"net/http"
	"os"

	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"context"
	"os/signal"
	"syscall"
)
func init() {
	var err error
	if err = configs.InitCommonConf(); err != nil {
		log.Fatalf("Failed to initialize common configurations: %v", err)
	}

	if err = configs.InitConfigs(); err != nil {
		log.Fatalf("Failed to initialize configurations: %v", err)
	}	
}

func main() {
	// connect database
	if err := configs.DBsetup(); err != nil {
		log.Fatalf("Connect to PostgreSQL failed: %v", err)
	}

	if err := configs.RedisSetup(); err != nil {
		log.Fatalf("Connect to Redis failed: %v", err)
	}

	mainServer := gin.New()

	// 定義router呼叫格式與跨域限制
	mainServer.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin"},
		ExposeHeaders:    []string{"Content-Type", "application/javascript"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 連接Router
	routes.ApiRoutes(mainServer)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", configs.ServiceInfo.ServerPort),
		Handler: mainServer,
	}

	log.Printf("(Main HTTPS) listen on :%s\n", configs.ServiceInfo.ServerPort)
	
	go func() {
		// 開啟port
		if err := mainServer.Run(":" + configs.ServiceInfo.ServerPort); err != nil {
			log.Fatalf("HTTP service failed: %v", err)
		}
	}()

	// gracefully shutdown 
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown: %v", err)
	}

	log.Println("Server exiting")
}
