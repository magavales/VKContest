package model

import (
	"VK/pkg/session"
	"net/http"
)

type CookieHolder struct {
	SessionHolder map[string]http.Cookie
}

var instance *CookieHolder = nil

func (c *CookieHolder) init() {
	c.SessionHolder = make(map[string]http.Cookie)
}

func GetCookieHolderInstance() *CookieHolder {
	if instance == nil {
		instance = new(CookieHolder)
		instance.init()
	}
	return instance
}

func (c *CookieHolder) Add(s session.SessionCookie) {
	c.SessionHolder[s.Cookie.Value] = s.Cookie
}

func (c *CookieHolder) Get(session string) http.Cookie {
	return c.SessionHolder[session]
}
