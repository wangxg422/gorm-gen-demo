package main

import (
	"gorm-gen-demo/dal"
	"gorm-gen-demo/dal/query"
)

func init() {
	dal.DB = dal.ConnectDB().Debug()
}

func main() {
	query.SetDefault(dal.DB)
}
