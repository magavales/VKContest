package del

import (
	"TelegramBot/pkg/database"
	"TelegramBot/pkg/dialog"
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func HandleDialogDeleteMessage(message *tgbotapi.Message, state *dialog.State, dao database.Database) tgbotapi.MessageConfig {
	var (
		responseMessage = tgbotapi.NewMessage(message.Chat.ID, "")
		nextState       dialog.State
	)

	state.Value = message.Text
	switch state.Name {
	case dialog.Service:
		var service string
		nextState, service = serviceHandler(state)
		err := dao.Access.DeleteUserCredentials(dao.Pool, message.Chat.ID, service)
		if err != nil {
			responseMessage.Text = fmt.Sprintf("Ohhh! Cannot find service, sorry :(\n")
		} else {
			responseMessage.Text = fmt.Sprintf("Removed creds for %s\n", service)
		}
	default:
		nextState = defaultHandler(state, &responseMessage)
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

func defaultHandler(state *dialog.State, responseMessage *tgbotapi.MessageConfig) dialog.State {
	nextState := dialog.State{
		Type:      dialog.End,
		Value:     "",
		PrevState: state,
	}
	responseMessage.Text = fmt.Sprintf("Sorry, but I cannot handle this! Let start again!\n")
	return nextState
}
