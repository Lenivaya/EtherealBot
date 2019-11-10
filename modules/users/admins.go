package users

import (
	"fmt"

	"gopkg.in/tucnak/telebot.v2"
)

func CheckAdmin(adminlist []telebot.ChatMember, username string) bool {
	for i := range adminlist {
		if adminlist[i].User.Username == username {
			return true
		}
	}
	return false
}

func ListAdmins(adminlist []telebot.ChatMember) string {
	var admins string
	for i := range adminlist {
		admins += fmt.Sprintf("- %s (%d)\n", adminlist[i].User.Username, adminlist[i].User.ID)
	}
	return admins
}
