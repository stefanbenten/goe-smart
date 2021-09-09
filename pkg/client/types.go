package client

import (
	"encoding/json"
	"sync"
)

type PowerFoxStatus struct {
	Lock sync.Mutex
	PowerFoxStats
}

type PowerFoxStats struct {
	Outdated  bool    `json:"Outdated"`
	Timestamp int64   `json:"Timestamp"`
	Watt      float64 `json:"Watt"`
	WattIn    float64 `json:"A_Plus"`
	WattOut   float64 `json:"A_Minus"`
}

type ChargerStatus struct {
	Lock sync.Mutex
	ChargerStats
}

type ChargerStats struct {
	CarStatus   int   `json:"car,string"`
	Amperage    int   `json:"amp,string"`
	Allowed     int   `json:"alw,string"`
	Temperature int   `json:"tmp,string"`
	ChargedKWh  int64 `json:"dws,string"`
	TotalKWh    int64 `json:"eto,string"`
	Serial      int   `json:"sse,string"`
	Stats       Stats `json:"nrg,string"`
}

type Stats struct {
	L1Volt  int
	L2Volt  int
	L3Volt  int
	NVolt   int
	L1Amp   int
	L2Amp   int
	L3Amp   int
	L1Load  int
	L2Load  int
	L3Load  int
	NLoad   int
	SumLoad int
	L1Pct   int
	L2Pct   int
	L3Pct   int
	NPct    int
}

func (s *Stats) UnmarshalJSON(data []byte) (err error) {
	var array [16]int
	err = json.Unmarshal(data, &array)
	if err != nil {
		return err
	}
	s.L1Volt = array[0]
	s.L2Volt = array[1]
	s.L3Volt = array[2]
	s.NVolt = array[3]
	s.L1Amp = array[4]
	s.L2Amp = array[5]
	s.L3Amp = array[6]
	s.L1Load = array[7]
	s.L2Load = array[8]
	s.L3Load = array[9]
	s.NLoad = array[10]
	s.SumLoad = array[11]
	s.L1Pct = array[12]
	s.L2Pct = array[13]
	s.L3Pct = array[14]
	s.NPct = array[15]
	return nil
}
