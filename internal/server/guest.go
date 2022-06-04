package server

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"github.com/zubeyiro/secure-chat/internal/events"
	"github.com/zubeyiro/secure-chat/internal/message"
)

type Guest struct {
	id            string
	conn          net.Conn
	reader        *bufio.Reader
	writer        *bufio.Writer
	outgoing      chan string
	connectedOn   int64
	lastHeartbeat int64
}

var guests map[string]*Guest

func newGuest(conn net.Conn) *Guest {
	return &Guest{
		id:            conn.RemoteAddr().String(),
		conn:          conn,
		reader:        bufio.NewReader(conn),
		writer:        bufio.NewWriter(conn),
		outgoing:      make(chan string),
		connectedOn:   time.Now().Unix(),
		lastHeartbeat: time.Now().Unix(),
	}
}

func (guest *Guest) join() {
	guests[guest.id] = guest

	go guest.read()
	go guest.write()

	fmt.Printf("%s connected, waiting for auth\n", guest.id)
}

func (guest *Guest) sendMessage(msg *message.Message) {
	guest.outgoing <- msg.Serialize()
}

func (guest *Guest) read() {
	for {
		str, err := guest.reader.ReadString('\n')

		if err != nil {
			fmt.Printf("%s - %s\n", guest.id, err.Error())

			break
		}
		msg := message.Deserialize(str)
		fmt.Println("GUEST")
		fmt.Println(str)

		switch msg.Command {
		case events.AUTH_REQUEST:
			if auth(msg.Owner) {
				guest.sendMessage(message.NewMessage(events.AUTH_SUCCEEDED, "", msg.Owner))
				newUser(msg.Owner, msg.Message, guest.conn).login()
				guest.remove()
				goto DONE
			} else {
				guest.sendMessage(message.NewMessage(events.AUTH_FAILED, "", ""))
			}
		}
	}

DONE:
	guest.outgoing <- "STOP"

	guest.remove()
}

func (guest *Guest) write() {
	for str := range guest.outgoing {

		if str == "STOP" {
			goto DONE
		}

		_, err := guest.writer.WriteString(str)

		if err != nil {
			fmt.Println(err)
			guest.remove()

			break
		}

		err = guest.writer.Flush()

		if err != nil {
			fmt.Println(err)
			guest.remove()

			break
		}
	}

DONE:

	guest.remove()
}

func (guest *Guest) remove() {
	delete(guests, guest.id)
}
