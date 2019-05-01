package main

import (
	"GKD/handler"
	"GKD/session"
	"github.com/gin-gonic/gin"
)

// server: responce = handle(request)

func main() {
	server := gin.Default()

	server.Use(session.Middleware)

	server.LoadHTMLGlob("webapp/*")

	// /static/js/jjj.js
	server.Static("/static", "./webapp")

	server.GET("/", handler.Index)
	server.GET("/login", handler.Login)
	server.POST("/doLogin", handler.DoLogin)

	server.Run(":8080")
}
