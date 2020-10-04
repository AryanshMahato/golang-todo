package controllers

import "github.com/gin-gonic/gin"

func HealthController(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"alive": true,
	})
}
