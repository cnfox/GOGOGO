package model

import (
	"GOGOGO/config"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DB2.0 数据库连接 gorm2.0
var DB2 *gorm.DB

func initDB2() {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(config.DBConfig.URL)
	db, err := gorm.Open(mysql.Open(config.DBConfig.URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	DB2 = db
}
