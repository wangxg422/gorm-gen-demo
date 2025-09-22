package dao

import (
	"context"
	"gorm-gen-demo/dal/model"
	"gorm-gen-demo/dal/query"
)

func GetUserWithDataScope(ctx context.Context) ([]*model.SysUser, error) {
	return query.SysUser.WithContext(ctx).
		Where(query.SysUser.Status.Eq(0)).
		Find()
}

func GetUserRolesWithDataScope(ctx context.Context, user *model.SysUser) ([]*model.SysUser, error) {
	return query.SysUser.WithContext(ctx).
		Where(query.SysUser.Status.Eq(0)).
		Find()
}

func GetUserInfo() {

}
