package commands

import "github.com/diamondburned/arikawa/v2/gateway"

func (b *Bot) Ping(*gateway.MessageCreateEvent) (string, error) {
	return "Pong!", nil
}
