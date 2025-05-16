package main

import (
	"context"
	"gorm-gen-demo/client"
	"gorm-gen-demo/dal/query"
	"gorm-gen-demo/dao"
)

func init() {
	client.DB = client.ConnectDB().Debug()
}

func main() {
	query.SetDefault(client.DB)

	dao.GetRolesByUserId(context.Background(), 1)
}
