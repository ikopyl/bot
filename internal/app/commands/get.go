package commands

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong arg", args)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("successfully parsed argument: %v", arg),
	)
	c.bot.Send(msg)
}
