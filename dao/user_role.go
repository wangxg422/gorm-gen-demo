package dao

import (
	"context"

	"gorm-gen-demo/dal/query"
)

type R map[any]any

func GetUserWithRoles(data interface{}) error {
	ur := query.UserRole.As("ur")

	return query.User.WithContext(context.Background()).
		LeftJoin(ur, query.User.ID.EqCol(query.Role.ID)).
		LeftJoin(ur, query.UserRole.RoleID.EqCol(query.Role.ID)).Scan(data)
}
