package main

import (
	"Bishe/be/cmd/api/handlers"
	"Bishe/be/cmd/api/rpc"
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func Init() {
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.Default()
	
	// user
	r.POST("/login/", handlers.Login)
	r.POST("/register/", handlers.Register)

	// userInfo
	r.GET("/queryInfo/", handlers.Query)
	r.POST("/updateInfo/", handlers.Update)

	if err := http.ListenAndServe(":8080", r); err != nil {
		println(err)
	}
}
