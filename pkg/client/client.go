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

const (
	powerFoxUrl = "https://backend.powerfox.energy/api/2.0/my/main/current?unit=kwh"
)

func NewPowerFoxClient() {
	client := &http.Client{}
	var powerFoxData = PowerFoxStatus{}

	req, err := http.NewRequest("GET", powerFoxUrl, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user := os.Getenv("PF_USER")
	pass := os.Getenv("PF_PASS")

	req.SetBasicAuth(user, pass)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(data, &powerFoxData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(powerFoxData)

	mqttHost := os.Getenv("MQTT_HOST")
	goeSerial := os.Getenv("GOE_SERIAL")

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", mqttHost))
	mqttClient := mqtt.NewClient(opts)

	token := mqttClient.Connect()
	token.WaitTimeout(5 * time.Second)

	mqttClient.Subscribe("go-eCharger/"+goeSerial+"/cmd/req", 0, mqttHandler)
}

func mqttHandler(client mqtt.Client, message mqtt.Message) {
	fmt.Println(message.Payload())
}
