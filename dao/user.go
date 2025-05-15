package dao

import (
	"context"
	"gorm-gen-demo/dal/model"
	"gorm-gen-demo/dal/query"

	"gorm.io/gen"
)

func CreateUser(user *model.User) error {
	return query.User.WithContext(context.Background()).Create(user);
}

func UpdateUser(user *model.User) (gen.ResultInfo, error) {
	return query.User.WithContext(context.Background()).Where(query.User.ID.Eq(user.ID)).Updates(user);
}

func DeleteUser(id int64) (gen.ResultInfo, error) {
	return query.User.WithContext(context.Background()).Where(query.User.ID.Eq(id)).Delete();
}

func GetUserByUserID(id int64) (*model.User, error) {
	user, err := query.User.WithContext(context.Background()).Where(query.User.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}
