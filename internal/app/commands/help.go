package commands

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		fmt.Sprintf("%s\n%s",
			"/help - help",
			"/list - list products",
		))

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("failed to send a message %v: %s", msg, err)
	}
	log.Println("message is sent")
}
