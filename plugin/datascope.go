package plugin

type DataScopeType int

const (
	ScopeAll             DataScopeType = iota // 全部，不限制
	ScopeSelf                                 // 仅本人
	ScopeDept                                 // 本部门
	ScopeDeptAndChildren                      // 本部门及以下
	ScopeCustom                               // 自定义
)

type ScopeConfig struct {
	UserID  int64
	DeptID  int64
	DeptIDs []int64
	Scope   DataScopeType
	// 表别名配置，比如 "sys_user" -> "u"
	TableAlias map[string]string
}
