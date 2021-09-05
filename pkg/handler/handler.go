package handler

import (
	"stefan-benten.de/goe-smart/pkg/client"
	"stefan-benten.de/goe-smart/pkg/webserver"
)

type Config struct {
	MQTTAddress string
	GoESerial   string
	PfxUsername string
	PfxPassword string
	WebAddress  string
}

type Handler struct {
	Web      *webserver.Server
	PowerFox *client.PowerFoxClient
	Charger  *client.MQTTClient
}

func NewHandler(config Config) (Handler, error) {

	pfx := client.NewPowerFoxClient(config.PfxUsername, config.PfxPassword)

	mqtt := client.NewMQTTClient(config.MQTTAddress, config.GoESerial)

	web, err := webserver.NewServer(config.WebAddress)
	if err != nil {
		return Handler{}, err
	}

	return Handler{&web, &pfx, &mqtt}, nil
}

func (hdl *Handler) Start() (err error) {
	hdl.PowerFox.Sub()
	hdl.Charger.Sub()
	hdl.Web.Run()
	return
}
