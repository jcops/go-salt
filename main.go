package main

import (
	_ "go-salt/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"go-salt/common/jwt"
	"go-salt/models"
	"go-salt/common/permission"
	"go-salt/common/inject"
)


func init() {
	// 初始化数据库链接
	models.Setup()
	//初始化inject
	inject.Setup()
}
func main() {
	// 加载casbin策略数据，包括角色权限数据、用户角色数据
	err := inject.LoadCasbinPolicyData()
	if err != nil {
		panic("加载casbin策略数据发生错误: " + err.Error())
	}
	//beego.InsertFilter("/v1/salt/*", beego.BeforeRouter, func(ctx *context.Context) {
	//	cookie, err := ctx.Request.Cookie("Authorization")
	//	if err != nil || !user.CheckToken(cookie.Value) {
	//		http.Redirect(ctx.ResponseWriter, ctx.Request, "/v1/user/login", http.StatusMovedPermanently)
	//
	//	}
	//})



	//用户信息需要 jwt 认证
	beego.InsertFilter("/api/v1/*", beego.BeforeExec, jwt.JWT)
	// casbin 权限验证
	beego.InsertFilter("/api/v1/*", beego.BeforeExec, permission.CasbinMiddleware)
	// 跨域
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//AllowAllOrigins:  true,
		AllowOrigins: []string{"http://*.*.*.*:*","http://localhost:*","http://127.0.0.1:*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,

	}))

	if beego.BConfig.RunMode == "dev" || beego.BConfig.RunMode == "prod" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
