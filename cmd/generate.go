package main

import (
	"flag"
	"fmt"
	"gorm-gen-demo/cmd/mixin"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
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
	
	g := gen.NewGenerator(gen.Config{
		OutPath: "../dal/query",

		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode

		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: true, // generate pointer when field is nullable
	})

	gormdb, _ := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, addr, port, database)))
	g.UseDB(gormdb) // reuse your gorm db

	// 先定义模型的基础结构，只生成基础字段，不处理关系，确保模型都先存在，避免定义关系时循环引用冲突
	SysUser := g.GenerateModel("sys_user")
	SysRole := g.GenerateModel("sys_role")
	SysDept := g.GenerateModel("sys_dept")

	// 使用 ApplyBasic 后，单独追加关联（使用 ApplyInterface 或 GenerateModel 的补充模式）

	// 定义 SysUser 的关系
	SysUserRelate := g.GenerateModelAs("sys_user", "SysUser",
		mixin.AutoCreateTimeTag, mixin.AutoUpdateTimeTag,
		mixin.SoftDeleteFieldType, mixin.SoftDeleteFlagTag,
		// M2M: SysUser - SysRole
		gen.FieldRelate(field.Many2Many, "SysRoleList", SysRole, &field.RelateConfig{
			GORMTag: field.GormTag{
				mixin.M2MLabel:            []string{"sys_user_role"},
				mixin.JoinForeignKeyLabel: []string{"UserID"},
				mixin.JoinReferencesLabel: []string{"RoleID"},
			},
		}),
		// M2O: SysUser - SysDept
		gen.FieldRelate(field.BelongsTo, "BelongToSysDept", SysDept, &field.RelateConfig{
			GORMTag: field.GormTag{
				mixin.ForeignKeyLabel: []string{"ID"},
				mixin.ReferencesLabel: []string{"DeptID"},
			},
			RelatePointer: true,
		}),
	)

	// 定义 SysRole 的关系
	SysRoleRelate := g.GenerateModelAs("sys_role", "SysRole",
		mixin.AutoCreateTimeTag, mixin.AutoUpdateTimeTag,
		mixin.SoftDeleteFieldType, mixin.SoftDeleteFlagTag,
		// M2M: SysRole - SysUser
		gen.FieldRelate(field.Many2Many, "SysUserList", SysUser, &field.RelateConfig{
			GORMTag: field.GormTag{
				mixin.M2MLabel:            []string{"sys_user_role"},
				mixin.JoinForeignKeyLabel: []string{"RoleID"},
				mixin.JoinReferencesLabel: []string{"UserID"},
			},
		}),
	)

	// 定义 SysDept 的关系
	SysDeptRelate := g.GenerateModelAs("sys_dept", "SysDept",
		mixin.AutoCreateTimeTag, mixin.AutoUpdateTimeTag,
		mixin.SoftDeleteFieldType, mixin.SoftDeleteFlagTag,
		// O2M: SysDept - SysUser
		gen.FieldRelate(field.HasMany, "SysUserList", SysUser, &field.RelateConfig{
			GORMTag: field.GormTag{
				mixin.ForeignKeyLabel: []string{"ID"},
				mixin.ReferencesLabel: []string{"ID"},
			},
			RelateSlicePointer: true,
		}),
	)

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier){}, model.User{}, model.Company{})
	g.ApplyBasic(SysUserRelate, SysRoleRelate, SysDeptRelate)
	//g.ApplyBasic(g.GenerateAllTable()...)

	// Generate the code
	g.Execute()
}
