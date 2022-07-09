package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CommandData struct {
	Offset int `json:"offset"`
}

func (c *Commander) Default(inputMsg *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "You wrote: "+inputMsg.Text)
	// msg.ReplyToMessageID = update.Message.MessageID

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("failed to send a message %v: %s", msg, err)
	}
	log.Println("message is sent")
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {

	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {

		// args := strings.Split(update.CallbackQuery.Data, "_")

		parsedData := CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)

		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Parsed: %+v\n", parsedData),
			// fmt.Sprintf("Command: %s\n", args[0])+
			// 	fmt.Sprintf("Offset: %s\n", args[1]),
		)
		c.bot.Send(msg)
		return
	}

	if update.Message == nil {
		return
	}

	switch strings.ToLower(update.Message.Command()) {
	case "get":
		c.Get(update.Message)
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	default:
		c.Default(update.Message)
	}

}
