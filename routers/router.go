// @title Golang Beego API
// @version 1.0
// @description An example of Beego
// @termsOfService https://github.com/jcops/go-salt

// @license.name MIT
// @license.url  https://github.com/jcops/go-salt
package routers

import (
	"github.com/astaxie/beego"
	"go-salt/controllers/salt"
	"go-salt/controllers/user"
	"go-salt/controllers/role"
	"go-salt/controllers/menu"
)

func init() {
	//beego.Router("/user/info", &user.UserController{}, "get:GetInfo")
	beego.Router("/user/login", &user.UserController{}, "post:Login")
	//beego.Router("/user/register", &user.UserController{}, "post:UserRegister")

	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSInclude(
				&salt.SaltController{},
			),
		),
		beego.NSNamespace("/v1",
			beego.NSInclude(
				&user.UserController{},
			),
		),
		beego.NSNamespace("/v1",
			beego.NSInclude(
				&role.RoleController{},
			),
		),
		beego.NSNamespace("/v1",
			beego.NSInclude(
				&menu.MenuController{},
			),
		),
	)
	beego.AddNamespace(ns)
	//ns1 := beego.NewNamespace("/user",
	//	beego.NSInclude(
	//		&user.UserController{},
	//	),
	//)
	//beego.AddNamespace(ns1)

}
