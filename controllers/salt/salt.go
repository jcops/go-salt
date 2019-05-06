package salt

import (
	"github.com/astaxie/beego"
	"net/url"
	"fmt"
	"net/http"
	"strings"
	"crypto/tls"
	"io/ioutil"
	"github.com/bitly/go-simplejson"
	"k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/json"
	"go-salt/common/e"
)

// salt
type SaltController struct {
	beego.Controller
}
type Salt struct {
	Match string
	Cmd string
	Arg string
	Run string
	Tgt string
	Arg1 string
	Arg2 string
}
type LoginResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Token interface{} `json:"token"`
	} `json:"data"`
}

func (this *SaltController) ResponseJson(httpCode, errCode int, data interface{}) {
	this.Data["json"] = &e.Response{
		Code: httpCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	}
	this.ServeJSON()
	this.StopRun()
}
type jj struct {
	Return []map[string]string
}
func exec(run string, accept string, ctype string) (body []byte,err error){
	token := GetToken()
	saltapi := beego.AppConfig.String("Saltapi")
	req ,err := http.NewRequest("POST", saltapi,strings.NewReader(run))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Accept",accept)
	req.Header.Set("X-Auth-Token",token)
	req.Header.Set("Content-Type",ctype)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify:true},
		DisableCompression:true,
	}
	client := &http.Client{Transport: tr}
	resp,err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

func (this *SaltController) URLMapping() {
	this.Mapping("GetKeyList",this.GetKeyList)
	this.Mapping("KeyPing",this.KeyPing)
	this.Mapping("KeyAccept",this.KeyAccept)
	this.Mapping("KeyDelete",this.KeyDelete)
	this.Mapping("KeyRun",this.KeyRun)
	this.Mapping("KeyDeploy",this.KeyDeploy)
	this.Mapping("KeyCopyFile",this.KeyCopyFile)
}



// @Title 获取salt节点
// @Description 获取salt节点
// @Success 200 {string}
// @Failure 404 body is empty
// @router /keylist [get]
func (this *SaltController) GetKeyList() {
	parse := setparse("client","wheel","tgt","*","key.list_all")
	body ,_:= exec(parse, "application/json","application/x-www-form-urlencoded")
	fmt.Println("response Body:", string(body))
	jsons,err := simplejson.NewJson(body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsons)
	minions := jsons.Get("return").GetIndex(0).Get("data").Get("return").Get("minions").MustStringArray()
	minions_pre := jsons.Get("return").GetIndex(0).Get("data").Get("return").Get("minions_pre").MustStringArray()

	var min_str []string

	for i :=0; i< len(minions); i++ {
		min_str = append(min_str, minions[i])
	}
	out := strings.Join(min_str,",")
	fmt.Println(out)
	fmt.Println(minions,minions_pre)
	var pre = make(map[string]interface{})
	pre["minions"] = minions
	pre["minions_pre"] = minions_pre
	pre["msg"] = "ok"
	this.ResponseJson(http.StatusOK, e.SUCCESS, map[string]interface{}{"data": pre})

}


// @Title 获取salt节点连通性
// @Description 获取salt节点连通性
// @Success 200 {string}
// @Failure 404 body is empty
// @router /keyping [get]
func (this *SaltController) KeyPing() {
	parse := setparse("tgt","*","client","local","test.ping")
	body ,_:= exec(parse,"application/json","application/x-www-form-urlencoded")
	fmt.Println("resp",string(body))

	jsons,err := simplejson.NewJson(body)
	if err != nil {
		fmt.Println("err",err)
	}
	//y,_ :=yaml.YAMLToJSON(body)
	pre := make(map[string]interface{})
	pre["msg"] = "ok"
	pre["data"] = jsons
	this.ResponseJson(http.StatusOK, e.SUCCESS, pre)
}

// @Title 认证获取salt节点
// @Description 认证获取salt节点
// @Param	match		query 	string	true		"match"
// @Success 200 {string}
// @Failure 404 body is empty
// @router /keyaccept [post]
func (this *SaltController) KeyAccept() {
	//fmt.Println(this.Ctx.Request.Body)
	//match := this.GetString("match")
	//fmt.Println(match)
	var salt Salt
	err := json.Unmarshal(this.Ctx.Input.RequestBody,&salt)
	if err != nil {
		fmt.Println("参数错误")
	}
	parse :=setparse("client","wheel","match",salt.Match,"key.accept")
	fmt.Println(parse)
	body ,_:= exec(parse,"application/json","application/x-www-form-urlencoded")
	fmt.Println(string(body))
	pre := make(map[string]interface{})
	pre["msg"] = "ok"
	this.ResponseJson(http.StatusOK, e.SUCCESS, pre)
}

