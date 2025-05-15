package main

import (
	"gorm-gen-demo/client"
	"gorm-gen-demo/dal/query"
)

func init() {
	client.DB = client.ConnectDB().Debug()
}

func main() {
	query.SetDefault(client.DB)
}
