package bll

import (
	"go-salt/controllers/user_service"
	"go-salt/controllers/role_service"
)
type Common struct {
	UserAPI *user_service.User `inject:""`
	RoleAPI *role_service.Role `inject:""`
}
