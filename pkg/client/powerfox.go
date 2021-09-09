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
	powerFoxUrl = "https://backend.powerfox.energy/api/2.0/my/main/current?unit=kwh"
)

type PowerFoxClient struct {
	username     string
	password     string
	pollInterval time.Duration
	ticker       *time.Ticker
}

func NewPowerFoxClient(username, password string, pollInterval time.Duration) PowerFoxClient {
	return PowerFoxClient{username: username, password: password, pollInterval: pollInterval}
}

func (pfx *PowerFoxClient) Sub(status *PowerFoxStatus) {
	pfx.ticker = time.NewTicker(pfx.pollInterval)

	// Trigger once first.
	pfx.query(status)

	// Fire of go routine to handle polling.
	go func() {
		for {
			<-pfx.ticker.C
			pfx.query(status)
		}
	}()
}

func (pfx *PowerFoxClient) query(status *PowerFoxStatus) {
	client := &http.Client{}
	var powerFoxData = PowerFoxStats{}

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
	status.Lock.Lock()
	status.PowerFoxStats = powerFoxData
	status.Lock.Unlock()
}

func (pfx *PowerFoxClient) PollAdjust(dur time.Duration) {
	pfx.ticker.Reset(dur)
}
