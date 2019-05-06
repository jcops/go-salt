用户登录获取Token

- API: POST  /user/login
- 功能说明：
  - 获取Token
- input body：

```
{
	"username":"ada2",
	"password":"123456"
}
```

- input字段说明:

| 名称     | 类型   | 必填 | 默认值 | 说明   |
| -------- | ------ | ---- | ------ | ------ |
| username | string | 是   | 无     | 用户名 |
| password | string | 是   | 无     | 密码   |

- 返回信息

  ```
  {
      "code": 200,
      "msg": "ok",
      "data": {
          "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkYTIiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNTU1NjY3NzMwLCJpc3MiOiJqY29wcyJ9.jUH62ZvIiTGtAStE-YSGq0eyTQmVGpyvyIVbV9O93E4"
      }
  }
  ```

  

  

### 新建用户

- API: POST  /api/v1/register
- 功能说明：
  - 新建用户
- input body：

```
{
	"username":"ada2",
	"password":"123456",
	"role_id":30

}
```

- input字段说明:

| 名称     | 类型   | 必填 | 默认值 | 说明   |
| -------- | ------ | ---- | ------ | ------ |
| username | string | 是   | 无     | 用户名 |
| password | string | 是   | 无     | 密码   |
| role_id  | int    | 否   | 无     | 角色id |

- 返回信息

  ```
  {
      "code": 200,
      "msg": "ok",
      "data": null
  }
  ```

  

### 查询所有用户

- API: GET /api/v1/users
- 功能说明：

  - 查询所有用户

- 返回信息

  ```
  {
      "code": 200,
      "msg": "ok",
      "data": {
          "lists": [
              {
                  "id": 33,
                  "created_on": 1555585720,
                  "modified_on": 1555586199,
                  "deleted_on": 0,
                  "username": "ada",
                  "password": "d41d8cd98f00b204e9800998ecf8427e",
                  "role": null
              },
              {
                  "id": 35,
                  "created_on": 1555586467,
                  "modified_on": 1555586467,
                  "deleted_on": 0,
                  "username": "ada1",
                  "password": "e10adc3949ba59abbe56e057f20f883e",
                   "role": [
                      {
                          "id": 30,
                          "created_on": 1555584631,
                          "modified_on": 1555584631,
                          "deleted_on": 0,
                          "name": "sa2",
                          "menu": null
                      }
                  ]
              },
  
          ],
          "total": 2
      }
  }
  ```

  

### 查询单个用户

- API: GET /api/v1/users/:id
- 功能说明：
  - 查询单个用户



- 返回信息

  ```
  {
      "code": 200,
      "msg": "ok",
      "data": {
          "id": 36,
          "created_on": 1555586548,
          "modified_on": 1555586548,
          "deleted_on": 0,
          "username": "ada2",
          "password": "e10adc3949ba59abbe56e057f20f883e",
          "role": [
              {
                  "id": 30,
                  "created_on": 1555584631,
                  "modified_on": 1555584631,
                  "deleted_on": 0,
                  "name": "sa2",
                  "menu": null
              }
          ]
      }
  }
  ```

  



### 修改用户信息

- API: PUT /api/v1/users
- 功能说明：
  - 修改用户信息
- input body：

```
{
	"id":35,
	"role_id":[30]
}
```

- input字段说明:

| 名称    | 类型  | 必填 | 默认值 | 说明   |
| ------- | ----- | ---- | ------ | ------ |
| id      | int   | 是   | 无     | 用户id |
| role_id | []int | 否   | 无     | 角色id |

- 返回信息

  ```
  {
      "code": 200,
      "msg": "ok",
      "data": null
  }
  ```

  

### 删除用户

- API: DELETE  /api/v1/users/:id
- 功能说明：
  - 删除用户

```
{
	"id":35
}
```



- 返回信息

```
  {
      "code": 200,
      "msg": "ok",
      "data": null
  }
```



### 修改用户密码

- API: POST /api/v1/users
- 功能说明：
  - 修改用户密码
- input path：

```
{
	"id":35,
	"password":"123123"
}
```

- input字段说明:

| 名称     | 类型   | 必填 | 默认值 | 说明                         |
| -------- | ------ | ---- | ------ | ---------------------------- |
| id       | int    | 是   | 无     | 用户id                       |
| password | string | 是   | 无     | 新密码（前端判断密码一致性） |

- 返回信息

```
{
    "code": 200,
    "msg": "ok",
    "data": null
}
```





### 查询所有角色

- API: GET  /api/v1/role
- 功能说明：
  - 查询所有角色

- 返回信息

