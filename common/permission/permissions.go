package permission

import (
	"strings"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"go-salt/common/inject"
	"net/http"
	"github.com/astaxie/beego/context"
	Jwt "go-salt/common/jwt"
)

var CasbinMiddleware =func(ctx *context.Context) {
	Authorization := ctx.Input.Header("Authorization")
	token := strings.Split(Authorization, " ")
	t, _ := jwt.Parse(token[1], func(*jwt.Token) (interface{}, error) {
		return Jwt.JwtSecret, nil
	})
	fmt.Println("req path",Jwt.GetIdFromClaims("username", t.Claims), ctx.Request.URL.Path, ctx.Request.Method)

	if b, err := inject.Obj.Enforcer.EnforceSafe(Jwt.GetIdFromClaims("username", t.Claims), ctx.Request.URL.Path, ctx.Request.Method); err != nil {
		fmt.Println("err",err)
		fmt.Println(b)
		Jwt.ResponseWithJson(ctx.ResponseWriter,http.StatusUnauthorized,
			Jwt.Response{Code: http.StatusOK, Msg: "ok",Data:err})
		return

	} else if !b {
		fmt.Println("err",err)
		fmt.Println(b)
		fmt.Println(ctx.ResponseWriter,http.StatusUnauthorized,http.StatusForbidden, "无权限",err)
		Jwt.ResponseWithJson(ctx.ResponseWriter,http.StatusUnauthorized,
			Jwt.Response{Code: http.StatusForbidden, Msg: "无权限",Data:err})
		return
	}

}
