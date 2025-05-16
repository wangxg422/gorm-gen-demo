package dao

import (
	"context"
	"gorm-gen-demo/dal/model"
	"gorm-gen-demo/dal/query"

	"gorm.io/gen"
)

func CreateUser(ctx context.Context, user *model.User) error {
	return query.User.WithContext(ctx).Create(user);
}

func UpdateUser(ctx context.Context, user *model.User) (gen.ResultInfo, error) {
	return query.User.WithContext(ctx).Where(query.User.ID.Eq(user.ID)).Updates(user);
}

func DeleteUser(ctx context.Context, id int64) (gen.ResultInfo, error) {
	return query.User.WithContext(ctx).Where(query.User.ID.Eq(id)).Delete();
}

func GetUserByUserID(ctx context.Context, id int64) (*model.User, error) {
	user, err := query.User.WithContext(ctx).Where(query.User.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}
