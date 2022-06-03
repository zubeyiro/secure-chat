package server

import (
	"fmt"
	"strings"

	"github.com/zubeyiro/secure-chat/internal/events"
	"github.com/zubeyiro/secure-chat/internal/message"
	"github.com/zubeyiro/secure-chat/internal/security"
)

func (user *User) login() {
	users[user.id] = user
	userMap[user.name] = user.id

	go user.read()
	go user.write()

	fmt.Printf("%s logged in\n", user.name)
	fmt.Printf("Logged in users: %d\n", len(users))

	user.sendUserList()
	user.loggedIn()
}

func (user *User) logout() {
	fmt.Printf("%s logged out\n", user.name)
	delete(users, user.id)
	delete(userMap, user.name)
	fmt.Printf("Logged in users: %d\n", len(users))

	user.loggedOut()
}

func (user *User) loggedIn() {
	msg := message.NewMessage(events.NEW_USER_CONNECTED, user.name, security.ExportPublicKeyBase64(user.publicKey))

	for _, v := range users {
		if v.id == user.id {
			continue
		}

		v.sendMessage(msg)
	}
}

func (user *User) loggedOut() {
	msg := message.NewMessage(events.USER_DISCONNECTED, user.name, "")

	for _, v := range users {
		if v.id == user.id {
			continue
		}

		v.sendMessage(msg)
	}
}

func (user *User) sendUserList() {
	msgContent := []string{} // format: name,publicKey|name,publicKey

	for _, v := range users {
		if v.id == user.id {
			continue
		}

		msgContent = append(msgContent, fmt.Sprintf("%s,%s", v.name, security.ExportPublicKeyBase64(v.publicKey)))
	}

	msg := message.NewMessage(events.USER_LIST, "", strings.Join(msgContent, "|"))
	user.sendMessage(msg)
}

func (user *User) sendMessageToUser(who string, msg string) {
	getUserByName(who).sendMessage(message.NewMessage(events.NEW_MESSAGE, user.name, msg))
}
