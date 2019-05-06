package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["go-salt/controllers/menu:MenuController"] = append(beego.GlobalControllerRouter["go-salt/controllers/menu:MenuController"],
        beego.ControllerComments{
            Method: "Addmenu",
            Router: `/menu`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-salt/controllers/menu:MenuController"] = append(beego.GlobalControllerRouter["go-salt/controllers/menu:MenuController"],
        beego.ControllerComments{
            Method: "Updatemenu",
            Router: `/menu`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-salt/controllers/menu:MenuController"] = append(beego.GlobalControllerRouter["go-salt/controllers/menu:MenuController"],
        beego.ControllerComments{
            Method: "Getmenu",
            Router: `/menu`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-salt/controllers/menu:MenuController"] = append(beego.GlobalControllerRouter["go-salt/controllers/menu:MenuController"],
        beego.ControllerComments{
            Method: "Deletemenu",
            Router: `/menu/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
