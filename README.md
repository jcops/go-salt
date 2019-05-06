本项目使用 beego + vue + element-ui 进行开发

```text
    _       __  __  ___ _____
    /_\  /\ /\ \/ / / _ \\_   \
   //_\\/ / \ \  / / /_)/ / /\/
/  _  \ \_/ /  \/ ___/\/ /_
\_/ \_/\___/_/\_\/   \____/

🍭 A NEW API IMAGES STORE TOOL 🍭


```

### go-salt

一个go + vue 前后端项目,包含JWT,RBAC，用户管理，salt远程操作等等。目的是提供一套轻量的中后台开发框架，方便、快速的完成业务需求的开发。 

- [在线演示地址](118.25.39.84:10010) (用户名：admin，密码：123456)（`温馨提醒：为了达到更好的演示效果，这里给出了拥有最高权限的用户，请手下留情！谢谢！`）

  

### 目录结构

- common：基础包
- swagger： 接口文档
- controllers： 接口逻辑处理
- conf：配置文件
- models：应用数据库模型
- routers： 路由逻辑处理

### 权限验证说明

> 利用的casbin库, 将 user role menu 进行自动关联

```
项目启动时,会自动加载权限. 如有更改,会删除对应的权限,重新加载.

用户关联角色  
角色关联菜单  

权限关系为:
角色(role.name,menu.path,menu.method)  
用户(user.username,role.name)

例如:
test      /api/v1/users       GET
eric     test

eric  GET  /api/v1/users 地址的时候，会去检查权限，因为他属于test组，同时组有对应权限，所以本次请求会通过。

用户 admin 有所有的权限,不进行权限匹配

登录接口 /user/login  不进行验证
```

### 请求

> 请求和接收 都是 传递 json 格式 数据

```
例如:
访问 /auth    获取token
{
	"username": "admin",
	"password": "123456"
}

访问  /api/v1/users   
请求头设置  Authorization: Token xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```



## How to run

### Required

- Mysql

### Ready

Create a **go database** and import [SQL](https://github.com/jcops/go-salt/blob/master/data/go.sql)

创建一个库 go-salt,然后导入sql,创建表！

### Conf

You should modify `conf/app.conf`

```
appname = go-salt
httpport = 8080
runmode = dev
autorender = false
copyrequestbody = true
EnableDocs = true

Saltloginapi = https://127.0.0.1:8000/login
Saltapi = https://127.0.0.1:8000
Saltuser = saltapi
Saltpassword = saltapi123

sessionon = true

PageSize = 100
Dbtype = mysql
Dbuser = root
Dbassword = 123
Dbhost = 127.0.0.1:3306
Dbname = go-salt
TablePrefix = go_
```

## Installation

```
yum install go -y 
export GOPROXY=https://goproxy.io
go get github.com/jcops/go-salt
cd $GOPATH/src/github.com/jcops/go-salt
go build main.go
go run  main.go  
```

### Run

```
2019/05/05 13:15:18 [info] replacing callback `gorm:update_time_stamp` from E:/GoWork/src/go-salt/models/db.go:49
2019/05/05 13:15:18 [info] replacing callback `gorm:update_time_stamp` from E:/GoWork/src/go-salt/models/db.go:50
2019/05/05 13:15:18 [info] replacing callback `gorm:delete` from E:/GoWork/src/go-salt/models/db.go:51
2019/05/05 13:15:18.679 [I] [asm_amd64.s:1333]  http server Running on http://:8080

默认 账户admin 密码 123456  
```



## Features

```
- RESTful API
- Gorm
- logging
- Jwt-go
- Swagger
- Beego
- RBAC
- salt
```

### 接口

参考 [接口](https://github.com/jcops/go-salt/blob/master/salt-api.md)

http://127.0.0.1:8080/swagger/index.html



### *example* 

![01](https://github.com/jcops/go-salt/blob/master/data/04.png)

![02](https://github.com/jcops/go-salt/blob/master/data/04.png)

![03](https://github.com/jcops/go-salt/blob/master/data/04.png)

![04](https://github.com/jcops/go-salt/blob/master/data/04.png)



## 感谢以下框架的开源支持

- [Beego] - <http://beego.me/>
- [GORM] - <http://gorm.io/>
- [Casbin] - <https://casbin.org/>

## MIT License

```
Copyright (c) 2019 Eric
```

