package server

import (
	"fmt"
	"net"
	"os"

	"github.com/zubeyiro/secure-chat/internal/configuration"
)

var config *configuration.Configuration

func Start() {
	config = configuration.GetConfig()
	users = make(map[string]*User)
	guests = make(map[string]*Guest)
	userMap = make(map[string]string)

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

		newGuest(conn).join()
	}
}
