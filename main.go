// Package main provides
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Config struct {
	TelegramBotToken string
	debug            bool
}

func main() {
	var config string
	flag.StringVar(&config, "c", "config.json", "determine what config to use")
	flag.Parse()
	file, _ := os.Open(config)
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Using %s\n", configuration.TelegramBotToken)

	bot, err := tgbotapi.NewBotAPI(configuration.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = configuration.debug
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		var reply string
		if update.Message == nil {
			continue
		}
		switch update.Message.Command() {
		case "help":
			reply = "This is help"
		case "@ethereality_bot":
			reply = "Hello"
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}
