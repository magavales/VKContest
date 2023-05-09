package model

import (
	"github.com/Syfaro/telegram-bot-api"
	"strings"
)

type UserCredentials struct {
	ChatID   int64
	Service  string
	Login    string
	Password string
}

func SplitString(str string) string {
	newstr := strings.Split(str, " ")

	return newstr[1]
}

func (d *UserCredentials) GetFromMessage(message *tgbotapi.Message) {
	strs := strings.Split(message.Text, "\n")
	d.ChatID = message.Chat.ID
	d.Service = SplitString(strs[0])
	d.Login = SplitString(strs[1])
	d.Password = SplitString(strs[2])
}

func (d *UserCredentials) ParseColumns(value []interface{}) {
	d.Service = value[2].(string)
	d.Login = value[3].(string)
	d.Password = value[4].(string)
}
