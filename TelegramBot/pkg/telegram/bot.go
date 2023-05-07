package telegram

import (
	"TelegramBot/pkg/database"
	"TelegramBot/pkg/model"
	"errors"
	"fmt"
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
	var (
		lastCommand string
		db          database.Database
	)
	db.Connect()
	for upd := range tgb.Updates {
		if upd.Message == nil {
			continue
		}

		if upd.Message.IsCommand() {
			lastCommand = tgb.handleCommand(upd.Message)
		} else {
			err := tgb.handleMessage(upd.Message, lastCommand, db)
			if err == nil {
				lastCommand = ""
			}
		}
	}
}

func (tgb *TgBot) handleCommand(message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	switch message.Command() {
	case "start":
		msg.Text = fmt.Sprintf("Hello! I'm BotStorage. I can save your login and password from anything services, if you want.\n\n\n%s, you are so cute!", message.From.UserName)
		tgb.Bot.Send(msg)
	case "set":
		msg.Text = "Ok. You can trust me your personal data. PLease use this format:\nService: Telegram\nLogin: user\nPassword: 123"
		tgb.Bot.Send(msg)
	case "get":
		msg.Text = "Ok. Write name of service:"
		tgb.Bot.Send(msg)
	case "del":
		msg.Text = "Ok. Write name of service:"
		tgb.Bot.Send(msg)
	default:
		msg.Text = "I don't know this command.\nYou must chose command from menu."
		tgb.Bot.Send(msg)
	}

	return message.Command()
}

func (tgb *TgBot) handleMessage(message *tgbotapi.Message, command string, db database.Database) error {
	var (
		data model.Data
		err  error
	)
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	switch command {
	case "set":
		if checkSetData(message) {
			data.Add(message)
			db.DataEntity.AddData(db.Pool, data)
			msg.Text = "Ok. Your data has been saved."
			tgb.Bot.Send(msg)
			return nil
		} else {
			msg.Text = "Uncorrected data!"
			tgb.Bot.Send(msg)
			err = errors.New("uncorrected data")
			return err
		}
	case "get":
		if strings.Contains(message.Text, "Service ") || strings.Contains(message.Text, "Service: ") || strings.Contains(message.Text, "service ") || strings.Contains(message.Text, "service: ") {
			service := model.SplitData(message.Text)
			data, err = db.DataEntity.GetData(db.Pool, message.Chat.ID, service)
			if err != nil {
				msg.Text = "I think, you send uncorrected name of service!"
				tgb.Bot.Send(msg)
				return err
			}
			msg.Text = fmt.Sprintf("Service: %s\nLogin: %s\nPassword: %s\n", data.Service, data.Login, data.Password)
			tgb.Bot.Send(msg)
			return nil
		} else {
			data, err = db.DataEntity.GetData(db.Pool, message.Chat.ID, message.Text)
			if err != nil {
				msg.Text = "I think, you send uncorrected name of service!"
				tgb.Bot.Send(msg)
				return err
			} else {
				msg.Text = fmt.Sprintf("Service: %s\nLogin: %s\nPassword: %s\n", data.Service, data.Login, data.Password)
				tgb.Bot.Send(msg)
				return nil
			}
		}
	case "del":
		service := model.SplitData(message.Text)
		if strings.Contains(message.Text, "Service ") || strings.Contains(message.Text, "Service: ") || strings.Contains(message.Text, "service ") || strings.Contains(message.Text, "service: ") {
			data, err = db.DataEntity.GetData(db.Pool, message.Chat.ID, service)
			if err != nil {
				msg.Text = "I think, you send uncorrected name of service!"
				tgb.Bot.Send(msg)
				return err
			} else {
				msg.Text = fmt.Sprintf("I delete your data about %s", service)
				tgb.Bot.Send(msg)
				return nil
			}
		} else {
			data, err = db.DataEntity.GetData(db.Pool, message.Chat.ID, message.Text)
			if err != nil {
				msg.Text = "I think, you send uncorrected name of service!"
				tgb.Bot.Send(msg)
				return err
			} else {
				msg.Text = fmt.Sprintf("I delete your data about %s", service)
				tgb.Bot.Send(msg)
				return nil
			}
		}
	default:
		msg.Text = "I don't know this command.\nYou must chose command from menu."
		tgb.Bot.Send(msg)
	}
	return nil
}

func checkSetData(message *tgbotapi.Message) bool {
	if !strings.Contains(message.Text, "Service ") && !strings.Contains(message.Text, "service ") && !strings.Contains(message.Text, "Service: ") && !strings.Contains(message.Text, "service: ") {
		return false
	}
	if !strings.Contains(message.Text, "Login ") && !strings.Contains(message.Text, "login ") && !strings.Contains(message.Text, "Login: ") && !strings.Contains(message.Text, "login: ") {
		return false
	}
	if !strings.Contains(message.Text, "Password ") && !strings.Contains(message.Text, "password ") && !strings.Contains(message.Text, "Password: ") && !strings.Contains(message.Text, "password: ") {
		return false
	}
	return true
}
