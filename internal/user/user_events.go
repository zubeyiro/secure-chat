package user

import (
	"fmt"
	"strings"

	"github.com/zubeyiro/secure-chat/internal/message"
	"github.com/zubeyiro/secure-chat/internal/security"
)

func userListReceived(msg *message.Message) {
	if len(msg.Message) > 1 {
		for _, v := range strings.Split(msg.Message, "|") {
			user := strings.Split(v, ",")

			addUser(user[0], strings.Join(user[1:], ""))
		}
	}

	printLobby()
}

func newUserConnected(msg *message.Message) {
	fmt.Printf("%s joined\n", msg.Owner)
	addUser(msg.Owner, msg.Message)
	printLobby()
}

func userDisconnected(msg *message.Message) {
	fmt.Printf("%s left\n", msg.Owner)
	removeUser(msg.Owner)
	printLobby()
}

func addUser(name string, publicKey string) {
	users[name] = security.ParsePublicKeyFromBase64(publicKey)
}

func removeUser(name string) {
	delete(users, name)
}
func messageRecieved(msg *message.Message) {
	m, err := security.Decrypt(msg.Message, privateKey)

	if err != nil {
		fmt.Println("Error while decrypting the message")
		fmt.Println(err)
		fmt.Println(err.Error())

		return
	}

	fmt.Printf("%s: %s\n", msg.Owner, m)
}
