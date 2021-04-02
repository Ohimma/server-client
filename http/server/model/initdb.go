package model

import (
	"fmt"

	"github.com/Ohimma/server-client/http/server/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		config.Conf.Mysql.User,
		config.Conf.Mysql.Password,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Db,
		config.Conf.Mysql.Config,
	)
	fmt.Println("initmysql dsn = ", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connect mysql succuss", db)

	// 迁移 添加测试数据
	if config.Conf.Mysql.AutoMigrate {
		db.AutoMigrate(new(Health))

		user := []Health{
			{Host: "127.0.0.2"},
			{Host: "127.0.0.3"},
		}
		result_user := db.Create(&user) // 通过数据的指针来创建
		fmt.Println("result2 = ", &result_user, result_user.Error)

		// 1. 验证该用户是否存在
		// result := db.Where("name = 'admin'").Find(&UserUser{})
		// if result.RowsAffected == 0 {
		// 	middleware.Logger.Error("未初始化数据库 = ", result.Error, result.RowsAffected)

		// 	user := []UserUser{
		// 		{Name: "admin", Phone: "1234", RoleIds: "0", Password: "d59dd9b2dd80ac58", Salt: "c15d52b3-d71d-413e-b671-dc1adab56c78", Avatar: "admin"},
		// 		{Name: "user", Phone: "12345", RoleIds: "2", Password: "a151b466e68e9724", Salt: "14e35c41-b8d3-437e-b05d-792e9ba418d0", Avatar: "test"},
		// 		{Name: "view", Phone: "view123", RoleIds: "3", Password: "cd4bc6ec70f6178f", Salt: "be2f459d-2d7f-460d-8670-0ff5ffaf11c7", Avatar: "view"},
		// 	}
		// 	result_user := db.Create(&user) // 通过数据的指针来创建
		// 	fmt.Println("result2 = ", &result_user, result_user.Error)
		// }
	}

	DB = db
}
