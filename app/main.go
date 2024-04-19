package main

import (
	"api/config"
	"api/db"
	"api/routes"
	"api/sse"
	"fmt"
	"log"
	"time"

	"api/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	db.Init()
	router := gin.Default()
	stream := sse.NewServer()
	docs.SwaggerInfo.BasePath = "/"

	router.Use(gin.Logger())

	go func() {
		for {
			time.Sleep(time.Second * 10)
			now := time.Now().Format("2006-01-02 15:04:05")
			currentTime := fmt.Sprintf("The Current Time Is %v", now)

			// Send current time to clients message channel
			stream.Message <- currentTime
		}
	}()

	routes.AffectRoutes(router)
	sse.InitRoute(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if err := router.Run(":" + config.PORT); err != nil {
		log.Panicf("error: %s", err)
	}
}
