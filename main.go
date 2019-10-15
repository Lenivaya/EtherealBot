package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/pelletier/go-toml"
	telebot "gopkg.in/tucnak/telebot.v2"
)

type Telegram struct {
	TelegramBotToken string
	debug            bool
}

type Config struct {
	Telegram Telegram
}

func main() {
	configuration := GetConfig()

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  configuration.Telegram.TelegramBotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
	}

	bot.Handle("/hello", func(message *telebot.Message) {
		bot.Send(message.Sender, "hello world", telebot.ForceReply)
	})

	bot.Handle("/randomshit", func(message *telebot.Message) {
		randomshiturl, _ := GetRandomShittyImage(message.Text)
		bot.Send(message.Sender, randomshiturl, &telebot.SendOptions{
			ReplyTo: message,
		})
	})

	bot.Handle("/wallhaven", func(message *telebot.Message) {
		wallurl, _ := GetWallFromWallhaven()
		bot.Send(message.Sender, wallurl, &telebot.SendOptions{
			ReplyTo: message,
		})
	})

	bot.Start()
}

func GetConfig() *Config {
	var configuration Config
	var ConfigPath string

	flag.StringVar(&ConfigPath, "c", os.Getenv("HOME")+"/.config/EtherealBot/config.toml", "determine what config to use")
	flag.Parse()

	file, _ := ioutil.ReadFile(ConfigPath)
	toml.Unmarshal(file, &configuration)

	return &configuration
}
