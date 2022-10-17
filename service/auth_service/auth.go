package auth_service

import (
	"example.com/my-gin/models"
	"example.com/my-gin/pkg/util"
)

type Auth struct {
	Id       int
	Username string
	Password string
}

func (a *Auth) Check() (bool, error, int) {
	return models.CheckAuth(a.Username, a.Password)
}

func (a *Auth) CheckAuthById() (bool, error) {
	err, auth := models.GetAuthById(a.Id)
	if err == nil && (models.Auth{}) != auth {
		if util.EncodeMD5(auth.Username) == a.Username && util.EncodeMD5(auth.Password) == a.Password {
			return true, nil
		}
	}
	return false, err
}

func (a *Auth) Add() error {
	if err := models.AddUser(a.Username, a.Password); err != nil {
		return err
	}
	return nil
}

func (a *Auth) ModifyPwd() error {
	if err := models.ModifyPwd(a.Id, a.Password); err != nil {
		return err
	}
	return nil
}
