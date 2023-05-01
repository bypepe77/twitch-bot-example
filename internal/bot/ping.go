package bot

import "github.com/gempir/go-twitch-irc/v2"

func (b *Bot) Ping() func(message twitch.PrivateMessage, bot *Bot) {
	return func(message twitch.PrivateMessage, bot *Bot) {
		b.client.Say(message.Channel, "Pong!")
	}
}
