package main

import (
	"github.com/gin-gonic/gin"
	"github.com/weifuchuan/GKD/handler"
	"github.com/weifuchuan/GKD/session"
)

// server: responce = handle(request)
// responce = handle...(request) ->
// client -> req -> | -> | -> addCar

func main() {
	server := gin.Default()

	server.Use(session.Middleware)

	server.LoadHTMLGlob("templates/*")

	// /static/js/jjj.js
	server.Static("/static", "resources")

	server.GET("/", handler.Index)
	server.GET("/login", handler.Login)
	server.POST("/doLogin", handler.DoLogin)

	server.Run(":8080")
}
