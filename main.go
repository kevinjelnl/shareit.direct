package main

import (
	"fmt"
	"log"

	"github.com/kevinjelnl/shareit.direct/router"
)

const (
	ip string = "0.0.0.0"
	// ip   string = "localhost"
	port int = 8080
)

func main() {
	log.Print("Server started, exposed on: ", ip)
	e := router.New()
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", ip, port)))
}