```
{
    "code": 200,
    "msg": "ok",
    "data": {
        "lists": [
            {
                "id": 30,
                "created_on": 1555584631,
                "modified_on": 1555584631,
                "deleted_on": 0,
                "name": "sa2",
                "menu": [
                    {
                        "id": 1,
                        "created_on": 0,
                        "modified_on": 0,
                        "deleted_on": 0,
                        "name": "查询所有用户",
                        "path": "/api/v1/users",
                        "method": "GET"
                    },
                    {
                        "id": 2,
                        "created_on": 0,
                        "modified_on": 0,
                        "deleted_on": 0,
                        "name": "注册用户",
                        "path": "/api/v1/register",
                        "method": "POST"
                    },
                    {
                        "id": 3,
                        "created_on": 0,
                        "modified_on": 0,
                        "deleted_on": 0,
                        "name": "查询单个用户",
                        "path": "/api/v1/users/:id",
                        "method": "GET"
                    },
                    {
                        "id": 4,
                        "created_on": 0,
                        "modified_on": 0,
                        "deleted_on": 0,
                        "name": "删除用户",
                        "path": "/api/v1/users/:id",
                        "method": "DELETE"
                    },
                    {
                        "id": 9,
                        "created_on": 0,
                        "modified_on": 0,
                        "deleted_on": 0,
                        "name": "修改角色",
                        "path": "/api/v1/role",
                        "method": "PUT"
                    },
                    {
                        "id": 11,
                        "created_on": 0,
                        "modified_on": 0,
                        "deleted_on": 0,
                        "name": "接受节点认证请求",
                        "path": "/api/v1/keyaccept/:id",
                        "method": "GET"
                    }
                ]
            },
            {
                "id": 32,
                "created_on": 1555587496,
                "modified_on": 1555587496,
                "deleted_on": 0,
                "name": "dev1",
                "menu": [
                    {
                        "id": 1,
                        "created_on": 0,
                        "modified_on": 0,
                        "deleted_on": 0,
                        "name": "查询所有用户",
                        "path": "/api/v1/users",
                        "method": "GET"
                    },
                    {
                        "id": 2,
                        "created_on": 0,
                        "modified_on": 0,
                        "deleted_on": 0,
                        "name": "注册用户",
                        "path": "/api/v1/register",
                        "method": "POST"
                    },
                    {
                        "id": 3,
                        "created_on": 0,
                        "modified_on": 0,
                        "deleted_on": 0,
                        "name": "查询单个用户",
                        "path": "/api/v1/users/:id",
                        "method": "GET"
                    },
                    {
                        "id": 4,
                        "created_on": 0,
                        "modified_on": 0,
                        "deleted_on": 0,
                        "name": "删除用户",
                        "path": "/api/v1/users/:id",
                        "method": "DELETE"
                    },
                    {
                        "id": 5,
                        "created_on": 0,
                        "modified_on": 0,
                        "deleted_on": 0,
                        "name": "修改用户信息",
                        "path": "/api/v1/users",
                        "method": "PUT"
                    },
                    {
                        "id": 6,
                        "created_on": 0,
                        "modified_on": 0,
                        "deleted_on": 0,
                        "name": "查询所有角色",
                        "path": "/api/v1/role",
                        "method": "GET"
                    }
                ]
            }
        ],
        "total": 2
    }
}
```



### 查询单个角色

- API: GET /api/v1/users/:id
- 功能说明：
  - 查询单个用户



- 返回信息

  

### 新建角色

- API: POST  /api/v1/role
- 功能说明：
  - 新建角色
- input body：

```
{
	"name":"test111",
	"menu_id":[1,2,3,4]
}
```

- input字段说明:

| 名称    | 类型   | 必填 | 默认值 | 说明   |
| ------- | ------ | ---- | ------ | ------ |
| name    | string | 是   | 无     | 用户名 |
| menu_id | []int  | 否   | 无     | 菜单id |

- 返回信息

```
{
    "code": 200,
    "msg": "ok",
    "data": null
}
```





### 修改角色信息

- API: PUT /api/v1/role
- 功能说明：
  - 修改角色信息
- input body：

```
{
	"id":30,
	"menu_id":[1,2,3,4,11,9]
}
```

- input字段说明:

| 名称    | 类型  | 必填 | 默认值 | 说明   |
| ------- | ----- | ---- | ------ | ------ |
| id      | int   | 是   | 无     | 用户id |
| role_id | []int | 否   | 无     | 角色id |

- 返回信息

```
{
    "code": 200,
    "msg": "ok",
    "data": null
}
```



### 删除角色

- API: DELETE  /api/v1/role/:id
- 功能说明：
  - 删除用户
- input path

```
{
	"id":35
}
```

- 返回信息

```
{
    "code": 200,
    "msg": "ok",
    "data": null
}
```



### 新建菜单

- API: POST  /api/v1/menu
- 功能说明：
  - 新建菜单
- input body：

```
{
	"name":"test11122",
	"path":"/api/v1/test",
	"method":"GET"
}
```

- input字段说明:

