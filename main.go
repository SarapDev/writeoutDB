package main

import (
	"fmt"

	"github.com/SarapDev/writeoutDB/src/core"
	"github.com/SarapDev/writeoutDB/src/server"
)

func main() {
	fmt.Println("WriteOutDB v0.0.0")

	// http server to send requests
	go server.RunServer() 
	
	// run main application loop
	core.Run()
}

