package main

import (
	"gorm-gen-demo/client"
	"gorm-gen-demo/dal/query"
	"gorm-gen-demo/plugin"
)

func main() {
	client.DB = client.NewClient().Debug()
	// 注册数据权限插件
	client.DB.Use(&plugin.DataScopePlugin{})

	query.SetDefault(client.DB)
}
