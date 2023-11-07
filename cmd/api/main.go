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
	user := r.Group("/user")
	user.POST("/login/", handlers.Login)
	user.POST("/register/", handlers.Register)

	if err := http.ListenAndServe(":8080", r); err != nil {
		println(err)
	}
}