package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Config structure
type Config struct {
	TelegramBotToken string
	debug            bool
}

func main() {
	var ConfigPath string
	configuration := Config{}

	flag.StringVar(&ConfigPath, "c", os.Getenv("HOME")+"/.config/EtherealBot/config.json", "determine what config to use")
	flag.Parse()

	file, _ := os.Open(ConfigPath)
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&configuration); err != nil {
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
		if update.Message == nil && !update.Message.IsCommand() {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "help":
			msg.Text = "This is help"
		case "host":
			msg.Text, _ = os.Hostname()
		case "randomshit":
			msg.Text, err = GetRandomShittyImage(update.Message.Text)
			if err != nil {
				log.Printf("Something went wrong: %s", err)
			}
		default:
			msg.Text = "Sorry, i cant understand..."
		}

		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
