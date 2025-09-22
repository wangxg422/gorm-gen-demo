package dao

import (
	"context"
	"gorm-gen-demo/plugin"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryDataScope(t *testing.T) {
	// 构造数据范围
	cfg := plugin.ScopeConfig{
		UserID:  1001,
		DeptID:  2001,
		DeptIDs: []int64{2001, 2002, 2003},
		Scope:   plugin.ScopeAll,
		TableAlias: map[string]string{
			//"sys_user": "u",
		},
	}

	ctx := plugin.WithScopeConfig(context.Background(), cfg)
	_, err := GetUserWithDataScope(ctx)
	assert.NoError(t, err)
	//assert.Equal(t, 1, len(users))
}
