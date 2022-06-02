package server

import (
	"fmt"
	"net"
	"os"

	"github.com/zubeyiro/secure-chat/internal/configuration"
	"github.com/zubeyiro/secure-chat/internal/events"
	"github.com/zubeyiro/secure-chat/internal/message"
)

var config *configuration.Configuration
var users map[string]*User // mutex

func Start() {
	config = configuration.GetConfig()
	users = make(map[string]*User)

	listener, err := net.Listen(config.Server.Type, config.Server.Port)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	defer listener.Close()

	fmt.Println("Server started listening on " + config.Server.Port)

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error: ", err)

			continue
		}

		joinUser(conn)
	}
}

func joinUser(conn net.Conn) {
	users[conn.RemoteAddr().String()] = newUser(conn)
	users[conn.RemoteAddr().String()].listen()

	msg := message.NewMessage(events.USER_INFO_REQUESTED, "", "").Serialize()
	users[conn.RemoteAddr().String()].outgoing <- msg

	fmt.Println("Requested info from client")
	fmt.Println(msg)
}

func (user *User) listen() {
	go user.read()
	go user.write()
}

func (user *User) read() {
	for {
		str, err := user.reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			break
		}
		message := message.Deserialize(str)
		user.incoming <- message
	}
	close(user.incoming)

	fmt.Println("Client disconnected")
}

func (user *User) write() {
	for str := range user.outgoing {
		_, err := user.writer.WriteString(str)
		if err != nil {
			fmt.Println(err)

			break
		}

		err = user.writer.Flush()

		if err != nil {
			fmt.Println(err)

			break
		}
	}

	fmt.Println("Closed client's write thread")
}

// job to ping user
// job to check users info

/*
- Server
	- Events
		- server_started
			- create public and private keys
		- user_connected
			- Add to users list
			- Update all other users
		- new_message
			- Relay message
		- user_list_request
			- Send list to requester
		- user_disconnected
*/
