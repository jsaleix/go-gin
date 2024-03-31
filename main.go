package main

import (
	"fmt"
	"sse/routes"
	"sse/sse"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	stream := sse.NewServer()

	go func() {
		for {
			time.Sleep(time.Second * 10)
			now := time.Now().Format("2006-01-02 15:04:05")
			currentTime := fmt.Sprintf("The Current Time Is %v", now)

			// Send current time to clients message channel
			stream.Message <- currentTime
		}
	}()

	routes.AffectRoutes(router, stream)
	sse.InitRoute(router, stream)

	router.Run("localhost:8080")
}
