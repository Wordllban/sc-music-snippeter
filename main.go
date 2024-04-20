package main

import (
	"flag"
	"log"

	tg "sc-music-snippeter/clients/telegram"
	logger "sc-music-snippeter/lib/logger"
)

func main() {
	
	platform, token := mustArgs()

	switch platform {
		case "tg":
			tgClient := tg.New(token)
			tgClient.Updates()
		default:
			logger.Log("Platform does not supported")
			return
	}

}

func mustArgs() (string, string) {
	log.Println("Available platforms: tg")
	platform := flag.String("platform", "", "platform name: ")
	token := flag.String("token", "", "platform name: ")
	
	flag.Parse()

	if *platform == "" {
		log.Fatal("platform is not specified")
	}

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *platform, *token
}
/* 
func mustToken() string {
	token := flag.String("client-token", "", "access token for messanger client")
	
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
 */