package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"reflect"
	"fmt"
	"strings"
	"github.com/astaxie/beego/context"
	"net/http"
	"encoding/json"
	"go-salt/common/e"
	"go-salt/common"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// 验证token拦截器
var JWT =func(ctx *context.Context) {
	var code int
	var data interface{}
	code = e.SUCCESS
	Authorization := ctx.Input.Header("Authorization")
	fmt.Println(Authorization)
	token := strings.Split(Authorization, " ")
	fmt.Println(token)
	if Authorization == "" {
		code = e.NO_AUTH
		fmt.Println(Authorization)
	} else {
		fmt.Println("token",token[1])
		_, err := ParseToken(token[1])
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			default:
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}
		}
	}
	if code != e.SUCCESS {
		//http.Error(ctx.ResponseWriter, "Token verification error", http.StatusBadRequest)
		ResponseWithJson(ctx.ResponseWriter,http.StatusBadRequest,
			Response{Code: code, Msg: e.GetMsg(code),Data:data})
		return
	}

}


var JwtSecret = []byte("jcops")

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}


// 生成Jwt Token
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		username,
		common.EncodeMD5(password),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "jcops",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecret)
	return token, err
}

// 验证JWT Token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}




func GetIdFromClaims(key string, claims jwt.Claims) string {
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)

			if fmt.Sprintf("%s", k.Interface()) == key {
				return fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return ""
}
