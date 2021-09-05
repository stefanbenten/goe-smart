package handler

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"stefan-benten.de/goe-smart/pkg/client"
	"stefan-benten.de/goe-smart/pkg/template"
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

	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for {
			<-ticker.C
			hdl.HandleCharging()
		}
	}()

	return
}

func (hdl *Handler) HandleCharging() {
	if hdl.Data.Charger.CarStatus >= 0 {
		hdl.ManageAmp()
	}
}

func (hdl *Handler) getVoltage() int {
	/*volt := (hdl.Data.Charger.Stats.L1Volt + hdl.Data.Charger.Stats.L2Volt + hdl.Data.Charger.Stats.L3Volt)/3
	if volt == 0 {
		return 230
	}
	return volt*/
	return 230
}

func (hdl *Handler) ManageAmp() {
	watt := -hdl.Data.PowerFox.Watt
	if watt > 0 {
		amp := int(watt) / hdl.getVoltage() / 3
		fmt.Println(watt, amp)
	}
}

func (hdl *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var buf bytes.Buffer

	data := struct {
		Charger  client.ChargerStats
		PowerFox client.PowerFoxStats
	}{
		Charger:  hdl.Data.Charger.ChargerStats,
		PowerFox: hdl.Data.PowerFox.PowerFoxStats,
	}

	if err = template.Index.Execute(&buf, data); err != nil {
		log.Println(err)
	}

	_, err = w.Write(buf.Bytes())
}
