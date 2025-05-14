package dal

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const MySQLDSN = "root:XishangAdmin@tcp(10.66.66.66:13307)/gorm_gen?charset=utf8mb4&parseTime=True"

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}
