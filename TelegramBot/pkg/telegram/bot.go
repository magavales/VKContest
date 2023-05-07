package telegram

import (
	"TelegramBot/pkg/model"
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"strings"
)

type TgBot struct {
	Bot     *tgbotapi.BotAPI
	Uconf   tgbotapi.UpdateConfig
	Updates tgbotapi.UpdatesChannel
}

func (tgb *TgBot) InitBot(token string) {
	var err error
	tgb.Bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Unable to launch the bot: %s\n", err)
	}
	tgb.Bot.Debug = true
	tgb.Uconf = tgbotapi.NewUpdate(0)
	tgb.Uconf.Timeout = 30
	tgb.Updates, err = tgb.Bot.GetUpdatesChan(tgb.Uconf)
}

func (tgb *TgBot) RunBot() {
	var lastCommand string
	for upd := range tgb.Updates {
		if upd.Message == nil {
			continue
		}

		if upd.Message.IsCommand() {
			lastCommand = tgb.handleCommand(upd.Message)
		} else {
			tgb.handleMessage(upd.Message, lastCommand)
		}
	}
}

func (tgb *TgBot) handleCommand(message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	switch message.Command() {
	case "start":
		msg.Text = "Hello! I'm BotStorage. I can save your login and password from anything services, if you want."
		tgb.Bot.Send(msg)
	case "set":
		msg.Text = "Ok. You can trust me your personal data. PLease use this format:\nService: Telegram\nLogin: user\nPassword: 123"
		tgb.Bot.Send(msg)
	case "get":
		msg.Text = "Ok. Write your secret key:"
		tgb.Bot.Send(msg)
	default:
		msg.Text = "I don't know this command.\nYou must chose command from menu."
	}

	return message.Command()
}

func (tgb *TgBot) handleMessage(message *tgbotapi.Message, command string) {
	var data *model.Data
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	switch command {
	case "set":
		if checkSetData(message) {
			data.Add(message)
			msg.Text = "Ok. Your data has been saved."
			tgb.Bot.Send(msg)
		} else {
			msg.Text = "Uncorrected data!"
			tgb.Bot.Send(msg)
		}
	case "get":
		msg.Text = "Ok. Write your secret key:"
		tgb.Bot.Send(msg)
	default:
		msg.Text = "I don't know this command.\nYou must chose command from menu."
	}
}

func checkSetData(message *tgbotapi.Message) bool {
	if !strings.Contains(message.Text, "Service ") && !strings.Contains(message.Text, "service ") {
		return false
	}
	if !strings.Contains(message.Text, "Login ") && !strings.Contains(message.Text, "login ") {
		return false
	}
	if !strings.Contains(message.Text, "Password ") && !strings.Contains(message.Text, "password ") {
		return false
	}
	return true
}