// @Title 认证获取salt节点
// @Description 认证获取salt节点
// @Param	match		query	string	false		"match"
// @Success 200 {string}
// @Failure 404 body is empty
// @router /keydelete [post]
func (this *SaltController) KeyDelete() {
	var salt Salt

	err := json.Unmarshal(this.Ctx.Input.RequestBody,&salt)
	if err != nil {
		fmt.Println("参数错误")
	}
	parse := setparse("client","wheel","match",salt.Match,"key.delete")
	fmt.Println(parse)

	body ,_:= exec(parse,"application/json","application/x-www-form-urlencoded")
	fmt.Println(string(body))
	pre := make(map[string]interface{})
	pre["msg"] = "ok"
	this.ResponseJson(http.StatusOK, e.SUCCESS, pre)

}

// @Title salt执行命令
// @Description salt执行命令
// @Param	match		query	string	false		"match"
// @Param	cmd		query	string	false		"cmd"
// @Success 200 {string}
// @Failure 404 body is empty
// @router /keyrun [post]
func (this *SaltController) KeyRun() {
	var salt Salt
	err := json.Unmarshal(this.Ctx.Input.RequestBody,&salt)
	if err != nil {
		fmt.Println("参数错误")
	}
	var parse = url.Values{}
	parse.Add("client", "local")
	parse.Add("tgt", salt.Match)
	parse.Add("fun","cmd.run")
	parse.Add("arg",salt.Cmd)
	data := parse.Encode()
	fmt.Println(data)
	body ,_:= exec(data,"application/json","application/x-www-form-urlencoded")
	fmt.Println(string(body))
	jsons ,err:=simplejson.NewJson([]byte(body))
	if err != nil {
		fmt.Println("err",err)
	}
	pre := make(map[string]interface{})
	pre["msg"] = "ok"
	pre["data"] = jsons
	this.ResponseJson(http.StatusOK, e.SUCCESS, pre)
}
//arg=redis.install&client=local&expr_form=list&fun=state.sls&tgt=%2A




// @Title salt执行部署
// @Description salt执行部署
// @Param	match		query	string	true		"match"
// @Param	arg		query	string	true		"arg"
// @Success 200 {string}
// @Failure 404 body is empty
// @router /keydeploy [post]
func (this *SaltController) KeyDeploy() {
	var salt Salt
	err := json.Unmarshal(this.Ctx.Input.RequestBody,&salt)
	if err != nil {
		fmt.Println("参数错误")
	}
	var parse =  url.Values{}
	parse.Add("client","local")
	parse.Add("expr_form","list")
	parse.Add("fun","state.sls")
	parse.Add("tgt",salt.Match)
	parse.Add("arg",salt.Arg)
	data := parse.Encode()
	fmt.Println(data)
	body ,_:= exec(data,"application/json","application/x-www-form-urlencoded")
	fmt.Println(string(body))
	jsons ,err:=simplejson.NewJson([]byte(body))
	if err != nil {
		fmt.Println("err",err)
	}
	pre := make(map[string]interface{})
	pre["data"] = jsons
	pre["msg"] = "ok"
	this.ResponseJson(http.StatusOK, e.SUCCESS, pre)
}
// -H "Accept: application/json" -H "X-Auth-Token: b274499cac06243d9f5a28ccd7f083de5341e53d"
// -d client='local' -d tgt='10.105.75.82' -d fun='cp.get_file' -d arg="salt://top.sls" -d arg="/root/aaa.txt" -d expr_form='list'


// @Title salt推送文件
// @Description salt执行部署
// @Param	run		query	string	true		"run"
// @Param	tgt		query	string	true		"match"
// @Param	arg1		query	string	true		"arg1"
// @Param	arg2		query	string	true		"arg1"
// @Success 200 {string}
// @Failure 404 body is empty
// @router /keycopy [post]
func (this *SaltController) KeyCopyFile() {
	var salt Salt
	err := json.Unmarshal(this.Ctx.Input.RequestBody,&salt)
	if err != nil {
		fmt.Println("参数错误")
	}
	var parse = url.Values{}
	parse.Add("fun",salt.Run)
	parse.Add("tgt",salt.Tgt)
	parse.Add("arg",salt.Arg1)
	parse.Add("arg",salt.Arg2)
	parse.Add("expr_form","list")
	parse.Add("client","local")
	data := parse.Encode()
	fmt.Println(data)
	body ,_:= exec(data,"application/json","application/x-www-form-urlencoded")
	fmt.Println(string(body))
	jsons,err := simplejson.NewJson([]byte(body))
	if err != nil {
		fmt.Println("err", err)
	}
	pre := make(map[string]interface{})
	pre["data"] = jsons
	pre["msg"] ="ok"
	this.ResponseJson(http.StatusOK, e.SUCCESS, pre)

}

func setparse(cctype, tgt, ctype, match,fun string)  string{
	var parse  = url.Values{}
	parse.Add(cctype, tgt)
	parse.Add(ctype, match)
	parse.Add("fun",fun)
	data := parse.Encode()
	return data
}