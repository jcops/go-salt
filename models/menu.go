package models

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/jinzhu/gorm"
)

type Menu struct {
	Model
	Name  string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

func ExistsMenu(name string) (bool,error) {
	var menu Menu
	if err := db.Where("name = ? AND deleted_on = ?",name,0).Find(&menu).Error;err!= nil {
		return false,err
	}
	if menu.ID > 0 {
		return true, nil
	}
	return false,nil
}

func ExistsMenuById(id int) (bool,error) {
	var menu Menu
	if err := db.Where("id = ? AND deleted_on = ?",id,0).Find(&menu).Error;err!= nil {
		return false,err
	}
	if menu.ID > 0 {
		return true, nil
	}
	return false,nil
}

func UpdateMenu(data map[string]interface{}) error {
	var menu Menu
	db.Where("id = ? AND deleted_on = ?",data["id"],0).Find(&menu)
	if err := db.Model(&menu).Updates(data).Error; err != nil {
		return err
	}
	return nil
}


func AddMenu(data map[string]string) (int,error) {
	menu  := Menu{
		Name:data["name"],
		Method:data["method"],
		Path:data["path"],
	}
	if err := db.Create(&menu).Error;err != nil {
		return 0,err
	}
	fmt.Println(menu)
	return menu.ID,nil
}

func DeleteMenu(mid int) (error) {
	var menu Menu
	err := db.Where("id = ? AND deleted_on = ?",mid,0).Find(&menu).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println(err)
		return errors.New("not")
	}
	db.Model(&menu).Association("Role").Delete()
	if err := db.Where("id = ? AND deleted_on = ?",mid,0).Delete(&menu).Error;err != nil {
		return err
	}
	return nil
}
func GetMenuTotal(maps map[string]interface{}) (id int,err error){
	var count int
	if err := db.Model(&Menu{}).Where(maps).Count(&count).Error;err != nil {
		return 0,err
	}
	return count,err
}

func GetMenus(PageNum,PageSize int,maps interface{} ) ([]*Menu,error) {
	var menu []*Menu
	fmt.Println("页码",PageNum,"每页条数",PageSize)
	if err := db.Model(&Menu{}).Where(maps).Offset(PageNum).Limit(PageSize).Find(&menu).Error;err != nil {
		return nil,err
	}
	return menu,nil
}