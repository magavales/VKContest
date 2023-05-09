package get

import (
	"TelegramBot/pkg/database"
	"TelegramBot/pkg/dialog"
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func HandleDialogGetMessage(message *tgbotapi.Message, state *dialog.State, dao database.Database) tgbotapi.MessageConfig {
	var (
		responseMessage = tgbotapi.NewMessage(message.Chat.ID, "")
		nextState       dialog.State
	)

	state.Value = message.Text
	switch state.Name {
	case dialog.Service:
		var service string
		nextState, service = serviceHandler(state)
		credentials, err := dao.Access.GetUserCredentials(dao.Pool, message.Chat.ID, service)
		if err != nil {
			responseMessage.Text = fmt.Sprintf("Ohhh! Cannot find service, sorry :(\n")
		} else {
			login := credentials.Login
			password := credentials.Password
			responseMessage.Text = fmt.Sprintf("Your creds for %s \nLogin: %s\nPassword: %s\n\nRight? I save it!", service, login, password)
		}
	default:
		nextState = defaultHandler(state, &responseMessage, nextState)
	}
	dialog.SetState(message.Chat.ID, nextState)

	return responseMessage
}

func serviceHandler(state *dialog.State) (dialog.State, string) {
	nextState := dialog.State{
		Type:      dialog.Set,
		Name:      dialog.Login,
		Value:     "",
		PrevState: state,
	}
	service := state.Value
	return nextState, service
}

func defaultHandler(state *dialog.State, responseMessage *tgbotapi.MessageConfig, nextState dialog.State) dialog.State {
	nextState = dialog.State{
		Type:      dialog.End,
		Value:     "",
		PrevState: state,
	}
	responseMessage.Text = fmt.Sprintf("Sorry, but I cannot handle this! Let start again!\n")
	return nextState
}
