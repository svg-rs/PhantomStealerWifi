package main

import (
	"flag"
	"fmt"
	"phantomstealer/client"
	"phantomstealer/server"
)

func main() {
	mode := flag.String("mode", "", "Set to 'client' or 'server'")
	flag.Parse()

	switch *mode {
	case "client":
		client.Send()
	case "server":
		server.Serve()
	default:
		fmt.Println("Please specify -mode=client or -mode=server")
	}
}
