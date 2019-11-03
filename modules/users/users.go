package users

import "gopkg.in/tucnak/telebot.v2"

func CheckAdmin(adminlist []telebot.ChatMember, username string) bool {
	for i := range adminlist {
		if adminlist[i].User.Username == username {
			return true
		}
	}
	return false
}
