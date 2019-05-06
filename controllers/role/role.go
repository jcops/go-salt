package role

import (
	"github.com/astaxie/beego"
	"fmt"
	"net/http"
	"github.com/Unknwon/com"
	"go-salt/common/e"
	"go-salt/controllers/role_service"
	"encoding/json"
	"go-salt/common/inject"
	"strconv"
)

type Role struct {
	Id int `json:"id"`
	Name string `json:"name"`
	User int `json:"user_id"`
	Menu int `json:"menu_id"`
}
type Roles struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Menu []int `json:"menu_id"`
}
type RoleController struct {
	beego.Controller
}
func (this *RoleController) URLMapping() {
	this.Mapping("GetAllRole",this.GetAllRole)
	this.Mapping("GetOneRole",this.GetOneRole)
	this.Mapping("AddlRole",this.AddlRole)
	this.Mapping("UpdateRole",this.UpdateRole)
	this.Mapping("DeleteRole",this.DeleteRole)
}
func GetPage(c *RoleController) int {
	result := 0
	page := com.StrTo(c.GetString("page")).MustInt()
	if page > 0 {
		result = (page - 1) * 10
	}
	fmt.Println(result)
	return result
}
type LoginResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Token interface{} `json:"token"`
	} `json:"data"`
}

func (this *RoleController) ResponseJson(httpCode, errCode int, data interface{}) {
	this.Data["json"] = &e.Response{
		Code: httpCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	}
	this.ServeJSON()
	this.StopRun()
}

// @Summary   查询所有角色
// @Tags 查询所有角色
// @Description 查询所有角色
// @Accept json
// @Produce  json
// @Failure 400 {string} json "{"code":400,  "data":null,"msg":"请求参数错误"}"
// @Failure 404 {string} json "{ "code": 404, "data":null,"msg":"请求参数错误"}"
// @Success 200 {string} json "{ "code": 200, "data": { "token": "xxx" }, "msg": "ok" }"
// @router /role [get]
func (this *RoleController) GetAllRole()  {
	roleService := role_service.Role{
		Name: this.GetString("name"),
		PageNum:  GetPage(this),
		PageSize: 10,
	}
	fmt.Println(GetPage(this))
	total, err := roleService.Count()
	if err != nil {
		this.ResponseJson(http.StatusInternalServerError, e.ERROR_COUNT_FAIL, nil)
	}
	role, err := roleService.GetAll()
	if err != nil {
		this.ResponseJson(http.StatusInternalServerError, e.ERROR_GET_S_FAIL, nil)
	}
	res := make(map[string]interface{})
	res["lists"] = role
	res["total"] = total
	this.ResponseJson(http.StatusOK, e.SUCCESS, res)

}

func (this *RoleController) GetOneRole()  {


}

// @Summary   添加角色
// @Tags 添加角色
// @Description 添加角色
// @Accept json
// @Produce  json
// @Failure 400 {string} json "{"code":400,  "data":null,"msg":"请求参数错误"}"
// @Failure 404 {string} json "{ "code": 404, "data":null,"msg":"请求参数错误"}"
// @Success 200 {string} json "{ "code": 200, "data": { "token": "xxx" }, "msg": "ok" }"
// @router /role [post]
func (this *RoleController) AddlRole()  {
	var ro Roles
	err := json.Unmarshal(this.Ctx.Input.RequestBody,&ro)
	if err!= nil {
		this.ResponseJson(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}
	fmt.Println(ro)
	roleservice := role_service.Role{Name:ro.Name}

	//// 检查角色是否存在
	//if err := roleservice.ExistsRoleByName(); err != nil {
	//	this.ResponseJson(http.StatusInternalServerError, e.ERROR_ADD_FAIL, err)
	//}
	role := make(map[string][]int)
	role["menu_id"] = ro.Menu

	id ,err := roleservice.AddRole(role)
	if err == nil{
		fmt.Println(id)
		fmt.Println(err)
		err = inject.Obj.Common.RoleAPI.LoadPolicy(id)
		if err != nil {
			this.ResponseJson(http.StatusInternalServerError, e.ERROR_EDIT_FAIL, nil)
		}
		this.ResponseJson(http.StatusOK, e.SUCCESS, nil)
	} else {
		fmt.Println(err)
		this.ResponseJson(http.StatusInternalServerError, e.ERROR_ADD_FAIL, err)
	}
}


// @Summary   修改角色信息
// @Tags 修改角色信息
// @Description 修改角色信息
// @Accept json
// @Produce  json
// @Failure 400 {string} json "{"code":400,  "data":null,"msg":"请求参数错误"}"
// @Failure 404 {string} json "{ "code": 404, "data":null,"msg":"请求参数错误"}"
// @Success 200 {string} json "{ "code": 200, "data": { "token": "xxx" }, "msg": "ok" }"
// @router /role [put]
func (this *RoleController)UpdateRole()  {
	var ro Roles
	err := json.Unmarshal(this.Ctx.Input.RequestBody,&ro)
	if err!= nil {
		this.ResponseJson(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}
	fmt.Println(ro)
	roleservice := role_service.Role{ID:ro.Id}
	role := make(map[string][]int)
	role["menu_id"] = ro.Menu
	id ,err := roleservice.UpdateRole(role)
	if err == nil{
		fmt.Println(id)
		fmt.Println(err)
		err = inject.Obj.Common.RoleAPI.LoadPolicy(id)
		if err != nil {
			this.ResponseJson(http.StatusInternalServerError, e.ERROR_EDIT_FAIL, nil)
		}
		this.ResponseJson(http.StatusOK, e.SUCCESS, nil)
	} else {
		fmt.Println(err)
		this.ResponseJson(http.StatusInternalServerError, e.ERROR_EDIT_FAIL, err)
	}

}

// @Summary   删除角色
// @Tags 删除角色
// @Description 删除角色
// @Accept json
// @Produce  json
// @Failure 400 {string} json "{"code":400,  "data":null,"msg":"请求参数错误"}"
// @Failure 404 {string} json "{ "code": 404, "data":null,"msg":"请求参数错误"}"
// @Success 200 {string} json "{ "code": 200, "data": { "token": "xxx" }, "msg": "ok" }"
// @router /role/:id [delete]
func (this *RoleController) DeleteRole()  {
	rid, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	roleservice := role_service.Role{ID:rid}

	exists ,err:= roleservice.ExistsRole()
	if err != nil {
		this.ResponseJson(http.StatusInternalServerError,e.ERROR_NOT_EXIST,nil)
	}
	if !exists {
		this.ResponseJson(http.StatusInternalServerError,e.ERROR_NOT_EXIST,nil)
	}
	if err := roleservice.DeleteRole();err != nil {
		this.ResponseJson(http.StatusInternalServerError,e.ERROR_DELETE_FAIL,nil)
	}
	this.ResponseJson(http.StatusOK,e.SUCCESS,nil)
}



