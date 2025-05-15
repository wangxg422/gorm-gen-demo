package client

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const MySQLDSN = "root:123456@tcp(127.0.0.1:3306)/gorm-gen-demo?charset=utf8mb4&parseTime=True"

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}
