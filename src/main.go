package main

import (
	"log"
	"os"

	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/joho/godotenv"
)

type Bot struct {
	Ctx *bot.Context
}

func main() {
	load_env()
	token := os.Getenv("bot_key")
	if token == "" {
		log.Fatalln("Bot key not valid!")
	}
	bot.Run(token, &Bot{}, func(c *bot.Context) error {
		c.HasPrefix = bot.NewPrefix("!")
		return nil
	})
}

func (b *Bot) Ping(*gateway.MessageCreateEvent) (string, error) {
	return "Pong!", nil
}

func load_env() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}
