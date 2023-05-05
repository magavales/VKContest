package auth

import (
	"VK/pkg/model"
	"crypto/sha1"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
	"log"
)

type UserPassword struct{}

func (up *UserPassword) PasswordVerification(user model.User, expectedPassword string) bool {
	pwd := hex.EncodeToString(pbkdf2.Key([]byte(user.Password), []byte(user.Salt), 8192, 32, sha1.New))
	if expectedPassword == pwd {
		return true
	} else {
		log.Println(pwd)
	}
	return false
}
