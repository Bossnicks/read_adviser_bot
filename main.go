package main

import (
	"flag"
	"log"

	tgClient "github.com/Bossnicks/read_adviser_bot/clients/telegram"
	event_consumer "github.com/Bossnicks/read_adviser_bot/consumer/event-consumer"
	"github.com/Bossnicks/read_adviser_bot/events/telegram"
	"github.com/Bossnicks/read_adviser_bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

	//fetcher = fetcher.New(telegramClient)

	//processor = processor.New(telegramClient)

	//consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
