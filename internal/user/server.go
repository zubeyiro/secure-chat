package user

import (
	"bufio"
	"crypto/rsa"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/zubeyiro/secure-chat/internal/events"
	"github.com/zubeyiro/secure-chat/internal/message"
)

var conn net.Conn
var wg sync.WaitGroup
var users map[string]*rsa.PublicKey
var name string
var CURRENT_STATE string

func connectServer() {
	wg.Add(1)

	users = make(map[string]*rsa.PublicKey)

	connection, err := net.Dial(config.Server.Type, config.Server.Port)
	if err != nil {
		fmt.Println(err)
	}

	conn = connection
	CURRENT_STATE = STATE_AUTH

	go Read()
	go Write()

	fmt.Printf("Please login, enter your username: ")

	wg.Wait()
}

func Read() {
	reader := bufio.NewReader(conn)

	for {
		str, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("DISCONNNECTED")
			wg.Done()

			return
		}

		msg := message.Deserialize(str)

		switch msg.Command {
		case events.AUTH_FAILED:
			fmt.Println("Auth failed")
			fmt.Printf("Please login, enter your username: ")
		case events.AUTH_SUCCEEDED:
			name = msg.Message
			fmt.Printf("Successfully logged in, welcome %s\n", name)
			CURRENT_STATE = STATE_CHAT
		case events.USER_LIST:
			userListReceived(msg)
		case events.NEW_USER_CONNECTED:
			newUserConnected(msg)
		case events.USER_DISCONNECTED:
			userDisconnected(msg)
		case events.NEW_MESSAGE:
			messageRecieved(msg)
		}
	}
}

func Write() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(conn)

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var msg *message.Message

		switch CURRENT_STATE {
		case STATE_AUTH:
			msg, err = authMessage(str)
		case STATE_CHAT:
			msg, err = chatMessage(str)
		}

		if err == nil {
			_, err = writer.WriteString(msg.Serialize())

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		err = writer.Flush()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
