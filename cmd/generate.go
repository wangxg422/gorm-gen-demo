package main

import (
	"gorm-gen-demo/client"
	"strings"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../dal/query",

		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	//gormdb, _ := gorm.Open(mysql.Open("root:passowrd@(127.0.0.1:3306)/gorm_gen?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(client.ConnectDB()) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	//g.ApplyBasic(model.User{})

	// 自定义模型结体字段的标签
	// 将特定字段名的 json 标签加上`string`属性,即 MarshalJSON 时该字段由数字类型转成字符串类型
	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
		toStringField := `balance, `
		if strings.Contains(toStringField, columnName) {
			return columnName + ",string"
		}
		return columnName
	})
	// 将非默认字段名的字段定义为自动时间戳和软删除字段;
	// 自动时间戳默认字段名为:`updated_at`、`created_at, 表字段数据类型为: INT 或 DATETIME
	// 软删除默认字段名为:`deleted_at`, 表字段数据类型为: DATETIME
	autoUpdateTimeField := gen.FieldGORMTag("update_time", func(tag field.GormTag) field.GormTag {
		return field.GormTag{
			"update_time":    []string{},
			"column":         []string{"update_time"},
			"type":           []string{"int unsigned"},
			"autoUpdateTime": []string{},
		}
	})
	autoCreateTimeField := gen.FieldGORMTag("create_time", func(tag field.GormTag) field.GormTag {
		return field.GormTag{
			"create_time":    []string{},
			"column":         []string{"create_time"},
			"type":           []string{"int unsigned"},
			"autoCreateTime": []string{},
		}
	})
	autoDeleteTimeField := gen.FieldGORMTag("delete_time", func(tag field.GormTag) field.GormTag {
		return field.GormTag{
			"delete_time":    []string{},
			"column":         []string{"delete_time"},
			"type":           []string{"int unsigned"},
			"autoDeleteTime": []string{},
		}
	})
	softDeleteField := gen.FieldType("delete_time", "soft_delete.DeletedAt")

	fieldID := gen.FieldNewTag("id", field.Tag{
		"json": "id",
		"gorm": "column:id;primaryKey;autoIncrement:true;type:int unsigned"})

	// 先定义模型的基础结构，只生成基础字段，不处理关系，确保模型都先存在，避免循环引用冲突
	AppInstance := g.GenerateModel("app_instance",
		fieldID,
		gen.FieldNewTag("app_id", field.Tag{
			"json": "app_id",
			"gorm": "column:app_id"}),
		gen.FieldNewTag("app_name", field.Tag{
			"json": "app_name",
			"gorm": "column:app_name"}),
		gen.FieldNewTag("app_version", field.Tag{
			"json": "app_version",
			"gorm": "column:app_version"}),
		jsonField, autoCreateTimeField, autoDeleteTimeField, autoUpdateTimeField, softDeleteField,
	)

	Role := g.GenerateModel("role",
		fieldID,
		jsonField, autoCreateTimeField, autoDeleteTimeField, autoUpdateTimeField, softDeleteField,
		gen.FieldNewTag("role_code", field.Tag{
			"json": "app_code",
			"gorm": "column:role_code"}),
		gen.FieldNewTag("role_name", field.Tag{
			"json": "role_name",
			"gorm": "column:role_name"}),
		jsonField, autoCreateTimeField, autoDeleteTimeField, autoUpdateTimeField, softDeleteField,
	)

	User := g.GenerateModel("user",
		fieldID,
		jsonField, autoCreateTimeField, autoDeleteTimeField, autoUpdateTimeField, softDeleteField,
		gen.FieldNewTag("user_name", field.Tag{
			"json": "user_name",
			"gorm": "column:user_name"}),
		gen.FieldNewTag("real_name", field.Tag{
			"json": "real_name",
			"gorm": "column:real_name"}),
		gen.FieldNewTag("password", field.Tag{
			"json": "password",
			"gorm": "column:password"}),
		jsonField, autoCreateTimeField, autoDeleteTimeField, autoUpdateTimeField, softDeleteField,
	)

  // 

  //g.ApplyBasic(g.GenerateAllTable()...)
  g.ApplyBasic(AppInstance, Role, User)

	// 使用 ApplyBasic 后，单独追加关联（使用 ApplyInterface 或 GenerateModel 的补充模式）

	// 给 User 追加关系
	UserRelate := g.GenerateModelAs("user", "User",
		gen.FieldRelate(field.HasMany, "AppInstances", AppInstance, &field.RelateConfig{
			GORMTag: field.GormTag{
				"foreignKey": []string{"create_user_id"},
			},
		}),
		gen.FieldRelate(field.Many2Many, "Roles", Role, &field.RelateConfig{
			GORMTag: field.GormTag{
				"many2many":      []string{"user_role"},
				"joinForeignKey": []string{"UserID"},
				"JoinReferences": []string{"RoleID"},
			},
		}),
	)

	// 给 AppInstance 追加关系
	AppInstanceRelate := g.GenerateModelAs("app_instance", "AppInstance",
		gen.FieldRelate(field.BelongsTo, "CreateUser", User, &field.RelateConfig{
			GORMTag: field.GormTag{
				"foreignKey": []string{"create_user_id"},
			},
		}),
	)

	// 给 Role 追加关系
	RoleRelate := g.GenerateModelAs("role", "Role",
		gen.FieldRelate(field.Many2Many, "Users", User, &field.RelateConfig{
			GORMTag: field.GormTag{
				"many2many":      []string{"user_role"},
				"joinForeignKey": []string{"RoleID"},
				"JoinReferences": []string{"UserID"},
			},
		}),
	)

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier){}, model.User{}, model.Company{})
	g.ApplyBasic(UserRelate, AppInstanceRelate, RoleRelate)

	// Generate the code
	g.Execute()
}
