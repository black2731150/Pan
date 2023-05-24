package utils

import (
	"database/sql"
	"pan/global"

	"gorm.io/gorm"
)

//获取Gorm类型数据库连接
func GetGormDB() *gorm.DB {
	return global.Panserver.DB
}

//获取Sql类型数据库连接
func GetSqlDB() (*sql.DB, error) {
	db, err := GetGormDB().DB()
	if err != nil {
		return nil, err
	}
	return db, err
}
