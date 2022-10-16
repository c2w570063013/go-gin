package main

import (
	"example.com/my-gin/models"
	"example.com/my-gin/pkg/logging"
	"example.com/my-gin/pkg/setting"
	"example.com/my-gin/pkg/util"
	"example.com/my-gin/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	//"os"
	//"time"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	util.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
	//routersInit.Run("localhost:8080")
}
