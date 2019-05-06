package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
	Role     []Role `json:"role" gorm:"many2many:user_role;"`

}

func CheckUser(username,password string) (*User,error){
	var user User
	err := db.Model(&User{}).Where(User{Username:username,Password:password}).First(&user).Error
	if err != nil {
		fmt.Println("model",err)
		return nil,err
	}
	fmt.Println(user)
	//if user.ID > 0 {
	//	return true, nil
	//}
	return &user,nil
}

func CheckUserUsername(username string) (bool,error){
	var user User
	err := db.Where("username = ? and deleted_on = ?",username,0).First(&user).Error
	if err != nil{
		fmt.Println("model err",err)
		return false,err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false,err
}
func AddUser(data map[string]interface{}) (id int,err error){
	user := User{
		Username: data["username"].(string),
		Password: data["password"].(string),
	}
	var role []Role
	db.Where("id in (?)", data["role_id"].(int)).Find(&role)
	if err := db.Create(&user).Association("Role").Append(role).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func GetUser(id int) (*User, error) {
	var user User
	err := db.Preload("Role").Where("id = ? AND deleted_on = ? ", id, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}




func GetUsersAll() ([]*User, error) {
	var user []*User
	err := db.Where("deleted_on = ? ", 0).Preload("Role").Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return user, nil
}

func GetUserTotal(maps interface{})(int,error) {
	var count int
	if err := db.Model(&User{}).Where(maps).Count(&count).Error; err != nil {
		return 0,err
	}
	return count,nil
}
func GetUsers(PageNum,PageSize int,maps interface{})([]*User,error) {
	var user []*User
	fmt.Println("页码",PageNum,"每页条数",PageSize)
	err := db.Preload("Role").Where(maps).Offset(PageNum).Limit(PageSize).Find(&user).Error
		if err != nil {
			return nil, err
	}
	return user,nil
}

func ExistUserById(id int) (bool,error) {
	var user User
	err := db.Select("id").Where("id = ? AND deleted_on = ?",id,0).First(&user).Error
	if err != nil {
		return false,err
	}
	if user.ID >0 {
		return true,err
	}
	return false,err

}

func DeleteUser(uid int) (error) {
	var user User
	db.Where("id = ?",uid).Find(&user)
	fmt.Println(user)
	db.Model(&user).Association("Role").Delete()
	if err := db.Where("id = ?",uid).Delete(&user).Error;err != nil {
		return err
	}
	fmt.Println(user)
	return nil
}

func EditUser(uid int,data map[string][]int) error {
	var role []Role
	var user User
	db.Where("id in (?)",data["role_id"]).Find(&role)
	fmt.Println("ro",role)
	if err := db.Where("id = ? AND deleted_on = ?",uid,0).First(&user).Error;err != nil {
		return err
	}
	// 使用覆盖之前的角色
	if err := db.Model(&user).Association("Role").Replace(role).Error;err != nil {
		fmt.Println("修改关联角色错误",err)
	}
	db.Model(&user).Update(data)
	return nil
}

func UpdateUswePassword(id int, password string) error {
	var user User
	if err := db.Where("id = ? AND deleted_on = ?",id,0).First(&user).Update("password",password).Error;err != nil {
		return err
	}
	//db.Model(&user).Update(password)
	return nil
}

func MigrateUsers() error {
	if db.HasTable(&User{}) {
		err := db.DropTable(&User{}).Error
		err = db.CreateTable(&User{}).Error
		return err
	} else {
		err := db.CreateTable(&User{}).Error
		return err
	}
}
