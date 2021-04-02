package health

import (
	"github.com/Ohimma/server-client/http/server/handler"
	"github.com/Ohimma/server-client/http/server/middleware"
	"github.com/Ohimma/server-client/http/server/model"
	"github.com/gin-gonic/gin"
)

type HealthRequest struct {
	Id   uint   `json:"id" form:"id"`
	Host string `json:"host" form:"host"`
}

func HealthCheck(c *gin.Context) {
	middleware.Logger.Info("success")
	handler.ResponseOK(c, 200, "health", "success")

}

func AddHealth(c *gin.Context) {
	// 1. 验证数据格式
	var req HealthRequest
	if err := c.Bind(&req); err != nil {
		handler.ResponseError(c, 500, err, "解析请求参数失败")
		return
	}
	middleware.Logger.Info("req = ", req)
	if req.Host == "" {
		handler.ResponseError(c, 500, req, "请求参数错误")
		return
	}

	Info, count, err := model.ExitsHealth(req.Host)
	if err != nil {
		middleware.Logger.Fatal("查询遇到其他错误，退出", err)
	} else {
		if count == 0 {
			middleware.Logger.Warn("记录不存在 is ", Info, count, err)
			db := model.Health{
				Id:   Info.Id,
				Host: req.Host,
			}
			Info, count, err = model.CreateHealth(&db)
			if err != nil {
				middleware.Logger.Info("创建失败 = ", err)
				handler.ResponseError(c, 500, err, "创建失败")
				return
			}
		} else {
			middleware.Logger.Warn("记录已经存在 is ", Info, count, err)
			db := model.Health{
				Id:   Info.Id,
				Host: req.Host,
			}
			Info, count, err = model.UpdateHealth(&db)
			if err != nil {
				middleware.Logger.Info("更新失败 = ", err)
				handler.ResponseError(c, 500, err, "更新失败")
				return
			}
		}
	}

	middleware.Logger.Info("success = ", Info)
	handler.ResponseOK(c, 200, Info, "success")
}

func UpdateHealth(db *model.Health) {
	// 1. 验证数据格式
	// var req HealthRequest
	// if err := c.Bind(&req); err != nil {
	// 	handler.ResponseError(c, 500, err, "解析请求参数失败")
	// 	return
	// }
	// middleware.Logger.Info("req = ", req)
	// if req.Host == "" {
	// 	handler.ResponseError(c, 500, req, "请求参数错误")
	// 	return
	// }

	// Info, count, err := model.ExitsHealth(req.Host)
	// if err != nil || count > 0 {
	// db := model.Health{
	// 	Id:   data.Id,
	// 	Host: Info.Host,
	// }
	Info, count, err := model.UpdateHealth(db)
	if err != nil {
		middleware.Logger.Info("创建失败 = ", err, count)
		return
	}
	middleware.Logger.Info("success = ", Info)
	return
	// handler.ResponseOK(c, 200, Info, "success")
	// }

	// // 3. 合并要插入的数据
	// db := model.Health{
	// 	Id:   req.Id,
	// 	Host: req.Host,
	// }

	// // 4. 将用户插入数据库
	// if err := model.CreateHealth(&db); err != nil {
	// 	middleware.Logger.Info("创建失败 = ", err)
	// 	handler.ResponseError(c, 500, err, "创建失败")
	// 	return
	// }

	// middleware.Logger.Info("success = ", req)
	// handler.ResponseOK(c, 200, req, "success")
}
