package common

import (
	"database/sql"
	"pan/global"

	"gorm.io/gorm"
)

func GetGormDB() *gorm.DB {
	return global.Panserver.DB
}

func GetSqlDB() (*sql.DB, error) {
	db, err := GetGormDB().DB()
	if err != nil {
		return nil, err
	}
	return db, err
}
