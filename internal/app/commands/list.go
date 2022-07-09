package commands

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here is the list of our products: \n\n"

	products := c.productService.List()
	for _, p := range products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CommandData{
		Offset: 21,
	})

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			// tgbotapi.NewInlineKeyboardButtonData("Next page", "list_10"),
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(serializedData)),
		),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("failed to send a message %v: %s", msg, err)
	}
	log.Println("message is sent")
}
