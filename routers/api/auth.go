package api

import (
	"example.com/my-gin/pkg/app"
	"example.com/my-gin/pkg/e"
	"example.com/my-gin/pkg/util"
	"example.com/my-gin/service/auth_service"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50); MinSize(5);"`
	Password string `valid:"Required; MaxSize(50); MinSize(5);"`
}

func Register(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	username := c.PostForm("username")
	password := c.PostForm("password")

	a := auth{Username: username, Password: password}
	if ok, _ := valid.Valid(&a); !ok {
		fmt.Println(username, password)
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	encryptedPwd := util.EncodeMD5(password)
	authService := &auth_service.Auth{Username: username, Password: encryptedPwd}
	isExist, err, _ := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	if isExist {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_USERNAME_EXISTS, nil)
		return
	}

	if err := authService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func Login(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username, password)

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	encryptedPwd := util.EncodeMD5(password)
	authService := &auth_service.Auth{Username: username, Password: encryptedPwd}
	isExist, err, authId := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH_USERNAME_OR_PWD, nil)
		return
	}

	token, err := util.GenerateToken(username, encryptedPwd, authId)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
