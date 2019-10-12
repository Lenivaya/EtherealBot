package main

import (
	"flag"
	"log"
	"os"

	"io/ioutil"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pelletier/go-toml"
)

type Telegram struct {
	TelegramBotToken string
	debug            bool
}

type Config struct {
	Telegram Telegram
}

func main() {
	var ConfigPath string
	configuration := Config{}

	flag.StringVar(&ConfigPath, "c", os.Getenv("HOME")+"/.config/EtherealBot/config.toml", "determine what config to use")
	flag.Parse()

	file, _ := ioutil.ReadFile(ConfigPath)
	toml.Unmarshal(file, &configuration)

	bot, err := tgbotapi.NewBotAPI(configuration.Telegram.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = configuration.Telegram.debug
	log.Printf("Using %s\n", configuration.Telegram.TelegramBotToken)
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
		case "wallhaven":
			msg.Text, err = GetWallFromWallhaven()
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
