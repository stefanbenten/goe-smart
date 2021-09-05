package main

import (
	"log"
	"os"

	"stefan-benten.de/goe-smart/pkg/handler"
	"stefan-benten.de/goe-smart/pkg/webserver"
)

func main() {
	mqttHost := os.Getenv("MQTT_HOST")
	goeSerial := os.Getenv("GOE_SERIAL")

	user := os.Getenv("PF_USER")
	pass := os.Getenv("PF_PASS")

	config := handler.Config{
		MQTTAddress: mqttHost,
		GoESerial:   goeSerial,
		PfxUsername: user,
		PfxPassword: pass,
	}

	hdl, err := handler.NewHandler(config)
	if err != nil {
		log.Fatalln(err)
	}
	err = hdl.Start()
	if err != nil {
		log.Fatalln(err)
	}

	web, err := webserver.NewServer(":8080", &hdl)
	if err != nil {
		log.Fatalln(err)
	}
	web.Run()
}
