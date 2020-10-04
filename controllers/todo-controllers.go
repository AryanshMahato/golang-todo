package controllers

import (
	"GoLang/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func HealthController(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"alive": true,
	})
}

func GetTodosController(ctx *gin.Context) {
	todos, err := model.DB.Query("SELECT * FROM TODO")
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Cannot fetch todos",
			"error":   err,
			"todos":   todos,
		})
		fmt.Println(err)
		return
	}

	ctx.JSON(200, gin.H{
		"todos": todos,
	})
}
