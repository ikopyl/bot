package commands

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ikopyl/bot/internal/service/product"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {

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
