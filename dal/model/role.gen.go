// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

const TableNameRole = "role"

// Role mapped from table <role>
type Role struct {
	ID         int64                 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RoleCode   string                `gorm:"column:role_code;not null;comment:角色编码" json:"role_code"`                            // 角色编码
	RoleName   string                `gorm:"column:role_name;not null;comment:角色名称" json:"role_name"`                            // 角色名称
	DelFlag    soft_delete.DeletedAt `gorm:"column:del_flag;not null;default:0;comment:0可用1已删除;softDelete:flag" json:"del_flag"` // 0可用1已删除
	Status     int32                 `gorm:"column:status;not null;comment:0可用1停用" json:"status"`                                // 0可用1停用
	CreateTime *time.Time            `gorm:"column:create_time;autoCreateTime:true" json:"create_time"`
	UpdateTime *time.Time            `gorm:"column:update_time;autoUpdateTime:true" json:"update_time"`
	DeleteTime *time.Time            `gorm:"column:delete_time" json:"delete_time"`
	Remark     *string               `gorm:"column:remark" json:"remark"`
	Users      []User                `gorm:"joinForeignKey:RoleID;joinReferences:UserID;many2many:user_role" json:"users"`
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}
