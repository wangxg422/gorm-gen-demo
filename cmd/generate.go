package main

import (
	"gorm-gen-demo/client"
	"gorm-gen-demo/mixin"

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

		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: true, // generate pointer when field is nullable
	})

	//gormdb, _ := gorm.Open(mysql.Open("root:passowrd@(127.0.0.1:3306)/gorm_gen?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(client.ConnectDB()) // reuse your gorm db

	// 先定义模型的基础结构，只生成基础字段，不处理关系，确保模型都先存在，避免定义关系时循环引用冲突
	AppPackage := g.GenerateModel("app_package")
	AppInstance := g.GenerateModel("app_instance")
	Role := g.GenerateModel("role")
	User := g.GenerateModel("user")

	// 这里需要先生成模型，在定义完关系后在生成模型会被覆盖
	g.ApplyBasic(g.GenerateAllTable()...)

	// 使用 ApplyBasic 后，单独追加关联（使用 ApplyInterface 或 GenerateModel 的补充模式）

	// 定义 User 的关系
	UserRelate := g.GenerateModelAs("user", "User",
		mixin.AutoCreateTimeTag, mixin.AutoUpdateTimeTag,
		mixin.SoftDeleteFieldType, mixin.SoftDeleteFlagTag,
		gen.FieldRelate(field.Many2Many, "Roles", Role, &field.RelateConfig{
			GORMTag: field.GormTag{
				"many2many":      []string{"user_role"},
				"joinForeignKey": []string{"UserID"},
				"JoinReferences": []string{"RoleID"},
			},
		}),
		gen.FieldRelate(field.HasMany, "AppInstances", AppInstance, &field.RelateConfig{
			GORMTag: field.GormTag{
				"foreignKey": []string{"CreateUserID"},
			},
		}),
		gen.FieldRelate(field.HasMany, "AppPackages", AppPackage, &field.RelateConfig{
			GORMTag: field.GormTag{
				"foreignKey": []string{"CreateUserID"},
			},
		}),
	)

	// 定义 Role 的关系
	RoleRelate := g.GenerateModelAs("role", "Role",
		mixin.AutoCreateTimeTag, mixin.AutoUpdateTimeTag,
		mixin.SoftDeleteFieldType, mixin.SoftDeleteFlagTag,
		gen.FieldRelate(field.Many2Many, "Users", User, &field.RelateConfig{
			GORMTag: field.GormTag{
				"many2many":      []string{"user_role"},
				"joinForeignKey": []string{"RoleID"},
				"JoinReferences": []string{"UserID"},
			},
		}),
	)

	// 定义 AppPackage 的关系
	AppPackageRelate := g.GenerateModelAs("app_package", "AppPackage",
		mixin.AutoCreateTimeTag, mixin.AutoUpdateTimeTag,
		mixin.SoftDeleteFieldType, mixin.SoftDeleteFlagTag,
		gen.FieldRelate(field.HasMany, "AppInstance", AppInstance, &field.RelateConfig{
			GORMTag: field.GormTag{
				"foreignKey": []string{"AppPackageID"},
			},
		}),
		gen.FieldRelate(field.BelongsTo, "CreateUser", User, &field.RelateConfig{
			GORMTag: field.GormTag{
				"foreignKey": []string{"CreateUserID"},
			},
		}),
	)

	// 定义 AppInstance 的关系
	AppInstanceRelate := g.GenerateModelAs("app_instance", "AppInstance",
		mixin.AutoCreateTimeTag, mixin.AutoUpdateTimeTag,
		mixin.SoftDeleteFieldType, mixin.SoftDeleteFlagTag,
		gen.FieldRelate(field.BelongsTo, "CreateUser", User, &field.RelateConfig{
			GORMTag: field.GormTag{
				"foreignKey": []string{"CreateUserID"},
			},
		}),
		gen.FieldRelate(field.BelongsTo, "AppPackage", AppPackage, &field.RelateConfig{
			GORMTag: field.GormTag{
				"foreignKey": []string{"AppPackageID"},
			},
		}),
	)

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier){}, model.User{}, model.Company{})
	g.ApplyBasic(UserRelate, RoleRelate, AppPackageRelate, AppInstanceRelate)

	// Generate the code
	g.Execute()
}
