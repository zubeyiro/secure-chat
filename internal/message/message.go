package message

import (
	"fmt"
	"strings"
)

type Message struct {
	Command string
	Target  string
	Message string
}

func NewMessage(command, target, message string) *Message {
	return &Message{
		Command: command,
		Target:  target,
		Message: message,
	}
}

func (msg *Message) Serialize() string {
	return fmt.Sprintf("%s,%s,%s\n", msg.Command, msg.Target, msg.Message)
}

func Deserialize(msg string) *Message {
	parts := strings.Split(strings.TrimSuffix(msg, "\n"), ",")

	if len(parts) < 3 {
		panic(fmt.Sprintf("Invalid Message %s", string(msg)))
	}

	return &Message{
		Command: parts[0],
		Target:  parts[1],
		Message: strings.Join(parts[2:], ","),
	}
}
