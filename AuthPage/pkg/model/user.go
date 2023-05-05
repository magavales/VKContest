package model

import (
	"encoding/base64"
	"log"
	"strings"
)

type User struct {
	ID       int64
	Username string
	Password string
	Salt     string
}

func (u *User) GetUserFromHeader(str string) {
	str = strings.Fields(str)[1]
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatal("error:", err)
	}
	idx := strings.IndexByte(string(data), ':')
	u.Username = string(data[:idx])
	u.Password = string(data[idx+1:])
}

func (u *User) ParseData(values []interface{}) {
	u.ID = values[0].(int64)
	u.Username = values[1].(string)
	u.Password = values[2].(string)
	u.Salt = values[3].(string)
}
