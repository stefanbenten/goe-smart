package handler

import (
	"stefan-benten.de/goe-smart/pkg/client"
)

type Config struct {
	MQTTAddress string
	GoESerial   string
	PfxUsername string
	PfxPassword string
}

type Handler struct {
	PowerFox *client.PowerFoxClient
	Charger  *client.MQTTClient

	Data
}

type Data struct {
	PowerFox client.PowerFoxStatus
	Charger  client.ChargerStatus
}

func NewHandler(config Config) (Handler, error) {

	pfx := client.NewPowerFoxClient(config.PfxUsername, config.PfxPassword)

	mqtt := client.NewMQTTClient(config.MQTTAddress, config.GoESerial)

	return Handler{&pfx, &mqtt, Data{PowerFox: client.PowerFoxStatus{}, Charger: client.ChargerStatus{}}}, nil
}

func (hdl *Handler) Start() (err error) {

	hdl.PowerFox.Sub(&hdl.Data.PowerFox)
	hdl.Charger.Sub(&hdl.Data.Charger)
	return
}
