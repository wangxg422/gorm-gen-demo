package dao

import (
	"fmt"
	"gorm-gen-demo/dal"
	"gorm-gen-demo/dal/model"
	"gorm-gen-demo/dal/query"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserCRUD(t *testing.T) {
	query.SetDefault(dal.ConnectDB().Debug())

	userId := "u" + strconv.FormatInt(time.Now().UnixMilli(), 13)
	user := model.User {
		UserID: userId,
		Username: "test",
		Password: "123456",
		Status: 1,
		Remark: "test",
	}

	err := CreateUser(&user)
	assert.NoError(t, err)
	assert.Equal(t, userId, user.UserID)

	FindUser, err := GetUserByUserID(user.UserID)
	assert.NoError(t, err)
	assert.Equal(t, user.UserID, FindUser.UserID)

	newUser := model.User {
		UserID: userId,
		Username: "new_username",
	}

	var one int64 = 1

	r, err := UpdateUser(&newUser)
	assert.NoError(t, err)
	assert.Equal(t,"new_username", newUser.Username)
	fmt.Println(r.RowsAffected)
	assert.Equal(t, one, r.RowsAffected)
	assert.NoError(t, r.Error)

	r, err = DeleteUser(user.UserID)
	assert.NoError(t, err)
	assert.Equal(t, one, r.RowsAffected)
	assert.NoError(t, r.Error)
}