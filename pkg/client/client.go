package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	powerFoxUrl  = "https://backend.powerfox.energy/api/2.0/my/main/current?unit=kwh"
	pollInterval = 30 * time.Second
)

type PowerFoxClient struct {
}

type MQTTClient struct {
	Client    mqtt.Client
	goeSerial string
}

func NewPowerFoxClient(username, password string) PowerFoxClient {

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

				req.SetBasicAuth(username, password)
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

	return PowerFoxClient{}
}

func NewMQTTClient(addr, goeSerial string) MQTTClient {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", addr))
	mqttClient := mqtt.NewClient(opts)

	token := mqttClient.Connect()
	token.WaitTimeout(5 * time.Second)

	return MQTTClient{Client: mqttClient}
}

func (cl *MQTTClient) Sub() {
	cl.Client.Subscribe("go-eCharger/"+cl.goeSerial+"/status", 0, mqttStatusHandler)
}

func mqttStatusHandler(_ mqtt.Client, message mqtt.Message) {
	var status ChargerStatus
	err := json.Unmarshal(message.Payload(), &status)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(status)
}
