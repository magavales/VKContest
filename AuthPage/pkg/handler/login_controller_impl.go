package handler

import (
	"VK/pkg/auth"
	"VK/pkg/database"
	"VK/pkg/model"
	"VK/pkg/session"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var (
		db            database.Database
		sessionCookie session.SessionCookie
		userPassword  auth.UserPassword
		resp          Response
		user          model.User
	)
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		resp.SetStatusBadRequest()
		return
	}

	resp.rw = c.Writer
	user.GetUserFromHeader(authorization)
	if user.Username == "" {
		resp.SetStatusUnauthorized()
		return
	}

	db.Connect()
	userDB, err := db.Users.GetUser(db.Pool, user.Username)
	if err != nil {
		resp.SetStatusUnauthorized()
		return
	}

	if userPassword.PasswordVerification(user, userDB.Password) {
		sessionCookie.CreateCookie()
		model.GetCookieHolderInstance().Add(sessionCookie)
		resp.SetSessionCookie(sessionCookie.Cookie.String())
		resp.SetStatusOk()
	} else {
		resp.SetStatusUnauthorized()
	}
}
