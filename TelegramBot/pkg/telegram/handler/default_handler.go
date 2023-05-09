package handler

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func HandleDialogDefaultResponse(message *tgbotapi.Message) tgbotapi.MessageConfig {
	var (
		responseMessage = tgbotapi.NewMessage(message.Chat.ID, "")
	)
	responseMessage.Text = "I don't know this command.\nYou must chose command from menu."
	return responseMessage
}
