package main

import (
	"GoLang/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	server := gin.Default()

	server.GET("/health", controllers.HealthController)

	err := server.Run()
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
