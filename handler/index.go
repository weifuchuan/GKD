package handler

import (
	"GKD/session"
	"github.com/gin-gonic/gin"
)

type Any = interface{}

func Index(c *gin.Context) {
	session := session.DefaultSession(c)

	account, _ := session.Get("loginAccount")

	c.HTML(200, "index.html", map[string]Any{
		"account": account,
	})
}
