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
	return query.User.WithContext(context.Background()).Where(query.User.UserID.Eq(user.UserID)).Updates(user);
}

func DeleteUser(userID string) (gen.ResultInfo, error) {
	return query.User.WithContext(context.Background()).Where(query.User.UserID.Eq(userID)).Delete();
}

func GetUserByUserID(userID string) (*model.User, error) {
	user, err := query.User.WithContext(context.Background()).Where(query.User.UserID.Eq(userID)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}
