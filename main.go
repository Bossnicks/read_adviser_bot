package main

import (
	"flag"
	"log"

	"github.com/Bossnicks/read_adviser_bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	telegramClient := telegram.New(tgBotHost, mustToken())

	//fetcher = fetcher.New(telegramClient)

	//processor = processor.New(telegramClient)

	//consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"default-token",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
