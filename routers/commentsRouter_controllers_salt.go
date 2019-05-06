package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["go-salt/controllers/salt:SaltController"] = append(beego.GlobalControllerRouter["go-salt/controllers/salt:SaltController"],
        beego.ControllerComments{
            Method: "KeyAccept",
            Router: `/keyaccept`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-salt/controllers/salt:SaltController"] = append(beego.GlobalControllerRouter["go-salt/controllers/salt:SaltController"],
        beego.ControllerComments{
            Method: "KeyCopyFile",
            Router: `/keycopy`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-salt/controllers/salt:SaltController"] = append(beego.GlobalControllerRouter["go-salt/controllers/salt:SaltController"],
        beego.ControllerComments{
            Method: "KeyDelete",
            Router: `/keydelete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-salt/controllers/salt:SaltController"] = append(beego.GlobalControllerRouter["go-salt/controllers/salt:SaltController"],
        beego.ControllerComments{
            Method: "KeyDeploy",
            Router: `/keydeploy`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-salt/controllers/salt:SaltController"] = append(beego.GlobalControllerRouter["go-salt/controllers/salt:SaltController"],
        beego.ControllerComments{
            Method: "GetKeyList",
            Router: `/keylist`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-salt/controllers/salt:SaltController"] = append(beego.GlobalControllerRouter["go-salt/controllers/salt:SaltController"],
        beego.ControllerComments{
            Method: "KeyPing",
            Router: `/keyping`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-salt/controllers/salt:SaltController"] = append(beego.GlobalControllerRouter["go-salt/controllers/salt:SaltController"],
        beego.ControllerComments{
            Method: "KeyRun",
            Router: `/keyrun`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
