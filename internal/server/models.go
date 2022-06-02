package server

import (
	"bufio"
	"crypto/rsa"
	"net"
	"time"

	"github.com/zubeyiro/secure-chat/internal/message"
)

type User struct {
	isReady       bool // this is flag that represents we have all necessary information for user to start chatting
	name          string
	conn          net.Conn
	key           *rsa.PublicKey
	reader        *bufio.Reader
	writer        *bufio.Writer
	incoming      chan *message.Message
	outgoing      chan string
	connectedOn   int64
	lastHeartbeat int64
}

func newUser(conn net.Conn) *User {
	return &User{
		isReady:       false,
		name:          "",
		conn:          conn,
		key:           nil,
		reader:        bufio.NewReader(conn),
		writer:        bufio.NewWriter(conn),
		incoming:      make(chan *message.Message),
		outgoing:      make(chan string),
		connectedOn:   time.Now().Unix(),
		lastHeartbeat: time.Now().Unix(),
	}
}
