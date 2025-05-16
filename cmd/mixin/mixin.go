package mixin

import (
	"strings"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

const (
	AutoCreateTime = "autoCreateTime"
	AutoUpdateTime = "autoUpdateTime"
	AutoDeleteTime = "autoDeleteTime"
	SoftDelete     = "softDelete"
	Delflag = "del_flag"
	CreateTime = "create_time"
	UpdateTime = "update_time"
	DeleteTime = "delete_time"
)

// 自定义模型结体字段的标签
//	将特定字段名的 json 标签加上`string`属性,即 MarshalJSON 时该字段由数字类型转成字符串类型
var jsonField = gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
	toStringField := `balance, `
	if strings.Contains(toStringField, columnName) {
		return columnName + ",string"
	}
	return columnName
})

// 将非默认字段名的字段定义为自动时间戳和软删除字段;
// 自动时间戳默认字段名为:`updated_at`、`created_at`, 表字段数据类型为: DATETIME
var AutoCreateTimeTag = gen.FieldGORMTag(CreateTime, func(tag field.GormTag) field.GormTag {
	tag.Append(AutoCreateTime, "true")
	return tag
})

var AutoUpdateTimeTag = gen.FieldGORMTag(UpdateTime, func(tag field.GormTag) field.GormTag {
	tag.Append(AutoUpdateTime, "true")
	return tag
})

// var AutoDeleteTimeTag = gen.FieldGORMTag(DeleteTime, func(tag field.GormTag) field.GormTag {
// 	tag.Append(AutoDeleteTime, "true")
// 	return tag
// })

var SoftDeleteFlagTag = gen.FieldGORMTag(Delflag, func(tag field.GormTag) field.GormTag {
	tag.Append(SoftDelete, "flag")
	return tag
})

// 软删除默认字段名为:`deleted_at`, 表字段数据类型为: DATETIME. 这里指定为`valid`,数据类型为char(1)
var SoftDeleteFieldType = gen.FieldType(Delflag, "soft_delete.DeletedAt")

//var CommonTag = []gen.ModelOpt{AutoCreateTimeTag, AutoUpdateTimeTag, AutoDeleteTimeTag}
