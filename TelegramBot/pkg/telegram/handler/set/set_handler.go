package set

import (
	"TelegramBot/pkg/database"
	"TelegramBot/pkg/dialog"
	"TelegramBot/pkg/model"
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
)

func HandleDialogSetMessage(message *tgbotapi.Message, state *dialog.State, dao database.Database) tgbotapi.MessageConfig {
	var (
		responseMessage = tgbotapi.NewMessage(message.Chat.ID, "")
		nextState       dialog.State
	)

	state.Value = message.Text

	switch state.Name {
	case dialog.Service:
		nextState = serviceHandler(state, &responseMessage)
	case dialog.Login:
		nextState = loginHandler(state, &responseMessage)
	case dialog.Password:
		var creds model.UserCredentials
		nextState, creds = passwordHandler(message, state, &responseMessage)
		dao.Access.AddUserCredentials(dao.Pool, creds)
	default:
		nextState = defaultHandler(state, &responseMessage)
	}
	dialog.SetState(message.Chat.ID, nextState)

	return responseMessage
}

func serviceHandler(state *dialog.State, responseMessage *tgbotapi.MessageConfig) dialog.State {
	nextState := dialog.State{
		Type:      dialog.Set,
		Name:      dialog.Login,
		Value:     "",
		PrevState: state,
	}
	responseMessage.Text = fmt.Sprintf("Service is: %s\nWrite login.", state.Value)
	return nextState
}

func loginHandler(state *dialog.State, responseMessage *tgbotapi.MessageConfig) dialog.State {
	nextState := dialog.State{
		Type:      dialog.Set,
		Name:      dialog.Password,
		Value:     "",
		PrevState: state,
	}
	responseMessage.Text = fmt.Sprintf("Login is: %s\nWrite password.", state.Value)
	return nextState
}

func passwordHandler(message *tgbotapi.Message, state *dialog.State, responseMessage *tgbotapi.MessageConfig) (dialog.State, model.UserCredentials) {
	nextState := dialog.State{
		Type:      dialog.End,
		Value:     "",
		PrevState: state,
	}
	password := state.Value
	login := state.PrevState.Value
	service := state.PrevState.PrevState.Value
	creds := model.UserCredentials{
		ChatID:   message.Chat.ID,
		Service:  service,
		Login:    login,
		Password: password,
	}
	responseMessage.Text = fmt.Sprintf("Ok! Your creds for %s \nLogin: %s\nPassword: %s\n\nRight? I save it!", service, login, password)
	return nextState, creds
}

func defaultHandler(state *dialog.State, responseMessage *tgbotapi.MessageConfig) dialog.State {
	nextState := dialog.State{
		Type:      dialog.End,
		Value:     "",
		PrevState: state,
	}
	responseMessage.Text = fmt.Sprintf("Sorry, but I cannot handle this! Let start again!\n")
	return nextState
}
