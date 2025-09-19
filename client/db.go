package client

import (
	"flag"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var (
		username string
		password string
		addr     string
		port     string
		database string
	)

	flag.StringVar(&username, "u", "root", "username")
	flag.StringVar(&password, "p", "123456", "password")
	flag.StringVar(&addr, "a", "127.0.0.1", "address")
	flag.StringVar(&port, "P", "3306", "port")
	flag.StringVar(&database, "d", "gorm-gen-demo", "database")
	flag.Parse()

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, addr, port, database)))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}
