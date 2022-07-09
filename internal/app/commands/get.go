package commands

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong arg", args)
		return
	}

	product, err := c.productService.Get(idx)
	if err != nil {
		log.Printf("failed to get product with idx %d: %v\n", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		product.Title,
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("failed to send a message %v: %s", msg, err)
	}
	log.Println("message is sent")
}
