package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ikopyl/bot/internal/app/commands"
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
	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		commander.HandleUpdate(update)
	}

}
