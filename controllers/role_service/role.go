package role_service

import (
	"github.com/casbin/casbin"
	"go-salt/models"
	"fmt"
	"errors"
)
type Role struct {
	ID   int
	Name string
	Menu int

	CreatedBy  string
	ModifiedBy string

	PageNum  int
	PageSize int

	Enforcer *casbin.Enforcer `inject:""`
}


func (a *Role) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	return maps
}

func (a *Role) Count() (int, error) {
	return models.GetRoleTotal(a.getMaps())
}

func (a *Role) GetAll()([]*models.Role,error) {
	if a.Name != "" {
		maps := make(map[string]interface{})
		maps["deleted_on"] = 0
		maps["name"] = a.Name
		role,err := models.GetRoles(a.PageNum,a.PageSize,maps)
		if err != nil {
			return nil,err
		}
		return role, nil
	}else {
		role, err := models.GetRoles(a.PageNum, a.PageSize, a.getMaps())
		if err != nil {
			return nil, err
		}
		return role,nil
	}
}

func (a *Role) ExistsRole() (bool,error){
	return models.ExistsRoleByID(a.ID)
}
func (a *Role) DeleteRole() error {
	return models.DeleteRole(a.ID)
}

func  (a *Role) AddRole(data map[string][]int) (id int, err error){
	exists,_ :=models.ExistsRole(a.Name)
	if exists {
		return 0,errors.New("角色已存在")
	}
	fmt.Println("传参",data)
	if id, err := models.AddRole(a.Name,data); err == nil {
		return id, nil
	} else {
		return 0, err
	}
}


func  (a *Role) UpdateRole(data map[string][]int) (id int, err error){
	exists,_ :=models.ExistsRoleByID(a.ID)
	if !exists {
		return 0,errors.New("角色不存在")
	}
	fmt.Println("传参",data)
	if id, err := models.UpdateRole(a.ID,data); err == nil {
		return id, nil
	} else {
		return 0, err
	}
}



// LoadAllPolicy 加载所有的角色策略
func (a *Role) LoadAllPolicy() error {
	roles, err := models.GetRolesAll()
	if err != nil {
		return err
	}

	for _, role := range roles {
		err = a.LoadPolicy(role.ID)
		if err != nil {
			return err
		}
	}
	//fmt.Println("权限关系", a.Enforcer.GetGroupingPolicy())
	fmt.Println("role_service/role.go,角色菜单权限关系", a.Enforcer.GetGroupingPolicy())
	return nil
}

// LoadPolicy 加载角色权限策略
func (a *Role) LoadPolicy(id int) error {
	role, err := models.GetRole(id)
	if err != nil {
		return err
	}
	// 如有更改,会删除对应的权限,重新加载.
	a.Enforcer.DeletePermissionsForUser(role.Name)
	for _, menu := range role.Menu {
		//fmt.Println("菜单",menu.ID,menu.Name,menu.Path, menu.Method)
		if menu.Path == "" || menu.Method == "" {
			continue
		}
		a.Enforcer.AddPermissionForUser(role.Name,menu.Path, menu.Method)
		// 重新加载.
		//a.Enforcer.AddRoleForUserInDomain(role.Name,menu.Path, menu.Method)
		//a.Enforcer.AddPolicy(role.Name,menu.Path, menu.Method)
		//a.Enforcer.AddGroupingPolicySafe(role.Name,menu.Path, menu.Method)
	}
	//fmt.Println("role_service/role.go,更新菜单权限关系", a.Enforcer.GetPolicy())
	return nil
}

