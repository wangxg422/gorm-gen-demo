package dao

import (
	"gorm-gen-demo/client"
	"gorm-gen-demo/dal/query"
	"gorm-gen-demo/plugin"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	client.DB = client.NewClient().Debug()
	// 注册数据权限插件
	client.DB.Use(&plugin.DataScopePlugin{})

	query.SetDefault(client.DB)

	// 运行所有测试
	code := m.Run()

	os.Exit(code)
}
