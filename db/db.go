package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func init() {
	Database = Connect()
}

func InitTable(dst any) {
	err := Database.AutoMigrate(dst)
	if err != nil {
		fmt.Println("初始化数据库表失败")
		panic(err)
	}
}

// Connect 数据库连接
func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=123456 port=5432 dbname=cms_dev sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}

	return db
}
