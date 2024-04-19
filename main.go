package main

import (
	"flag"
	"log"
	"sc-music-snippeter/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient = telegram.New(tgBotHost, mustToken())

	//fetcher = fetcher.New()

	//processor = processor.New()

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String("tg-token", "", "access token for telegram bot")
	
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}