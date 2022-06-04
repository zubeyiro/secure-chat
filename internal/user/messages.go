package user

import (
	"errors"
	"fmt"
	"strings"

	"github.com/zubeyiro/secure-chat/internal/events"
	"github.com/zubeyiro/secure-chat/internal/message"
	"github.com/zubeyiro/secure-chat/internal/security"
)

func authMessage(str string) (*message.Message, error) {
	if len(str) < 3 {
		fmt.Println("Username must be at least 3 characters")
		fmt.Printf("Please login, enter your username: ")

		return nil, errors.New("invalid username")
	}
	return message.NewMessage(events.AUTH_REQUEST, str, security.ExportPublicKeyBase64(&privateKey.PublicKey)), nil
}

func chatMessage(str string) (*message.Message, error) {
	if !strings.Contains(str, ":") {
		fmt.Println("Please select recipient, type your message as recipient:message")

		return nil, errors.New("invalid message")
	}

	parts := strings.Split(str, ":")
	recipient := strings.Trim(parts[0], " ")
	content := strings.Trim(parts[1], " ")

	if len(recipient) < 3 || len(content) < 1 {
		fmt.Println("Please select valid user and your message must contain at least 1 character")
		fmt.Println("Please select recipient, type your message as recipient:message")

		return nil, errors.New("invalid message")
	}

	_, ok := users[recipient]

	if !ok {
		fmt.Println("Unknown user")
		fmt.Println("Please select recipient, type your message as recipient:message")

		return nil, errors.New("invalid user")
	}

	encrypted, err := security.Encrypt([]byte(content), users[recipient])

	if err != nil {
		fmt.Println("Error while encrypting the message")
		fmt.Println("Please select recipient, type your message as recipient:message")

		return nil, errors.New("error while encrypting the message")
	}

	return message.NewMessage(events.NEW_MESSAGE, recipient, encrypted), nil
}
