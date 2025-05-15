package dao

import (
	"gorm-gen-demo/client"
	"gorm-gen-demo/dal/model"
	"gorm-gen-demo/dal/query"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserCRUD(t *testing.T) {
	query.SetDefault(client.ConnectDB().Debug())

	username := "u" + strconv.FormatInt(time.Now().UnixMilli(), 13)
	user := model.User {
		Username: username,
		Password: "123456",
		Status: 1,
		Remark: "test",
	}

	err := CreateUser(&user)
	assert.NoError(t, err)
	assert.Equal(t, username, user)
	assert.Greater(t, user.ID, 0)

	FindUser, err := GetUserByUserID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, username, FindUser.Username)

	newUser := model.User {
		Username: "new_" + username,
	}

	var one int64 = 1

	r, err := UpdateUser(&newUser)
	assert.NoError(t, err)
	assert.Equal(t, one, r.RowsAffected)
	assert.NoError(t, r.Error)

	r, err = DeleteUser(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, one, r.RowsAffected)
	assert.NoError(t, r.Error)
}