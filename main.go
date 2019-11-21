package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/Lenivaya/EtherealBot/modules/media/gifs"
	"github.com/Lenivaya/EtherealBot/modules/media/images"
	"github.com/Lenivaya/EtherealBot/modules/media/wikipedia"
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

	bot.Handle("/wiki", func(message *telebot.Message) {
		wikipages, _ := wikipedia.WikipediaAPI(urlEncoded(trimCommand(message.Text)))

		for _, page := range wikipages {
			bot.Send(message.Chat, page, &telebot.SendOptions{
				ReplyTo: message,
			})

			time.Sleep(1 * time.Second)
		}
	})

	bot.Handle("/adminlist", func(message *telebot.Message) {
		adminlist, _ := bot.AdminsOf(message.Chat)
		admins := users.ListAdmins(adminlist)

		bot.Send(message.Chat, admins, &telebot.SendOptions{
			ReplyTo: message,
		})
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
		gifurl, _ := gifs.GetRandomGif(trimCommand(message.Text))
		gif := &telebot.Video{File: telebot.FromURL(gifurl)}

		bot.Send(message.Chat, gif, &telebot.SendOptions{
			ReplyTo: message,
		})
	})

	bot.Handle("/randomshit", func(message *telebot.Message) {
		randomshiturl, err := images.GetRandomShittyImage(trimCommand(message.Text))
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

func trimCommand(message string) string {
	var textDefault = "cat"
	args := strings.SplitN(message, " ", 2)

	for i, v := range args {
		if i == 1 {
			text := v
			return text
		}
	}
	return textDefault
}

func urlEncoded(str string) string {
	u, err := url.Parse(str)
	if err != nil {
		log.Printf("Something went wrong: %s", err)
	}
	return u.String()
}
