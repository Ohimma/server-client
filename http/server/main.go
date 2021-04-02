package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Ohimma/server-client/http/server/config"
	"github.com/Ohimma/server-client/http/server/middleware"
	"github.com/Ohimma/server-client/http/server/model"
	"github.com/Ohimma/server-client/http/server/router"
)

func main() {
	fmt.Println("run server")

	// 1. 初始化数据库
	model.InitMysql()
	// if config.Conf.Server.UseRedis {
	// 	model.InitRedis()
	// }

	// 3. 初始化log
	// 4. 进入 gin 主程序
	runServer()
}

func runServer() {
	router1 := router.InitRouter()

	// 启动
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Conf.Server.Port),
		Handler: router1,
	}
	middleware.Logger.Info(fmt.Sprintf("Listening and serving HTTP on Port: %d, Pid: %d", config.Conf.Server.Port, os.Getpid()))
	server.ListenAndServe()

	// 直接关停
	middleware.Logger.Info("server exit .....")
}
