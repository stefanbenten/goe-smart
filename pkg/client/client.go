package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type PowerFoxStatus struct {
	Outdated  bool    `json:"Outdated"`
	Timestamp int64   `json:"Timestamp"`
	Watt      float64 `json:"Watt"`
	WattIn    float64 `json:"A_Plus"`
	WattOut   float64 `json:"A_Minus"`
}

type ChargerStatus struct {
	CarStatus   int `json:"car"`
	Amperage    int `json:"amp"`
	Allowed     int `json:"alw"`
	Temperature int `json:"tmp"`
}

const (
	powerFoxUrl  = "https://backend.powerfox.energy/api/2.0/my/main/current?unit=kwh"
	pollInterval = 30 * time.Second
)

func NewPowerFoxClient() {
	user := os.Getenv("PF_USER")
	pass := os.Getenv("PF_PASS")

	timer := time.NewTicker(pollInterval)
	done := make(chan bool)

	// Fire of go routine to handle polling.
	go func() {
		for {
			select {
			case <-done:
				return
			case <-timer.C:

				client := &http.Client{}
				var powerFoxData = PowerFoxStatus{}

				req, err := http.NewRequest("GET", powerFoxUrl, nil)
				if err != nil {
					fmt.Println(err.Error())
					return
				}

				req.SetBasicAuth(user, pass)
				req.Header.Set("Accept", "application/json")

				resp, err := client.Do(req)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				defer resp.Body.Close()

				data, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err.Error())
					return
				}

				err = json.Unmarshal(data, &powerFoxData)
				if err != nil {
					fmt.Println(err.Error())
					return
				}

				fmt.Println(powerFoxData)
			}
		}
	}()
}

func NewMQTTClient() {
	mqttHost := os.Getenv("MQTT_HOST")
	goeSerial := os.Getenv("GOE_SERIAL")

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", mqttHost))
	mqttClient := mqtt.NewClient(opts)

	token := mqttClient.Connect()
	token.WaitTimeout(5 * time.Second)

	mqttClient.Subscribe("go-eCharger/"+goeSerial+"/status", 0, mqttStatusHandler)
}

func mqttStatusHandler(client mqtt.Client, message mqtt.Message) {
	var status ChargerStatus
	err := json.Unmarshal(message.Payload(), &status)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(status)
}
