package dao

import (
	"context"

	"gorm-gen-demo/dal/query"
)

type R map[any]any

func FindRolesOfUser(data interface{}) error {
	return query.User.WithContext(context.Background()).
		LeftJoin(query.UserRole, query.User.ID.EqCol(query.Role.ID)).
		LeftJoin(query.UserRole, query.UserRole.RoleID.EqCol(query.Role.ID)).Scan(data)
}
