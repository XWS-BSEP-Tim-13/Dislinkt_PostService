package main

import (
	"fmt"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/startup"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	fmt.Println("Server is starting...")
	server.Start()
}
