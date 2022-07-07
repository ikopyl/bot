package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		fmt.Sprintf("%s\n%s",
			"/help - help",
			"/list - list products",
		))
	c.bot.Send(msg)
}

func init() {
	registeredCommands["help"] = (*Commander).Help
}
