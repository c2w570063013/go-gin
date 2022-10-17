package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Auth struct {
	Model
	//ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) (bool, error, int) {
	var auth Auth
	err := db.Table("blog_auth").Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	fmt.Println(auth, "(((((((((((((((((((", auth.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err, auth.ID
	}

	if auth.ID > 0 {
		return true, nil, auth.ID
	}
	return false, nil, auth.ID
}

func GetAuthById(id int) (error, Auth) {
	var (
		auth Auth
		err  error
	)
	err = db.Select("id,username,password").Where("id=?", id).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err, auth
	}
	if auth.ID == id {
		return nil, auth
	}
	return nil, auth
}

func AddUser(username, password string) error {
	auth := Auth{
		Username: username,
		Password: password,
	}
	if err := db.Create(&auth).Error; err != nil {
		return err
	}
	return nil
}

func ModifyPwd(id int, password string) error {
	if err := db.Model(&Auth{}).Where("id=?", id).Update("password", password).Error; err != nil {
		return err
	}
	return nil
}
