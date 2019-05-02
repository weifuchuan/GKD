package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/weifuchuan/GKD/db"
	"github.com/weifuchuan/GKD/kit"
	"github.com/weifuchuan/GKD/session"
)

func AuthManager(c *gin.Context) {
	session := session.DefaultSession(c)
	_acc, ok := session.Get("loginAccount")
	if ok {
		account := _acc.(db.Account)
		if kit.BitAt(account.Status, 2) == 1 {
			c.Next()
		} else {
			c.String(404, "")
		}
	} else {
		c.String(404, "")
	}
}
