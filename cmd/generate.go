package main

import (
	"gorm-gen-demo/dal"

	"gorm.io/gen"
)

// Dynamic SQL
type Querier interface {
  // SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
  FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
  g := gen.NewGenerator(gen.Config{
    OutPath: "../dal/query",
    Mode: gen.WithDefaultQuery|gen.WithQueryInterface, // generate mode
    // Mode: gen.WithoutContext|gen.WithDefaultQuery|gen.WithQueryInterface, // generate mode
  })

  //gormdb, _ := gorm.Open(mysql.Open("root:XishangAdmin@(10.66.66.66:13307)/gorm_gen?charset=utf8mb4&parseTime=True&loc=Local"))
  g.UseDB(dal.ConnectDB()) // reuse your gorm db

  // Generate basic type-safe DAO API for struct `model.User` following conventions
  //g.ApplyBasic(model.User{})

  // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
  //g.ApplyInterface(func(Querier){}, model.User{}, model.Company{})
  g.ApplyBasic(g.GenerateAllTable()...)

  // Generate the code
  g.Execute()
}