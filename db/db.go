package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"server/config"
	"server/dal/query"
)

var Database *gorm.DB

func New() {
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
	conf := config.Config.Db
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v port=%v dbname=%v sslmode=disable TimeZone=Asia/Shanghai",
		conf.Host, conf.User, conf.Password, conf.Port, conf.Dbname,
	)
	fmt.Println("dsn", dsn)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}

	query.SetDefault(db)

	return db
}
