package main

import (
	"TelegramBot/pkg/telegram"
)

func main() {
	var bot telegram.TgBot
	bot.InitBot("6092617311:AAGTJ5iIZ_xyw7VC04gfCCAIi-NZbVFnBeI")
	bot.RunBot()
}
