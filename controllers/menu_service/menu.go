package menu_service

import (
	"github.com/casbin/casbin"
	"go-salt/models"
	"errors"
)

type Menu struct {
	ID     int
	Name   string
	Path   string
	Method string

	CreatedBy  string
	ModifiedBy string

	PageNum  int
	PageSize int

	Menu     *models.Menu     `inject:""`
	Enforcer *casbin.Enforcer `inject:""`
}


func (a *Menu) Addmenu() (id int,err error){
	exists,_ := models.ExistsMenu(a.Name)
	if exists {
		return 0,errors.New("菜单已重复")
	}
	menu := make(map[string]string)
	menu["name"] = a.Name
	menu["method"] = a.Method
	menu["path"] = a.Path

	id,err = models.AddMenu(menu)
	if err != nil {
		return 0,err
	}
	return id,nil
}

func (a *Menu) DeleteMenu() error {
	return models.DeleteMenu(a.ID)
}

func (a *Menu) UpdateMenu() error {
	exists,_ := models.ExistsMenuById(a.ID)
	if !exists {
		return errors.New("菜单不存在")
	}
	menu := make(map[string]interface{})
	menu["name"] = a.Name
	menu["method"] = a.Method
	menu["path"] = a.Path
	menu["id"] = a.ID
	err := models.UpdateMenu(menu)
	if err != nil {
		return err
	}
	return nil


}
func (a *Menu) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	return maps
}

func (a *Menu) Count()(id int ,err error) {
	return models.GetMenuTotal(a.getMaps())
}

func (a *Menu) GetAll() ([]*models.Menu,error) {
	menu, err := models.GetMenus(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}
	return menu,nil
}