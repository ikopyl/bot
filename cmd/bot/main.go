package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ikopyl/bot/internal/service/product"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("Error loading .env file: %s", err)
	}
	token := os.Getenv("TELEGRAM_BOT_API_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	productService := product.NewService()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch strings.ToLower(update.Message.Command()) {
		case "help":
			helpCommand(bot, update.Message)
		case "list":
			listCommand(bot, update.Message, productService)
		default:
			defaultBehavior(bot, update.Message)
		}
	}

}

func helpCommand(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		fmt.Sprintf("%s\n%s",
			"/help - help",
			"/list - list products",
		))
	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, productService *product.Service) {
	outputMsgText := "Here is the list of our products: \n\n"

	products := productService.List()
	for _, p := range products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "You wrote: "+inputMsg.Text)
	// msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)

}
