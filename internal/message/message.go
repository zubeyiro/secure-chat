package message

import (
	"fmt"
	"strings"
)

type Message struct {
	Command string
	Owner   string
	Message string
}

func NewMessage(command, owner, message string) *Message {
	return &Message{
		Command: command,
		Owner:   owner,
		Message: message,
	}
}

func (msg *Message) Serialize() string {
	m := fmt.Sprintf("%s,%s,%s", msg.Command, msg.Owner, msg.Message)
	m = strings.Replace(m, "\n", "", -1)

	return fmt.Sprintf("%s\n", m)
}

func Deserialize(msg string) *Message {
	parts := strings.Split(strings.TrimSuffix(msg, "\n"), ",")

	if len(parts) < 3 {
		fmt.Printf("Invalid Message %s\n", string(msg))

		return nil
	}

	return &Message{
		Command: parts[0],
		Owner:   parts[1],
		Message: strings.Join(parts[2:], ","),
	}
}
