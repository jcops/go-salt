package menu

import (
	"github.com/astaxie/beego"
	"go-salt/common/e"
	"encoding/json"
	"net/http"
	"go-salt/controllers/menu_service"
	"fmt"
	"strconv"
	"github.com/Unknwon/com"
)

type Menu struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	Method string `json:"method"`
}

type MenuController struct {
	beego.Controller
}

func (this *MenuController) URLMapping() {
	this.Mapping("AddMenu",this.Addmenu)
	this.Mapping("Deletemenu",this.Deletemenu)
	this.Mapping("Updatemenu",this.Updatemenu)
	this.Mapping("Getmenu",this.Getmenu)
}

func (this *MenuController) ResponseJson(httpCode, errCode int, data interface{}) {
	this.Data["json"] = &e.Response{
		Code: httpCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	}
	this.ServeJSON()
	this.StopRun()
}


// @Summary   添加菜单
// @Tags 添加菜单
// @Description 添加菜单
// @Accept json
// @Produce  json
// @Failure 400 {string} json "{"code":400,  "data":null,"msg":"请求参数错误"}"
// @Failure 404 {string} json "{ "code": 404, "data":null,"msg":"请求参数错误"}"
// @Success 200 {string} json "{ "code": 200, "data": { "token": "xxx" }, "msg": "ok" }"
// @router /menu [post]
func (this *MenuController) Addmenu() {
	var menu Menu
	err := json.Unmarshal(this.Ctx.Input.RequestBody,&menu)
	if err != nil {
		this.ResponseJson(http.StatusBadRequest,e.INVALID_PARAMS,nil)
	}
	menuservice := menu_service.Menu{
		Name:menu.Name,
		Method:menu.Method,
		Path:menu.Path,
	}
	_ ,err = menuservice.Addmenu()
	if err != nil {
		fmt.Println(err)
		this.ResponseJson(http.StatusInternalServerError,e.ERROR_ADD_FAIL,nil)
	}
	this.ResponseJson(http.StatusOK,e.SUCCESS,nil)


}

// @Summary   删除菜单
// @Tags 删除菜单
// @Description 删除菜单
// @Accept json
// @Produce  json
// @Failure 400 {string} json "{"code":400,  "data":null,"msg":"请求参数错误"}"
// @Failure 404 {string} json "{ "code": 404, "data":null,"msg":"请求参数错误"}"
// @Success 200 {string} json "{ "code": 200, "data": { "token": "xxx" }, "msg": "ok" }"
// @router /menu/:id [delete]
func (this *MenuController) Deletemenu() {
	mid,err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil {
		this.ResponseJson(http.StatusBadRequest,e.INVALID_PARAMS,nil)
	}
	menuservice := menu_service.Menu{ID:mid}
	err = menuservice.DeleteMenu()
	if err != nil {
		fmt.Println(err)
		this.ResponseJson(http.StatusInternalServerError,e.ERROR_DELETE_FAIL,nil)
	}
	this.ResponseJson(http.StatusOK,e.SUCCESS,nil)

}

// @Summary   更新菜单
// @Tags 更新菜单
// @Description 更新菜单
// @Accept json
// @Produce  json
// @Failure 400 {string} json "{"code":400,  "data":null,"msg":"请求参数错误"}"
// @Failure 404 {string} json "{ "code": 404, "data":null,"msg":"请求参数错误"}"
// @Success 200 {string} json "{ "code": 200, "data": { "token": "xxx" }, "msg": "ok" }"
// @router /menu [put]
func (this *MenuController) Updatemenu() {
	var menu Menu
	if err := json.Unmarshal(this.Ctx.Input.RequestBody,&menu);err != nil {
		this.ResponseJson(http.StatusBadRequest,e.INVALID_PARAMS,nil)
	}
	menuservice := menu_service.Menu{
		ID:menu.Id,
		Name:menu.Name,
		Path:menu.Path,
		Method:menu.Method,
	}
	err := menuservice.UpdateMenu()
	if err != nil {
		fmt.Println(err)
		this.ResponseJson(http.StatusInternalServerError,e.ERROR_EDIT_FAIL,nil)
	}
	this.ResponseJson(http.StatusOK,e.SUCCESS,nil)
}
func GetPage(c *MenuController) int {
	result := 0
	page := com.StrTo(c.GetString("page")).MustInt()
	if page > 0 {
		result = (page - 1) * 10
	}
	fmt.Println(result)
	return result
}
// @Summary   查询菜单
// @Tags 查询菜单
// @Description 查询菜单
// @Accept json
// @Produce  json
// @Failure 400 {string} json "{"code":400,  "data":null,"msg":"请求参数错误"}"
// @Failure 404 {string} json "{ "code": 404, "data":null,"msg":"请求参数错误"}"
// @Success 200 {string} json "{ "code": 200, "data": { "token": "xxx" }, "msg": "ok" }"
// @router /menu [get]
func (this *MenuController) Getmenu() {

	menuservice := menu_service.Menu{
		PageSize:1000,
		PageNum:GetPage(this),
	}
	total, err := menuservice.Count()
	if err != nil {
		this.ResponseJson(http.StatusInternalServerError, e.ERROR_COUNT_FAIL, nil)
	}
	menu, err := menuservice.GetAll()
	if err != nil {
		this.ResponseJson(http.StatusInternalServerError, e.ERROR_GET_S_FAIL, nil)
	}
	res := make(map[string]interface{})
	res["lists"] = menu
	res["total"] = total
	this.ResponseJson(http.StatusOK, e.SUCCESS, res)

	}
