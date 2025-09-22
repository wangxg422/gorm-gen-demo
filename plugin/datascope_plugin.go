package plugin

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DataScopePlugin struct{}

func (p *DataScopePlugin) Name() string {
	return "DataScopePlugin"
}

func (p *DataScopePlugin) Initialize(db *gorm.DB) (err error) {
	db.Callback().Query().Before("gorm:query").Register("data_scope:query", func(db *gorm.DB) {
		cfg, ok := GetScopeConfig(db.Statement.Context)
		if !ok {
			return
		}

		// 获取主表别名（没有就用表名）
		table := db.Statement.Table
		alias := cfg.TableAlias[table]
		if alias == "" {
			alias = table
		}

		column := func(col string) string {
			return alias + "." + col
		}

		switch cfg.Scope {
		case ScopeSelf:
			db.Statement.AddClause(clause.Where{
				Exprs: []clause.Expression{
					clause.Eq{Column: column("user_id"), Value: cfg.UserID},
				},
			})
		case ScopeDept:
			db.Statement.AddClause(clause.Where{
				Exprs: []clause.Expression{
					clause.Eq{Column: column("dept_id"), Value: cfg.DeptID},
				},
			})
		case ScopeDeptAndChildren:
			if len(cfg.DeptIDs) == 0 {
				// 无部门权限 —— 强制返回空集
				db.Statement.AddClause(clause.Where{
					Exprs: []clause.Expression{
						clause.Expr{SQL: "1=0"},
					},
				})
				return
			}
			// 推荐用 db.Where，GORM 会自动处理 slice 类型
			db.Where(fmt.Sprintf("%s.dept_id IN ?", alias), cfg.DeptIDs)

		case ScopeAll:
			// 不加条件
		}
	})

	return
}
