package request

type Policy struct {
	Path   string // API路径
	Method string // 请求方法
}

type UpdateRoleForm struct {
	Policys  []Policy `json:"policys" form:"policys" valid:"Required"`       // casbin 中 policy 规则
	RoleName string   `json:"role_name" form:"role_name" valid:"MinSize(2)"` // 角色名
}

type DeleteRoleForm struct {
	RoleName string `json:"role_name" form:"role_name" valid:"MinSize(2)"` // 角色名
}

type AddRoleForUserForm struct {
	UserName string `json:"user_name" form:"user_name" valid:"Match(/^[a-zA-Z0-9_-]{4,16}$/)"` // 用户名
	RoleName string `json:"role_name" form:"role_name" valid:"MinSize(2)"`                     // 角色名
}
