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
	status    *ChargerStatus
	goeSerial string
}

func NewMQTTClient(addr, goeSerial string) MQTTClient {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", addr))
	mqttClient := mqtt.NewClient(opts)

	token := mqttClient.Connect()
	token.WaitTimeout(5 * time.Second)

	return MQTTClient{client: mqttClient, goeSerial: goeSerial}
}

func (cl *MQTTClient) Sub(status *ChargerStatus) {
	cl.status = status
	cl.client.Subscribe("go-eCharger/"+cl.goeSerial+"/status", 0, cl.mqttStatusHandler)
}

func (cl *MQTTClient) mqttStatusHandler(_ mqtt.Client, message mqtt.Message) {
	var stats ChargerStats
	err := json.Unmarshal(message.Payload(), &stats)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	log.Println(stats)
	cl.status.Lock.Lock()
	cl.status.ChargerStats = stats
	cl.status.Lock.Unlock()
}
