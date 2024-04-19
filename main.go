package main

import (
	"flag"
	"log"
)

func main() {
	token := mustToken()

	// tgClient = telegram.New(token)

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