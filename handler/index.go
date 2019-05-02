package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/weifuchuan/GKD/session"
)

type Any = interface{}

func Index(c *gin.Context) {
	session := session.DefaultSession(c)

	account, _ := session.Get("loginAccount")

	c.HTML(200, "index.html", map[string]Any{
		"account": account,
	})
}
