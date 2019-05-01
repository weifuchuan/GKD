package session

import (
"github.com/gin-gonic/gin"
"github.com/chanxuehong/sid"
"github.com/clusterit/orca/timebuffer"
)

const sessionCookieName = "sessionZH3HS"

type H = map[string]interface{}
type Any = interface{}

func Middleware(c *gin.Context) {
	coo, err := c.Cookie(sessionCookieName)
	if err != nil {
		coo = sid.New()
		c.SetCookie(sessionCookieName, coo, 0, "", "", false, true)
	}
	c.Set(sessionCookieName, coo)
	if timebuffer.Get(coo) == nil {
		timebuffer.Put(coo, H{}, 60*60)
	}else{
		timebuffer.Put(coo, timebuffer.Get(coo), 60*60)
	}
	c.Next()
}

type Session interface {
	Set(key string, value Any)
	Get(key string) (value Any, ok bool)
}

type session struct {
	store H
}

func DefaultSession(c *gin.Context) Session {
	coo, _ := c.Get(sessionCookieName)
	store := timebuffer.Get(coo.(string))
	return &session{
		store: store.(H),
	}
}

func (s *session) Set(key string, value Any) {
	s.store[key] = value
}

func (s *session) Get(key string) (value Any, ok bool) {
	value, ok = s.store[key]
	return
}
