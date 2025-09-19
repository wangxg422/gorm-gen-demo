package mixin

import (
	"strings"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

const (
	PrimaryKeyLabel = "primaryKey"
	AutoCreateTimeLabel = "autoCreateTime"
	AutoUpdateTimeLabel = "autoUpdateTime"
	AutoDeleteTimeLabel = "autoDeleteTime"
	SoftDeleteLabel     = "softDelete"
	DelFlagLabel        = "del_flag"
	CreateTimeLabel     = "create_time"
	UpdateTimeLabel     = "update_time"
	DeleteTimeLabel     = "delete_time"
	M2MLabel            = "many2many"
	ForeignKeyLabel      = "foreignKey"
	ReferencesLabel     = "References"
	JoinForeignKeyLabel = "joinForeignKey"
	JoinReferencesLabel = "joinReferences"
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
// 自动时间戳默认字段名为:`updated_at`、`created_at`, 表字段数据类型为: DATETIME
var AutoCreateTimeTag = gen.FieldGORMTag(CreateTimeLabel, func(tag field.GormTag) field.GormTag {
	tag.Append(AutoCreateTimeLabel, "true")
	return tag
})

var AutoUpdateTimeTag = gen.FieldGORMTag(UpdateTimeLabel, func(tag field.GormTag) field.GormTag {
	tag.Append(AutoUpdateTimeLabel, "true")
	return tag
})

// var AutoDeleteTimeTag = gen.FieldGORMTag(DeleteTime, func(tag field.GormTag) field.GormTag {
// 	tag.Append(AutoDeleteTime, "true")
// 	return tag
// })

var SoftDeleteFlagTag = gen.FieldGORMTag(DelFlagLabel, func(tag field.GormTag) field.GormTag {
	tag.Append(SoftDeleteLabel, "flag")
	return tag
})

// 软删除默认字段名为:`deleted_at`, 表字段数据类型为: DATETIME. 这里指定为`del_flag`,数据类型为tinyint
var SoftDeleteFieldType = gen.FieldType(DelFlagLabel, "soft_delete.DeletedAt")

//var CommonTag = []gen.ModelOpt{AutoCreateTimeTag, AutoUpdateTimeTag, AutoDeleteTimeTag}
