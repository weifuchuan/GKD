package main

import (
	"GKD/handler"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

// server: responce = handle(request)

func main()  {
	server:=gin.Default()

	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("GKD", store))

	server.LoadHTMLGlob("webapp/*")

	// /static/js/jjj.js
	server.Static("/static", "./webapp")

	server.GET("/login", handler.Login)
	server.POST("/doLogin",handler.DoLogin)

	server.Run(":8080")
}
