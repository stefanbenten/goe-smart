package client

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTClient struct {
	client    mqtt.Client
	goeSerial string
	Status    ChargerStatus
}

func NewMQTTClient(addr, goeSerial string) MQTTClient {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", addr))
	mqttClient := mqtt.NewClient(opts)

	token := mqttClient.Connect()
	token.WaitTimeout(5 * time.Second)

	return MQTTClient{client: mqttClient, goeSerial: goeSerial}
}

func (cl *MQTTClient) Sub() {
	cl.client.Subscribe("go-eCharger/"+cl.goeSerial+"/status", 0, cl.mqttStatusHandler)
}

func (cl *MQTTClient) mqttStatusHandler(_ mqtt.Client, message mqtt.Message) {
	var status ChargerStatus
	err := json.Unmarshal(message.Payload(), &status)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	log.Println(status)
	cl.Status = status
}
