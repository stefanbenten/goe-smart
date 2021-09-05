package main

import (
	"stefan-benten.de/goe-smart/pkg/client"
	"stefan-benten.de/goe-smart/pkg/webserver"
)

func main() {
	srv, _ := webserver.NewServer()

	client.NewPowerFoxClient()

	srv.Run()
}
