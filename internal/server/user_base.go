package server

import (
	"bufio"
	"crypto/rsa"
	"fmt"
	"net"
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
	incoming      chan *message.Message
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
		incoming:      make(chan *message.Message),
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

		switch msg.Command {
		case events.NEW_MESSAGE:
			user.sendMessageToUser(msg.Owner, msg.Message)
		}
		user.incoming <- msg
	}
	close(user.incoming)

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
