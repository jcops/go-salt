package salt

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
	"crypto/tls"
	"io/ioutil"
)

type Juser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Eauth string `json:"eauth"`
}

type Jreturn struct {
	Perms []string `json:"perms"`
	Start float64 `json:"start"`
	Token string `json:"token"`
	Expire float64 `json:"expire"`
	User string `json:"user"`
	Eauth string `json:"eauth"`
}

type Jslice struct {
	Return []Jreturn `json:"return"`
}
// 获取token
func GetToken() string{
	saltloginapi := beego.AppConfig.String("Saltloginapi")
	var juser Juser
	juser.Username = beego.AppConfig.String("Saltuser")
	juser.Password = beego.AppConfig.String("Saltpassword")
	juser.Eauth = "pam"

	js ,err := json.Marshal(juser)
	if err != nil {
		fmt.Println("json marshal err",err)
	}
	fmt.Println(string(js))
	var json_str = js
	req,err := http.NewRequest("POST", saltloginapi, bytes.NewBuffer(json_str))
	req.Header.Set("Accept","application/json")
	req.Header.Set("Content-Type","application/json")

	tr := &http.Transport{
		TLSClientConfig:&tls.Config{InsecureSkipVerify:true},
		DisableCompression:true,
	}

	client := http.Client{Transport:tr}
	resp,err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer  resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var res Jslice
	if err = json.Unmarshal([]byte(string(body)),&res);err != nil {
		fmt.Println(err)
	}
	var token string
	for _,v := range res.Return {
		token = v.Token
	}
	fmt.Println(token)
	return token
}
