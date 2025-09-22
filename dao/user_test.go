package dao

import (
	"context"
	"fmt"

	"gorm-gen-demo/dal/model"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func genUsername() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10)
}

func TestUserCreate(t *testing.T) {
	username := genUsername()

	user := model.SysUser{
		UserName: username,
		Password: "123456",
		Status:   0,
	}

	err := CreateUser(context.Background(), &user)
	assert.NoError(t, err)
	assert.Greater(t, user.ID, int64(0))
}

func TestUserRead(t *testing.T) {
	username := genUsername()

	user := model.SysUser{
		UserName: username,
		Password: "123456",
		Status:   0,
	}

	err := CreateUser(context.Background(), &user)
	assert.NoError(t, err)

	FindUser, err := GetUserByUserID(context.Background(), user.ID)
	assert.NoError(t, err)
	assert.Equal(t, username, FindUser.UserName)
}

func TestUserUpdate(t *testing.T) {
	username := genUsername()

	user := model.SysUser{
		UserName: username,
		Password: "123456",
		Status:   0,
	}

	err := CreateUser(context.Background(), &user)
	assert.NoError(t, err)

	newUser := model.SysUser{
		ID:       user.ID,
		UserName: "new_" + username,
		Password: "123456",
		Status:   0,
	}

	r, err := UpdateUser(context.Background(), &newUser)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), r.RowsAffected)
	assert.NoError(t, r.Error)
}

func TestUserDelete(t *testing.T) {
	username := genUsername()

	user := model.SysUser{
		UserName: username,
		Password: "123456",
		Status:   0,
	}

	err := CreateUser(context.Background(), &user)
	assert.NoError(t, err)

	r, err := DeleteUser(context.Background(), user.ID)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), r.RowsAffected)
	assert.NoError(t, r.Error)

	GetUser, err := GetUserByUserID(context.Background(), user.ID)
	fmt.Println(err)
	assert.Error(t, err)
	assert.Nil(t, GetUser)
}

// func TestGetRolesByUserId(t *testing.T) {
// 	GetRolesByUserId(context.Background(), 1)
// 	//assert.NotErrorIs().Error(t, err,)
// }
