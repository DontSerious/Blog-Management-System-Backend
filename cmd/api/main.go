package main

import (
	"Bishe/be/cmd/api/handlers"
	"Bishe/be/cmd/api/rpc"
	"Bishe/be/pkg/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() {
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.Default()
	r.Use(middlewares.Cors())

	// user
	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)
	r.POST("/changePWD", handlers.ChangePWD)
	r.GET("/delUser", handlers.DelUser)
	r.GET("/getAllUser", handlers.GetAllUser)

	// userInfo
	r.GET("/queryInfo", handlers.Query)
	r.POST("/updateInfo", handlers.Update)

	// edit
	r.GET("/dirTree", handlers.GetDirTree)
	r.GET("/fileContent", handlers.GetFile)
	r.POST("/createFile", handlers.CreateFile)
	r.POST("/createDir", handlers.CreateDir)
	r.POST("/saveFile", handlers.SaveFile)
	r.DELETE("/delAll", handlers.DelAll)

	// file
	r.POST("/uploadFile", handlers.UploadFile)

	if err := http.ListenAndServe(":8080", r); err != nil {
		println(err)
	}
}
