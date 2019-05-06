package models

import (
	"fmt"
)

type Role struct {
	Model
	Name string `json:"name"`
	Menu []Menu `json:"menu" gorm:"many2many:role_menu;"`
}


func GetRoles(pageNum,pageSize int, data interface{}) ([]*Role,error) {
	var role []*Role
	err := db.Preload("Menu").Where(data).Offset(pageNum).Limit(pageSize).Find(&role).Error
	if err != nil {
		return nil, err
	}
	return role,nil
}
func GetRolesAll() ([]*Role,error) {
	var role []*Role
	err := db.Preload("Menu").Where("deleted_on=?",0).Find(&role).Error
	if err != nil {
		return nil, err
	}
	return role,nil
}
func ExistsRoleByID(rid int) (bool,error) {
	var role Role
	if err := db.Model(&Role{}).Where("id = ? AND deleted_on = ?",rid,0).First(&role).Error;err != nil {
		return false,err
	}
	return true,nil
}
func ExistsRole(ro string) (bool,error) {
	var role Role
	if err := db.Model(&Role{}).Where("name = ? AND deleted_on = ?",ro,0).First(&role).Error;err != nil {
		return false,err
	}
	if role.ID > 0 {
		return true, nil
	}
	return false,nil
}


func AddRole(name string,data map[string][]int) (id int,err error) {
	role  := Role{
		Name:name,
	}
	var menu []Menu
	//fmt.Println(strings.Replace(strings.Trim(fmt.Sprint(v), "[]"), " ", ",", -1))
	db.Where("id in (?)", data["menu_id"]).Find(&menu)
	fmt.Println(menu)
	if err = db.Create(&role).Association("Menu").Append(menu).Error; err != nil {
		return 0, err
	}
	return role.ID, nil
}

func UpdateRole(rid int,data map[string][]int) (id int,err error) {
	role  := Role{}
	var menu []Menu
	fmt.Println("query menu id",data["menu_id"])
	//fmt.Println(strings.Replace(strings.Trim(fmt.Sprint(v), "[]"), " ", ",", -1))
	db.Where("id in (?)", data["menu_id"]).Find(&menu)
	if err := db.Model(&Role{}).Where("id = ? AND deleted_on = ?",rid,0).First(&role).Error;err != nil {
		return 0,err
	}
	// 使用覆盖之前的菜单权限
	err = db.Model(&role).Association("Menu").Replace(menu).Error
	if err != nil {
		fmt.Println(err)
		return 0,err
	}
	//db.Model(&role).Update(data)
	return role.ID, nil
}




func DeleteRole(rid int )error {
	var role Role
	db.Where("id = ? AND deleted_on = ?",rid,0).Find(&role)
	db.Model(&role).Association("Menu").Delete()
	if err := db.Where("id = ? AND deleted_on = ?",rid,0).Delete(&role).Error;err!= nil {
		return err
	}
	db.Model(&role).Association("Menu").Clear()
	return nil
}
func GetRole(id int) (*Role,error) {
	var role Role
	err := db.Preload("Menu").Where("id = ? AND deleted_on = ? ", id, 0).First(&role).Error
	if err != nil {
		return nil,err
	}
	return &role,nil
}

func GetRoleTotal(data interface{}) (int ,error) {
	var count int
	if err := db.Model(&Role{}).Where(data).Count(&count).Error;err != nil {
		return 0,err
	}
	return count,nil
}
