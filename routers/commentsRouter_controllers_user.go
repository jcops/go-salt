package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["go-salt/controllers/user:UserController"] = append(beego.GlobalControllerRouter["go-salt/controllers/user:UserController"],
        beego.ControllerComments{
            Method: "GetInfo",
            Router: `/getinfo`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-salt/controllers/user:UserController"] = append(beego.GlobalControllerRouter["go-salt/controllers/user:UserController"],
        beego.ControllerComments{
            Method: "UserRegister",
            Router: `/register`,
            AllowHTTPMethods: []string{"POST"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-salt/controllers/user:UserController"] = append(beego.GlobalControllerRouter["go-salt/controllers/user:UserController"],
        beego.ControllerComments{
            Method: "GetAllUser",
            Router: `/users`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-salt/controllers/user:UserController"] = append(beego.GlobalControllerRouter["go-salt/controllers/user:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: `/users/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-salt/controllers/user:UserController"] = append(beego.GlobalControllerRouter["go-salt/controllers/user:UserController"],
        beego.ControllerComments{
            Method: "UpdateUserPwd",
            Router: `/users/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-salt/controllers/user:UserController"] = append(beego.GlobalControllerRouter["go-salt/controllers/user:UserController"],
        beego.ControllerComments{
            Method: "GetOneUser",
            Router: `/users/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-salt/controllers/user:UserController"] = append(beego.GlobalControllerRouter["go-salt/controllers/user:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: `/users/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
