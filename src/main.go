package main

import (
	"adrianvillanueva997/kfcbot/src/commands"
	"log"
	"os"

	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	token := os.Getenv("bot_key")
	if token == "" {
		log.Fatalln("Bot key not valid!")
	}
	bot.Run(token, &commands.Bot{}, func(c *bot.Context) error {
		c.HasPrefix = bot.NewPrefix("!")
		return nil
	})
}
