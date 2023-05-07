package model

import (
	"github.com/Syfaro/telegram-bot-api"
	"strings"
)

type Data struct {
	ChatID   int64
	Service  string
	Login    string
	Password string
}

func SplitData(str string) string {
	newstr := strings.Split(str, " ")

	return newstr[1]
}

func (d *Data) Add(message *tgbotapi.Message) {
	strs := strings.Split(message.Text, "\n")
	d.ChatID = message.Chat.ID
	d.Service = SplitData(strs[0])
	d.Login = SplitData(strs[1])
	d.Password = SplitData(strs[2])
}

func (d *Data) ParseData(value []interface{}) {
	d.Service = value[2].(string)
	d.Login = value[3].(string)
	d.Password = value[4].(string)
}