| 名称   | 类型   | 必填 | 默认值 | 说明     |
| ------ | ------ | ---- | ------ | -------- |
| name   | string | 是   | 无     | 菜单名   |
| path   | string | 是   | 无     | 菜单路径 |
| method | string | 是   | 无     | 请求方法 |

- 返回信息

```
{
    "code": 200,
    "msg": "ok",
    "data": null
}
```





### 删除菜单

- API: DELETE  /api/v1/menu/:id
- 功能说明：
  - 删除菜单
- input path

```
{
	"id":22
}
```

- 返回信息

```
{
    "code": 200,
    "msg": "ok",
    "data": null
}
```



### 修改菜单信息

- API: PUT /api/v1/menu
- 功能说明：
  - 修改菜单信息
- input body：

```
{
	"id":21,
	"name":"修改用户密码",
	"path":"/api/v1/users",
	"method":"POST"
}
```

- input字段说明:

| 名称   | 类型   | 必填 | 默认值 | 说明     |
| ------ | ------ | ---- | ------ | -------- |
| id     | int    | 是   | 无     | 用户id   |
| name   | string | 否   | 无     | 菜单名   |
| path   | string | 否   | 无     | 菜单路径 |
| method | string | 否   | 无     | 请求方法 |

- 返回信息

```
{
    "code": 200,
    "msg": "ok",
    "data": null
}
```



### 查询所有菜单

- API: GET /api/v1/users
- 功能说明：
  - 查询所有菜单
- 返回信息

```
{
    "code": 200,
    "msg": "ok",
    "data": {
        "lists": [
            {
                "id": 1,
                "created_on": 0,
                "modified_on": 0,
                "deleted_on": 0,
                "name": "查询所有用户",
                "path": "/api/v1/users",
                "method": "GET"
            },
            {
                "id": 2,
                "created_on": 0,
                "modified_on": 0,
                "deleted_on": 0,
                "name": "注册用户",
                "path": "/api/v1/register",
                "method": "POST"
            }
        ],
        "total": 2
    }
}
```



### 查询节点列表

- API: GET /api/v1/users
- 功能说明：
  - 查询节点列表
- 返回信息

```
{
    "data": {
        "minions": [
            "10.105.75.82"
        ],
        "minions_pre": [],
        "msg": "ok"
    }
}
```

### 检查节点连通性

- API: GET /api/v1/keyping
- 功能说明：
  - 检查节点连通性
- 返回信息

```
{
    "data": {
        "return": [
            {
                "10.105.75.82": true
            }
        ]
    },
    "msg": "ok"
}
```



### 认证节点

- API: POST /api/v1/keyaccept
- 功能说明：
  - 认证节点
- input body：

```
{
	"match":10.0.1.1
}
```

- input字段说明:

| 名称  | 类型   | 必填 | 默认值 | 说明       |
| ----- | ------ | ---- | ------ | ---------- |
| match | string | 是   | 无     | 认证节点ip |



### 删除认证节点

- API: POST /api/v1/keydelete
- 功能说明：
  - 删除认证节点
- input body：

```
{
	"match":10.0.1.1
}
```

- input字段说明:

| 名称  | 类型   | 必填 | 默认值 | 说明       |
| ----- | ------ | ---- | ------ | ---------- |
| match | string | 是   | 无     | 认证节点ip |



### 执行命令

- API: POST /api/v1/keyrun
- 功能说明：
  - 认证节点
- input body：

```
{
	"match":"10.0.1.1",
	"cmd":"ls"
}
```

- input字段说明:

| 名称  | 类型   | 必填 | 默认值 | 说明   |
| ----- | ------ | ---- | ------ | ------ |
| match | string | 是   | 无     | 节点ip |
| cmd   | string | 是   | 无     | 命令   |

### 

### 部署软件

- API: POST /api/v1/keydeploy
- 功能说明：
  - 部署软件
- input body：

```
{
	"match":"10.0.1.1",
	"arg":"system_init"
}
```

- input字段说明:

| 名称  | 类型   | 必填 | 默认值 | 说明          |
| ----- | ------ | ---- | ------ | ------------- |
| match | string | 是   | 无     | 认证节点ip    |
| arg   | string | 是   | 无     | 部署的sls文件 |



### 拷贝文件到节点

- API: POST /api/v1/keydeploy
- 功能说明：
  - 拷贝文件到节点
- input body：

```
{
	"tgt":"10.0.1.1",
	"arg1":"/opt/aa.txt",
	"arg2":"/opt/"
	"run":"cp.get_file"
}
```

- input字段说明:

| 名称 | 类型   | 必填 | 默认值 | 说明     |
| ---- | ------ | ---- | ------ | -------- |
| tgt  | string | 是   | 无     | 节点ip   |
| arg1 | string | 是   | 无     | 源路径   |
| arg2 | string | 是   | 无     | 目标路径 |
| run  | string | 是   | 无     | 模块     |


