package model

import (
	"time"

	"github.com/Ohimma/server-client/http/server/middleware"
)

// example
type Health struct {
	Id uint `json:"id" gorm:"primaryKey; autoIncrement; unique"`

	Host     string `json:"host" gorm:"size:32; uniqueIndex; not null;"`
	CreateAt string `json:"create_at" gorm:"not null; default:2021-01-01"`
	UpdateAt string `json:"update_at" gorm:"not null; default:2021-01-01"`
	DeleteAt string `json:"delete_at" gorm:"default: null; comment:value不为空时表示删除"`
}

func ExitsHealth(host string) (*Health, int64, error) {
	Info := &Health{}
	sql := DB.Where("host = ? ", host)
	result := sql.Find(&Info)

	middleware.Logger.Info("AuthExits result = ", sql, Info, result.RowsAffected, result.Error)
	return Info, result.RowsAffected, result.Error
}

func CreateHealth(db *Health) (*Health, int64, error) {
	db.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	result := DB.Create(&db)

	middleware.Logger.Info("CreateUsername = ", *db, result.RowsAffected, result.Error)
	return db, result.RowsAffected, result.Error
}

func UpdateHealth(db *Health) (*Health, int64, error) {
	db.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	sql := DB.Where("id = ? ", db.Id)
	result := sql.Updates(&db)

	middleware.Logger.Info("UpdateAuth = ", *db, result, result.Error)
	return db, result.RowsAffected, result.Error
}
