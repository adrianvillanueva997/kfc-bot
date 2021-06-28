package main

import (
	"adrianvillanueva997/kfcbot/src/commands"
	"log"
	"os"

	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/joho/godotenv"
)

func main() {
	load_env()
	token := os.Getenv("bot_key")
	if token == "" {
		log.Fatalln("Bot key not valid!")
	}
	bot.Run(token, &commands.Bot{}, func(c *bot.Context) error {
		c.HasPrefix = bot.NewPrefix("!")
		return nil
	})
}

func load_env() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}
