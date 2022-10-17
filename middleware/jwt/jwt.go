package jwt

import (
	"example.com/my-gin/pkg/e"
	"example.com/my-gin/pkg/setting"
	"example.com/my-gin/pkg/util"
	"example.com/my-gin/service/auth_service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data any

		code = e.SUCCESS
		//token := c.Query("token")
		token := c.Request.Header.Get("token")
		if token == "" {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else {
			tokenClaims, err := util.ParseToken(token)
			//fmt.Println(err, tokenClaims, "iiiiiiiiiiiiiiiiii")
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}

			//check whether username or password has changed in database
			if code == e.SUCCESS && tokenClaims != nil {
				authService := &auth_service.Auth{
					Username: tokenClaims.Username,
					Password: tokenClaims.Password,
					Id:       tokenClaims.Id,
				}
				isExist, err := authService.CheckAuthById()
				if !isExist || err != nil {
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}

				//Cache user info
				setting.UserId = tokenClaims.Id
				setting.UserMd5EncodedPwd = tokenClaims.Password
			}

		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}

}
