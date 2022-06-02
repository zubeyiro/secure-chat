package user

import (
	"bufio"
	"crypto/rsa"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/zubeyiro/secure-chat/internal/configuration"
	"github.com/zubeyiro/secure-chat/internal/security"
)

var config *configuration.Configuration
var wg sync.WaitGroup
var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func Start() {
	config = configuration.GetConfig()
	privateKey, publicKey = security.GenerateKeyPair()

	wg.Add(1)

	conn, err := net.Dial(config.Server.Type, config.Server.Port)
	if err != nil {
		fmt.Println(err)
	}

	go Read(conn)
	go Write(conn)

	wg.Wait()
}

func Read(conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		str, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("DISCONNNECTED")
			wg.Done()

			return
		}

		fmt.Println(str)
	}
}

func Write(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(conn)

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		_, err = writer.WriteString(str)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = writer.Flush()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

/*
- User
	- Events
		- Incoming
			- connected
				- save servers public key
			- user_list_received
			- message_received
		- Outgoing
			- user_list_request
			- send_message
*/
