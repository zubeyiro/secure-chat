package user

import (
	"crypto/rsa"
	"fmt"
	"strings"

	"github.com/zubeyiro/secure-chat/internal/configuration"
	"github.com/zubeyiro/secure-chat/internal/security"
)

var config *configuration.Configuration

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func Start() {
	setup()
	connectServer()
}

func setup() {
	config = configuration.GetConfig()
	privateKey, publicKey = security.GenerateKeyPair()
}

func printLobby() {
	if len(users) > 0 {
		keys := make([]string, 0, len(users))

		for k := range users {
			keys = append(keys, k)
		}

		fmt.Printf("Users (%d): %s\n", len(users), strings.Join(keys, ", "))
		fmt.Println("Please select recipient, type your message as recipient:message")
	}
}
