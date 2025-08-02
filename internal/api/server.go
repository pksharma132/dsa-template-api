package api

import (
	"dsa-template-api/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/template", handler.HandleTemplate)
	}
	router.Run(":8080")
}
