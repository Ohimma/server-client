package router

import (
	"fmt"

	"github.com/Ohimma/server-client/http/server/config"
	"github.com/Ohimma/server-client/http/server/handler"
	"github.com/Ohimma/server-client/http/server/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	gin.SetMode(config.Conf.Mode)

	// 引用公共中间件
	router.Use(
		middleware.RequestLogger(),
	)

	// 404组
	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		// middleware.Logger.Warn("404 http request = ", c.Request)
		handler.ResponseFound(c, 404, nil, fmt.Sprintf("%s %s not found", method, path))
	})

	// 测试组， 把 t追加到 RouterHealth 里
	t := router.Group("/test")
	RouterHealth(t)

	return router
}
