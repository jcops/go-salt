package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["go-salt/controllers/role:RoleController"] = append(beego.GlobalControllerRouter["go-salt/controllers/role:RoleController"],
        beego.ControllerComments{
            Method: "GetAllRole",
            Router: `/role`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-salt/controllers/role:RoleController"] = append(beego.GlobalControllerRouter["go-salt/controllers/role:RoleController"],
        beego.ControllerComments{
            Method: "AddlRole",
            Router: `/role`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-salt/controllers/role:RoleController"] = append(beego.GlobalControllerRouter["go-salt/controllers/role:RoleController"],
        beego.ControllerComments{
            Method: "UpdateRole",
            Router: `/role`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-salt/controllers/role:RoleController"] = append(beego.GlobalControllerRouter["go-salt/controllers/role:RoleController"],
        beego.ControllerComments{
            Method: "DeleteRole",
            Router: `/role/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
