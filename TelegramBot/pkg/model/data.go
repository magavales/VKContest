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

func (d *Data) parseData(message *tgbotapi.Message) (string, string, string) {
	strs := strings.Split(message.Text, "\n")
	service := strings.Split(strs[0], " ")
	login := strings.Split(strs[1], " ")
	password := strings.Split(strs[2], " ")

	return service[1], login[1], password[1]
}

func (d *Data) Add(message *tgbotapi.Message) {
	_, _, password := d.parseData(message)

	d.Password = password
}
