package routers

import (
	"example.com/my-gin/middleware/jwt"
	"example.com/my-gin/routers/api"
	"example.com/my-gin/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//r := gin.Default()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/login", api.Login)
	r.POST("/register", api.Register)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		apiV1.GET("/album", v1.GetAlbums)
		apiV1.GET("/getByAlbumId/:id", v1.GetAlbumByID)
		apiV1.POST("/album", v1.PostAlbums)
		apiV1.POST("/changePwd", api.ChangePwd)
	}
	return r
}
