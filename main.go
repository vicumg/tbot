package main

import (
	"flag"
	"log"
)

func main() {

	token := mustToken()

	// tgClient = telegram.New(token)

	// fetcher = fetcher.New(tgClient)

	// processor = processor.New(tgClient)

	//consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"bot-token", "", "token for access to telegramm")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is empty")
	}

	return *token
}
