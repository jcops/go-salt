package user

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"go-salt/common/jwt"
	"fmt"
	"go-salt/common/e"
	"net/http"
	"go-salt/controllers/user_service"
	"go-salt/common/inject"
	"github.com/Unknwon/com"
	"strconv"
	"go-salt/common"
)

// 用户
type UserController struct {
	beego.Controller
}

func (this *UserController) URLMapping() {
	this.Mapping("Login",this.Login)
	this.Mapping("UserRegister",this.UserRegister)
	this.Mapping("UpdateUser",this.UpdateUser)
	this.Mapping("UpdateUserPwd",this.UpdateUserPwd)
	this.Mapping("DeleteUser",this.DeleteUser)
	this.Mapping("GetAllUser",this.GetAllUser)
	this.Mapping("GetOneUser",this.GetOneUser)

}


type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role_id"`
}
type Users struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Role []int `json:"role_id"`
}

type LoginResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Token interface{} `json:"token"`
	} `json:"data"`
}

func (this *UserController) ResponseJson(httpCode, errCode int, data interface{}) {
	this.Data["json"] = &e.Response{
		Code: httpCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	}
	this.ServeJSON()
	this.StopRun()
}

// @Title 获取用户token
// @Description 获取用户token
// @Param	body		body 	User	true		"body"
// @Success 200 {string}
// @Failure 404 body is empty
func (this *UserController) Login() {
	result := LoginResult{}
	var user User
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err != nil {
		fmt.Println(err)
		this.ResponseJson(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}
	fmt.Println("req body", user)
	//isExist,err := models.CheckUser(user.Username,user.Password)
	authService := user_service.User{Username: user.Username, Password: user.Password}
	isExist, err := authService.Check()
	if err != nil {
		fmt.Println(err)
		this.ResponseJson(http.StatusInternalServerError, e.ERROR_USERORPASS_FAIL, nil)
	}
	if isExist == nil {
		this.ResponseJson(http.StatusUnauthorized, e.ERROR_NOT_EXIST, nil)
	}
	//token := GenToken()
	//this.Ctx.SetCookie("Authorization", token,  "/", 3600)


	token, err := jwt.GenerateToken(user.Username, user.Password)
	if err != nil {
		result.Code = e.ERROR
		result.Msg = e.GetMsg(e.ERROR_AUTH_TOKEN)
	} else {
		result.Code = e.SUCCESS
		result.Msg = e.GetMsg(e.SUCCESS)
		result.Data.Token = token
		this.Ctx.SetCookie("Authorization", token, "/", 3600)
	}
	this.SetSession("id",isExist.ID)
	this.SetSession("name",isExist.Username)
	this.Data["json"] = result
	this.ServeJSON()
}

// @Title 获取用户信息
// @Description 获取用户信息
// @Success 200 {string}
// @Failure 404 body is empty
// @router /getinfo [get]
func (this *UserController) GetInfo() {
		maps := make(map[string]interface{})
		maps["id"] = this.GetSession("id")
		maps["name"] = this.GetSession("name")
		fmt.Println(maps)
		this.Data["json"] = maps
		this.ServeJSON()
}

// @Summary   注册用户
// @Tags 注册用户
// @Description 注册用户
// @Accept json
// @Produce  json
// @Param  username  body   string true "username"
// @Param  password  body   string true "password"
// @Param  role_id  body   int false "role_id"
// @Success 200 {string} json "{ "code": 200, "data": { "token": "xxx" }, "msg": "ok" }"
// @Failure 400 {string} json "{"code":400,  "data":null,"msg":"请求参数错误"}"
// @Failure 404 {string} json "{ "code": 404, "data":null,"msg":"请求参数错误"}"
// @router /register [POST]
func (this *UserController) UserRegister() {
	var user User
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err != nil {
		fmt.Println(err)
		this.ResponseJson(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}
	fmt.Println("request body", user)
	userService := user_service.User{
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	}
	if id, err := userService.AddUser(); err == nil {
		err = inject.Obj.Common.UserAPI.LoadPolicy(id)
		if err != nil {
			this.ResponseJson(http.StatusInternalServerError, e.ERROR_EDIT_FAIL, nil)
		}
		this.ResponseJson(http.StatusOK, e.SUCCESS, nil)
	} else {
		this.ResponseJson(http.StatusBadRequest, e.ERROR_USERALREADYEXISTS_FAIL, nil)
	}

}



func GetPage(c *UserController) int {
	result := 0
	page := com.StrTo(c.GetString("page")).MustInt()
	if page > 0 {
		result = (page - 1) * 10
	}
	fmt.Println(result)
	return result
}



// @Summary   查询所有用户
// @Tags 查询所有用户
// @Description 查询所有用户
// @Accept json
// @Produce  json
// @Failure 400 {string} json "{"code":400,  "data":null,"msg":"请求参数错误"}"
// @Failure 404 {string} json "{ "code": 404, "data":null,"msg":"请求参数错误"}"
// @Success 200 {string} json "{ "code": 200, "data": { "token": "xxx" }, "msg": "ok" }"
// @router /users [get]
func (this *UserController) GetAllUser() {
	userService := user_service.User{
		Username: this.GetString("username"),
		PageNum:  GetPage(this),
		PageSize: 10,
	}
	fmt.Println(GetPage(this))
	total, err := userService.Count()
	if err != nil {
		this.ResponseJson(http.StatusInternalServerError, e.ERROR_COUNT_FAIL, nil)
	}
	user, err := userService.GetAll()
	if err != nil {
		this.ResponseJson(http.StatusInternalServerError, e.ERROR_GET_S_FAIL, nil)
	}
	res := make(map[string]interface{})
	res["lists"] = user
	res["total"] = total
	this.ResponseJson(http.StatusOK, e.SUCCESS, res)
}

// @Summary   查询单个用户
// @Tags 查询单个用户
// @Description 查询单个用户
// @Accept json
// @Produce  json
// @Failure 400 {string} json "{"code":400,  "data":null,"msg":"请求参数错误"}"
// @Failure 404 {string} json "{ "code": 404, "data":null,"msg":"请求参数错误"}"
// @Success 200 {string} json "{ "code": 200, "data": { "token": "xxx" }, "msg": "ok" }"
// @router /users/:id [get]
func (this *UserController) GetOneUser() {
	uid, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	userservice := user_service.User{ID: uid}
	exists, err := userservice.ExistById()
	if err != nil {
		this.ResponseJson(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
	}
	if !exists {
		this.ResponseJson(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
	}
	user, err := userservice.GetOneUser()
	if err != nil {
		this.ResponseJson(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
	}
	this.ResponseJson(http.StatusOK, e.SUCCESS, user)
}

// @Summary   删除用户
// @Tags 删除用户
// @Description 删除用户
// @Accept json
// @Produce  json
// @Param  id  path  int true "id"
// @Failure 400 {string} json "{"code":400,  "data":null,"msg":"请求参数错误"}"
// @Failure 404 {string} json "{ "code": 404, "data":null,"msg":"请求参数错误"}"
// @Success 200 {string} json "{ "code": 200, "data": { "token": "xxx" }, "msg": "ok" }"
// @router /users/:id [delete]
func (this *UserController) DeleteUser() {
	uid, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	userservice := user_service.User{ID: uid}
	exists, err := userservice.ExistById()
	if err != nil {
		this.ResponseJson(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
	}
	if exists {
		err := userservice.DeleteUser()
		if err != nil {
			this.ResponseJson(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
		}
		this.ResponseJson(http.StatusOK, e.SUCCESS, nil)

	}

}

// @Summary   修改用户信息
// @Tags 修改用户信息
// @Description 修改用户信息
// @Accept json
// @Produce  json
// @Param  id  body  int true "id"
// @Param  role_id  body  int false "role_id"
// @Failure 400 {string} json "{"code":400,  "data":null,"msg":"请求参数错误"}"
// @Failure 404 {string} json "{ "code": 404, "data":null,"msg":"请求参数错误"}"
// @Success 200 {string} json "{ "code": 200, "data": { "token": "xxx" }, "msg": "ok" }"
// @router /users/ [put]
func (this *UserController) UpdateUser() {
	var reqinfo Users
	fmt.Println("req",string(this.Ctx.Input.RequestBody))
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &reqinfo)
	if err != nil {
		fmt.Println(err)
		this.ResponseJson(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}
	fmt.Println("UpdateUser",reqinfo)
	userservice := user_service.User{
		ID:   reqinfo.Id,
	}

	exists, err := userservice.ExistById()
	if err != nil {
		this.ResponseJson(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
	}
	if !exists {
		this.ResponseJson(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
	}
	role_id := make(map[string][]int)
	role_id["role_id"] = reqinfo.Role
	err = userservice.EditUser(role_id)
	if err != nil {
		this.ResponseJson(http.StatusInternalServerError, e.ERROR_EDIT_FAIL, nil)
	}
	this.ResponseJson(http.StatusOK, e.SUCCESS, nil)
}

// @Summary   修改用户密码
// @Tags 修改用户密码
// @Description 修改用户密码
// @Accept json
// @Produce  json
// @Param	body		body 	User	true		"body content"
// @Success 200 {string} json "{ "code": 200, "data": { "token": "xxx" }, "msg": "ok" }"
// @router /users/ [post]
func (this *UserController) UpdateUserPwd() {
	// old 和conf密码在前端验证。
	var user User
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err != nil {
		fmt.Println(err)
		this.ResponseJson(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}
	userservice := user_service.User{
		ID:       user.Id,
		Password: common.EncodeMD5(user.Password),
	}

	exists, err := userservice.ExistById()
	if err != nil {
		this.ResponseJson(http.StatusInternalServerError, e.ERROR, nil)
	}
	if !exists {
		this.ResponseJson(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
	}
	err = userservice.UpdateUswePassword()
	if err != nil {
		this.ResponseJson(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
	}
	this.ResponseJson(http.StatusOK, e.SUCCESS, nil)

}
