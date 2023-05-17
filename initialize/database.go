package initialize

import (
	"fmt"
	"pan/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initMysqlGorm() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s&%s&%s",
		global.Panserver.Config.Database.Username,
		global.Panserver.Config.Database.Password,
		global.Panserver.Config.Database.Host,
		global.Panserver.Config.Database.Port,
		global.Panserver.Config.Database.Name,
		global.Panserver.Config.Database.Options[0],
		global.Panserver.Config.Database.Options[1],
		global.Panserver.Config.Database.Options[2])
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	fmt.Println("Database initialize SUCCESS!")

	db.AutoMigrate(&global.Panserver.Config.Pan)

	return db
}

func InitDB() {
	switch global.Panserver.Config.Database.Driver {
	case "mysql":
		global.Panserver.DB = initMysqlGorm()
	}
}
