package bot

import (
	"github.com/gempir/go-twitch-irc/v2"
)

type Command struct {
	Name        string
	Description string
	Handler     func(message twitch.PrivateMessage, bot *Bot)
}

func (b *Bot) AddCommand(command *Command) {
	b.commands[command.Name] = command
}

func (b *Bot) RegisterCommands() {
	commands := b.GetCommandList()
	for _, command := range commands {
		b.AddCommand(command)
	}
}
