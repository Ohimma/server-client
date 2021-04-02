package router

import (
	"github.com/Ohimma/server-client/http/server/controller/health"
	"github.com/gin-gonic/gin"
)

func RouterHealth(group *gin.RouterGroup) {
	router := group.Group("")
	{
		router.GET("/health", health.HealthCheck)
		router.POST("/health", health.AddHealth)
	}
}
