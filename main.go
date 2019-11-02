package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/pelletier/go-toml"
	telebot "gopkg.in/tucnak/telebot.v2"
)

type Telegram struct {
	TelegramBotToken string
}

type Config struct {
	Telegram Telegram
}

func main() {
	configuration := getConfig()

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  configuration.Telegram.TelegramBotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}
	log.Print("Bot successfully started")

	bot.Handle(telebot.OnText, func(message *telebot.Message) {
		// all the text messages that weren't
		// captured by existing handlers
	})

	bot.Handle("/pin", func(message *telebot.Message) {
		adminlist, _ := bot.AdminsOf(message.Chat)
		admin := checkAdmin(adminlist, message.Sender.Username)

		if admin {
			if message.ReplyTo == nil {
				bot.Pin(message)
			} else {
				bot.Pin(message.ReplyTo)
			}
		}
	})

	bot.Handle("/invitelink", func(message *telebot.Message) {
		link, err := bot.GetInviteLink(message.Chat)
		if err != nil {
			log.Printf("Something went wrong: %s", err)
		}
		bot.Send(message.Chat, link, &telebot.SendOptions{
			ReplyTo: message,
		})
	})

	bot.Handle("/id", func(message *telebot.Message) {
		id := fmt.Sprintf("Your id: %d\nChat id: %d", message.Sender.ID, message.Chat.ID)
		bot.Send(message.Chat, id, &telebot.SendOptions{
			ReplyTo: message,
		})
	})

	bot.Handle("/randomshit", func(message *telebot.Message) {
		randomshiturl, err := GetRandomShittyImage(message.Text)
		if err != nil {
			log.Printf("Something went wrong: %s", err)
		}
		bot.Send(message.Chat, randomshiturl, &telebot.SendOptions{
			ReplyTo: message,
		})
	})

	bot.Handle("/wallhaven", func(message *telebot.Message) {
		wallurl, err := GetWallFromWallhaven()
		if err != nil {
			log.Printf("Something went wrong: %s", err)
		}
		bot.Send(message.Chat, wallurl, &telebot.SendOptions{
			ReplyTo: message,
		})
	})

	bot.Start()
}

func getConfig() *Config {
	var configuration Config
	var ConfigPath string

	flag.StringVar(&ConfigPath, "c", os.Getenv("HOME")+"/.config/EtherealBot/config.toml", "determine what config to use")
	flag.Parse()

	file, _ := ioutil.ReadFile(ConfigPath)
	toml.Unmarshal(file, &configuration)

	return &configuration
}

func checkAdmin(adminlist []telebot.ChatMember, username string) bool {
	for i := range adminlist {
		if adminlist[i].User.Username == username {
			return true
		}
	}
	return false
}
