package main

import (
	"fmt"
	"go1f/pkg/server"
	"os"

)




func main() {
	port := os.Getenv("TODO_PORT")
	fmt.Println("port is: ", port)
	if port == "" {
		port = ":7540"
		fmt.Println("port is empty, setting to default: ", port)
	}
	server.Run(port)

}
