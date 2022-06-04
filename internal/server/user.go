package server

import (
	"bufio"
	"crypto/rsa"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/zubeyiro/secure-chat/internal/events"
	"github.com/zubeyiro/secure-chat/internal/message"
	"github.com/zubeyiro/secure-chat/internal/security"
)

type User struct {
	id            string
	name          string
	conn          net.Conn
	publicKey     *rsa.PublicKey
	reader        *bufio.Reader
	writer        *bufio.Writer
	outgoing      chan string
	connectedOn   int64
	lastHeartbeat int64
}

var users map[string]*User    // mutex
var userMap map[string]string // this is map for mapping users with [name]ip:port

func getUserByName(name string) *User {
	return users[userMap[name]]
}

func newUser(name string, publicKey string, conn net.Conn) *User {
	return &User{
		id:            conn.RemoteAddr().String(),
		name:          name,
		conn:          conn,
		publicKey:     security.ParsePublicKeyFromBase64(publicKey),
		reader:        bufio.NewReader(conn),
		writer:        bufio.NewWriter(conn),
		outgoing:      make(chan string),
		connectedOn:   time.Now().Unix(),
		lastHeartbeat: time.Now().Unix(),
	}
}

func (user *User) sendMessage(msg *message.Message) {
	user.outgoing <- msg.Serialize()
}

func (user *User) read() {
	for {
		str, err := user.reader.ReadString('\n')

		if err != nil {
			break
		}
		msg := message.Deserialize(str)
		fmt.Println("USER")
		fmt.Println(str)

		switch msg.Command {
		case events.NEW_MESSAGE:
			user.relayMessageToUser(msg.Owner, msg.Message)
		}
	}

	user.logout()
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

	user.logout()
}

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

// inform other users about new joiner
func (user *User) loggedIn() {
	msg := message.NewMessage(events.NEW_USER_CONNECTED, user.name, security.ExportPublicKeyBase64(user.publicKey))

	for _, v := range users {
		if v.id == user.id {
			continue
		}

		v.sendMessage(msg)
	}
}

// inform other users about user left
func (user *User) loggedOut() {
	msg := message.NewMessage(events.USER_DISCONNECTED, user.name, "")

	for _, v := range users {
		if v.id == user.id {
			continue
		}

		v.sendMessage(msg)
	}
}

// share user list with new joiner
func (user *User) sendUserList() {
	if len(users) == 0 {
		return
	}

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

func (user *User) relayMessageToUser(who string, msg string) {
	getUserByName(who).sendMessage(message.NewMessage(events.NEW_MESSAGE, user.name, msg))
}
