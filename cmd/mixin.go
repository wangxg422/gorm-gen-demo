package main

import (
	"strings"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

// 自定义模型结体字段的标签
//
//	将特定字段名的 json 标签加上`string`属性,即 MarshalJSON 时该字段由数字类型转成字符串类型
var jsonField = gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
	toStringField := `balance, `
	if strings.Contains(toStringField, columnName) {
		return columnName + ",string"
	}
	return columnName
})

// 将非默认字段名的字段定义为自动时间戳和软删除字段;
// 自动时间戳默认字段名为:`updated_at`、`created_at`、`delete_at`, 表字段数据类型为: DATETIME
var autoUpdateTimeField = gen.FieldGORMTag("update_time", func(tag field.GormTag) field.GormTag {
	return field.GormTag{
		"update_time":    []string{},
		"column":         []string{"update_time"},
		"type":           []string{"int unsigned"},
		"autoUpdateTime": []string{},
	}
})
var autoCreateTimeField = gen.FieldGORMTag("create_time", func(tag field.GormTag) field.GormTag {
	return field.GormTag{
		"create_time":    []string{},
		"column":         []string{"create_time"},
		"type":           []string{"int unsigned"},
		"autoCreateTime": []string{},
	}
})
var autoDeleteTimeField = gen.FieldGORMTag("delete_time", func(tag field.GormTag) field.GormTag {
	return field.GormTag{
		"delete_time":    []string{},
		"column":         []string{"delete_time"},
		"type":           []string{"int unsigned"},
		"autoDeleteTime": []string{},
	}
})

// 软删除默认字段名为:`deleted_at`, 表字段数据类型为: DATETIME，这里指定为`valid`,数据类型为char(1)
var softDeleteField = gen.FieldType("valid", "string")

var fieldOpts = []gen.ModelOpt{jsonField, autoCreateTimeField, autoUpdateTimeField, autoDeleteTimeField, softDeleteField}
