package dao

import (
	"context"
	"fmt"
	"gorm-gen-demo/client"
	"gorm-gen-demo/dal/model"
	"gorm-gen-demo/dal/query"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var zore int64 = 0
var one int64 = 1

func genUsername() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10)
}

func TestUserCreate(t *testing.T) {
	query.SetDefault(client.ConnectDB().Debug())

	username := genUsername()

	user := model.User{
		UserName: username,
		Password: "123456",
		Status:   0,
	}

	err := CreateUser(context.Background(), &user)
	assert.NoError(t, err)
	assert.Greater(t, user.ID, zore)
}

func TestUserRead(t *testing.T) {
	query.SetDefault(client.ConnectDB().Debug())

	username := genUsername()

	user := model.User{
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
	query.SetDefault(client.ConnectDB().Debug())

	username := genUsername()

	user := model.User{
		UserName: username,
		Password: "123456",
		Status:   0,
	}

	err := CreateUser(context.Background(), &user)
	assert.NoError(t, err)

	newUser := model.User{
		ID:       user.ID,
		UserName: "new_" + username,
		Password: "123456",
		Status:   0,
	}

	r, err := UpdateUser(context.Background(), &newUser)
	assert.NoError(t, err)
	assert.Equal(t, one, r.RowsAffected)
	assert.NoError(t, r.Error)
}

func TestUserDelete(t *testing.T) {
	query.SetDefault(client.ConnectDB().Debug())

	username := genUsername()

	user := model.User{
		UserName: username,
		Password: "123456",
		Status:   0,
	}

	err := CreateUser(context.Background(), &user)
	assert.NoError(t, err)

	r, err := DeleteUser(context.Background(), user.ID)
	assert.NoError(t, err)
	assert.Equal(t, one, r.RowsAffected)
	assert.NoError(t, r.Error)

	GetUser, err := GetUserByUserID(context.Background(), user.ID)
	fmt.Println(err)
	assert.Error(t, err)
	assert.Nil(t, GetUser)
}

func TestGetRolesByUserId(t *testing.T) {
	query.SetDefault(client.ConnectDB().Debug())

	GetRolesByUserId(context.Background(), 1)
	//assert.NotErrorIs().Error(t, err,)
}

