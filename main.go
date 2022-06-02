package main

import (
	"fmt"
	"os"

	"github.com/zubeyiro/secure-chat/internal/server"
	"github.com/zubeyiro/secure-chat/internal/user"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid call")

		return
	}

	switch os.Args[1] {
	case "server":
		server.Start()
	case "user":
		user.Start()
	default:
		fmt.Println("Invalid call")
	}
}
