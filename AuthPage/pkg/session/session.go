package session

import (
	"github.com/google/uuid"
	"net/http"
	"time"
)

type SessionCookie struct {
	Cookie http.Cookie
}

func (s *SessionCookie) CreateCookie() {
	s.Cookie = http.Cookie{
		Name:     "Session",
		Value:    uuid.New().String(),
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 6),
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true,
	}
}
