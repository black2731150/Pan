package initialize

import (
	"fmt"
	"pan/global"
	"pan/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//初始化mysql数据库
func initMysqlGorm() *gorm.DB {
	database := global.Panserver.Config.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s&%s&%s", database.Username, database.Password, database.Host, database.Port, database.Name, database.Options[0], database.Options[1], database.Options[2])
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	fmt.Println("Database initialize SUCCESS!")

	db.AutoMigrate(models.NewUser())

	return db
}

//初始化数据库
func InitDB() {
	switch global.Panserver.Config.Database.Driver {
	case "mysql":
		global.Panserver.DB = initMysqlGorm()
	default:
		panic("InitDB Error!")
	}
}
