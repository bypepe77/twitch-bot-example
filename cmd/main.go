package main

import (
	"fmt"
	"os"

	"github.com/bypepe77/twitch-bot-test/internal/bot"
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error cargando el archivo .env")
	}

	client := twitch.NewClient(os.Getenv("CHANNEL"), os.Getenv("OAUTH_TOKEN"))
	bot := bot.NewBot("bypepe77 bot", client)
	bot.RegisterCommands()

	err = bot.Start()
	if err != nil {
		println(err.Error())
	}

}
