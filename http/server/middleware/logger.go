package middleware

import (
	"os"
	"path"
	"time"

	"github.com/Ohimma/server-client/http/server/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	log := logrus.New()
	// 生产环境写入文件JSON格式
	if config.Conf.Server.Mode == "release" {
		fileName := path.Join("./logs", "odemo.log")
		_, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		// logger.SetOutput(bufio.NewWriter(file))
		// writeFile(fileName, config.MaxAge)

		if err != nil {
			panic(err)
		}
		//设置日志格式、级别、输出地方
		log.SetFormatter(&logrus.JSONFormatter{})
		log.SetLevel(logrus.InfoLevel)

	} else {
		log.SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
		log.WithFields(logrus.Fields{
			"app": "odemo",
		}).Info()
		log.SetLevel(logrus.DebugLevel)
		log.SetOutput(os.Stdout)
	}

	Logger = log
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		if uri == "/favicon.ico" {
			return
		}
		//开始时间
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		method := c.Request.Method
		statusCode := c.Writer.Status()
		ip := c.ClientIP()

		if config.Conf.Server.Mode == "release" {
			Logger.WithFields(logrus.Fields{
				"env":         "prod",
				"clientIp":    ip,
				"statusCode":  statusCode,
				"reqMethod":   method,
				"reqUri":      uri,
				"latencyTime": latencyTime,
			}).Info()
		} else {
			// now := time.Now().Format("2006-01-02 15:04:05")
			// Logger.Infof(" %3d | %13v | %15s | %s  %s",
			// 	statusCode,
			// 	latencyTime,
			// 	ip,
			// 	method,
			// 	uri,
			// )
			Logger.WithFields(logrus.Fields{"env": "dev", "clientIp": ip, "statusCode": statusCode, "reqMethod": method, "reqUri": uri, "latencyTime": latencyTime}).Info()
		}
	}
}
