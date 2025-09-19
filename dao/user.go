package dao

import (
	"context"
	"gorm-gen-demo/dal/model"
	"gorm-gen-demo/dal/query"

	"gorm.io/gen"
)

func CreateUser(ctx context.Context, user *model.SysUser) error {
	return query.SysUser.WithContext(ctx).Create(user);
}

func UpdateUser(ctx context.Context, user *model.SysUser) (gen.ResultInfo, error) {
	return query.SysUser.WithContext(ctx).Where(query.SysUser.ID.Eq(user.ID)).Updates(user);
}

func DeleteUser(ctx context.Context, id int64) (gen.ResultInfo, error) {
	return query.SysUser.WithContext(ctx).Where(query.SysUser.ID.Eq(id)).Delete();
}

func GetUserByUserID(ctx context.Context, id int64) (*model.SysUser, error) {
	user, err := query.SysUser.WithContext(ctx).Where(query.SysUser.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetRolesByUserId(ctx context.Context, id int64) ([]model.SysRole, error) {
	user, err := query.SysUser.WithContext(ctx).Preload(query.SysUser.SysRoleList).Where(query.SysUser.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return user.SysRoleList, nil
}
