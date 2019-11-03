package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Lenivaya/EtherealBot/modules/media/gifs"
	"github.com/Lenivaya/EtherealBot/modules/media/images"
	"github.com/Lenivaya/EtherealBot/modules/users"
	telebot "gopkg.in/tucnak/telebot.v2"
)

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
		admin := users.CheckAdmin(adminlist, message.Sender.Username)

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

	bot.Handle("/gif", func(message *telebot.Message) {
		gifurl, _ := gifs.GetRandomGif(message.Text)
		gif := &telebot.Video{File: telebot.FromURL(gifurl)}

		bot.Send(message.Chat, gif, &telebot.SendOptions{
			ReplyTo: message,
		})
	})

	bot.Handle("/randomshit", func(message *telebot.Message) {
		randomshiturl, err := images.GetRandomShittyImage(message.Text)
		if err != nil {
			log.Printf("Something went wrong: %s", err)
		}

		bot.Send(message.Chat, randomshiturl, &telebot.SendOptions{
			ReplyTo: message,
		})
	})

	bot.Handle("/wallhaven", func(message *telebot.Message) {
		wallurl, err := images.GetWallFromWallhaven()
		if err != nil {
			log.Printf("Something went wrong: %s", err)
		}

		bot.Send(message.Chat, wallurl, &telebot.SendOptions{
			ReplyTo: message,
		})
	})

	bot.Start()
}
