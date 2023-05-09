package telegram

import (
	"TelegramBot/pkg/database"
	"TelegramBot/pkg/dialog"
	"TelegramBot/pkg/telegram/handler"
	"TelegramBot/pkg/telegram/handler/del"
	"TelegramBot/pkg/telegram/handler/get"
	"TelegramBot/pkg/telegram/handler/set"
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"time"
)

type TgBot struct {
	Bot     *tgbotapi.BotAPI
	Uconf   tgbotapi.UpdateConfig
	Updates tgbotapi.UpdatesChannel
}

var dao database.Database

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
	dao = initDatabase()
	for upd := range tgb.Updates {
		dao.StatConn()
		if upd.Message == nil {
			continue
		}

		if upd.Message.IsCommand() {
			tgb.handleCommand(upd.Message)
		} else {
			tgb.handleMessage(upd.Message)
		}
	}
}

func initDatabase() database.Database {
	var db database.Database
	db.Connect()
	return db
}

func (tgb *TgBot) handleCommand(message *tgbotapi.Message) (string, error) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	go tgb.delayedDelete(message.Chat.ID, message.MessageID)
	switch message.Command() {
	case "start":
		msg.Text = fmt.Sprintf("Hello! I'm BotStorage. I can save your login and password from anything services, if you want.\n\n\n%s, you are so cute!", message.From.UserName)
		tgb.warningForExplosions(message.Chat.ID)
	case "set":
		msg.Text = "Ok. You can trust me your personal data. Write name of service:"
		dialog.SetState(message.Chat.ID, dialog.State{
			Type:  dialog.Set,
			Name:  dialog.Service,
			Value: "",
		})
		tgb.warningForExplosions(message.Chat.ID)
	case "get":
		msg.Text = "Ok. Write name of service:"
		dialog.SetState(message.Chat.ID, dialog.State{
			Type:  dialog.Get,
			Name:  dialog.Service,
			Value: "",
		})
		tgb.warningForExplosions(message.Chat.ID)
	case "del":
		msg.Text = "Ok. Write name of service:"
		dialog.SetState(message.Chat.ID, dialog.State{
			Type:  dialog.Delete,
			Name:  dialog.Service,
			Value: "",
		})
		tgb.warningForExplosions(message.Chat.ID)
	default:
		msg.Text = "I don't know this command.\nYou must chose command from menu."
	}
	send, err := tgb.Bot.Send(msg)
	go tgb.delayedDelete(message.Chat.ID, send.MessageID)

	return message.Command(), err
}

func (tgb *TgBot) handleMessage(requestMessage *tgbotapi.Message) error {
	var (
		responseMessage tgbotapi.MessageConfig
		err             error
	)

	go tgb.delayedDelete(requestMessage.Chat.ID, requestMessage.MessageID)
	state := dialog.GetState(requestMessage.Chat.ID)
	switch state.Type {
	case dialog.Set:
		responseMessage = set.HandleDialogSetMessage(requestMessage, &state, dao)
	case dialog.Get:
		responseMessage = get.HandleDialogGetMessage(requestMessage, &state, dao)
	case dialog.Delete:
		responseMessage = del.HandleDialogDeleteMessage(requestMessage, &state, dao)
	default:
		responseMessage = handler.HandleDialogDefaultResponse(requestMessage)
	}
	send, err := tgb.Bot.Send(responseMessage)
	go tgb.delayedDelete(requestMessage.Chat.ID, send.MessageID)

	return err
}

func (tgb *TgBot) delayedDelete(chatId int64, messageId int) {
	config := tgbotapi.DeleteMessageConfig{
		ChatID:    chatId,
		MessageID: messageId,
	}
	time.Sleep(time.Second * 30)
	tgb.Bot.DeleteMessage(config)
}

func (tgb *TgBot) warningForExplosions(chatId int64) {
	response := tgbotapi.NewMessage(chatId, "🧨 Warning! All messages in this chat explode after 30 seconds!")
	send, _ := tgb.Bot.Send(response)
	go tgb.delayedDelete(chatId, send.MessageID)
}
