package main

import (
	"flag"
	"log"
)

func main() {
	t, err := token()

	//telegramClient = telegram.New(token)

	//fetcher = fetcher.New(telegramClient)

	//processor = processor.New(telegramClient)

	//consumer.Start(fetcher, processor)
}

func getToken() string {
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
