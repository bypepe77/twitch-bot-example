package bot

func (b *Bot) GetCommandList() []*Command {
	return []*Command{
		{
			Name:        "ping",
			Description: "Pong!",
			Handler:     b.Ping(),
		},
	}
}
