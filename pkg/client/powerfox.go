package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	powerFoxUrl  = "https://backend.powerfox.energy/api/2.0/my/main/current?unit=kwh"
	pollInterval = 30 * time.Second
)

type PowerFoxClient struct {
	username string
	password string
	Status   PowerFoxStatus
}

func NewPowerFoxClient(username, password string) PowerFoxClient {
	return PowerFoxClient{username: username, password: password}
}

func (pfx *PowerFoxClient) Sub() {
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

				req.SetBasicAuth(pfx.username, pfx.password)
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

				log.Println(powerFoxData)
				pfx.Status = powerFoxData
			}
		}
	}()
}
