package user_service

import (
	"github.com/casbin/casbin"
	"go-salt/models"
	"go-salt/common"
	"errors"
	"fmt"
)

type User struct {
	ID       int
	Username string
	Password string
	Role     int

	CreatedBy  string
	ModifiedBy string

	PageNum  int
	PageSize int

	Enforcer *casbin.Enforcer `inject:""`
}
func (a *User) Check() (*models.User, error) {
	return models.CheckUser(a.Username, common.EncodeMD5(a.Password))
}

func (a *User) AddUser() (id int, err error) {
	menu := map[string]interface{}{
		"username": a.Username,
		"password": common.EncodeMD5(a.Password),
		"role_id":  a.Role,
	}
	username, _ := models.CheckUserUsername(a.Username)

	if username {
		return 0, errors.New("用户已存在")
	}
	if id, err := models.AddUser(menu); err == nil {
		return id, err
	} else {
		return 0, err
	}
}
func (a *User) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	return maps
}


func (a *User) Count() (int,error) {
	return models.GetUserTotal(a.getMaps())
}


func (a *User) GetAll() ([]*models.User,error) {
	fmt.Println(a.PageNum,a.PageSize)
	if a.Username != "" {
		maps := make(map[string]interface{})
		maps["deleted_on"] = 0
		maps["username"] = a.Username
		fmt.Println(a.Username)
		user,err := models.GetUsers(a.PageNum,a.PageSize,maps)
		if err != nil {
			return nil,err
		}
		return user,nil
	}else {
		user,err := models.GetUsers(a.PageNum,a.PageSize,a.getMaps())
		if err != nil {
			return nil,err
		}
		return user,nil
	}
}


func (a *User) ExistById() (bool, error){
	return models.ExistUserById(a.ID)
}

func (a *User) GetOneUser() (*models.User, error){
	user,err := models.GetUser(a.ID)
	if err != nil {
		return nil,err
	}
	return user,nil
}

func (a *User) DeleteUser() (error) {
	return models.DeleteUser(a.ID)
}

func (a *User) EditUser(data map[string][]int) error{
	err := models.EditUser(a.ID,data)
	if err != nil {
		return err
	}
	return nil
}

func (a *User) UpdateUswePassword() error{
	if err := models.UpdateUswePassword(a.ID,a.Password); err != nil{
		return err
	}
	return nil

}



// LoadAllPolicy 加载所有的用户策略
func (a *User) LoadAllPolicy() error {
	users, err := models.GetUsersAll()
	if err != nil {
		return err
	}
	for _, user := range users {
		if len(user.Role) != 0 {
			err = a.LoadPolicy(user.ID)
			if err != nil {
				return err
			}
		}
	}
	fmt.Println("user_service/user.go,用户角色权限关系", a.Enforcer.GetGroupingPolicy())
	return nil
}

// LoadPolicy 加载用户权限策略
func (a *User) LoadPolicy(id int) error {
	user, err := models.GetUser(id)
	if err != nil {
		return err
	}
	a.Enforcer.DeleteRolesForUser(user.Username)
	for _, ro := range user.Role {
		// 重新加载.
		a.Enforcer.AddRoleForUser(user.Username, ro.Name)
	}
	fmt.Println("user_service/user.go,更新角色权限关系", a.Enforcer.GetGroupingPolicy())
	return nil
}
