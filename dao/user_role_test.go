package dao

import (
	"gorm-gen-demo/dal"
	"gorm-gen-demo/dal/query"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserWithRoles(t *testing.T) {
	query.SetDefault(dal.ConnectDB().Debug())

	var data = make([]map[string]any, 0)
	err := GetUserWithRoles(&data)
	assert.NoError(t, err)
}