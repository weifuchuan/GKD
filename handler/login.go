package handler

import (
	"GKD/db"
	"GKD/session"
	"github.com/gin-gonic/gin"
)

type any = interface {
}

func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func DoLogin(c *gin.Context) {
	session := session.DefaultSession(c)
	username := c.PostForm("username")
	password := c.PostForm("password")
	account := new(db.Account)
	ok, _ := db.Engine.Where("username=?", username).Get(account)

	ret := make(map[string]any)

	if ok {
		if password == account.Password {
			session.Set("loginAccount", account)
			ret["state"] = "ok"
		} else {
			ret["state"] = "fail"
			ret["msg"] = "密码错误"
		}
	} else {
		ret["state"] = "fail"
		ret["msg"] = "用户不存在"
	}


	c.JSON(200, ret)
}
