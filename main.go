package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"project_demo/cmd/controller"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/QUAN_LY_DON_HANG?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("can't connect to database")
	}
	controller.Menu(db)
}
