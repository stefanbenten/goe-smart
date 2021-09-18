package handler

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"time"

	"stefan-benten.de/goe-smart/pkg/client"
	tmpl "stefan-benten.de/goe-smart/pkg/template"
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
	UserSettings
}

type UserSettings struct {
	SleepState    bool
	ForceCharge   bool
	AmpOverride   int
	PhaseOverride int
	LoadThreshold int
}

const (
	pollInterval      = 30 * time.Second
	pollIntervalSleep = 5 * time.Minute
	checkInterval     = 30 * time.Second
)

func NewHandler(config Config) (Handler, error) {

	pfx := client.NewPowerFoxClient(config.PfxUsername, config.PfxPassword, pollInterval)

	mqtt := client.NewMQTTClient(config.MQTTAddress, config.GoESerial)

	return Handler{&pfx, &mqtt, Data{PowerFox: client.PowerFoxStatus{}, Charger: client.ChargerStatus{}, UserSettings: UserSettings{SleepState: false, ForceCharge: false, AmpOverride: 16, PhaseOverride: 3, LoadThreshold: 2000}}}, nil
}

func (hdl *Handler) Start() (err error) {
	hdl.PowerFox.Sub(&hdl.Data.PowerFox)
	hdl.Charger.Sub(&hdl.Data.Charger)

	time.Sleep(5 * time.Second)
	hdl.HandleSleepState()
	hdl.HandleCharging()

	ticker := time.NewTicker(checkInterval)

	go func() {
		for {
			<-ticker.C
			hdl.HandleSleepState()
			hdl.HandleCharging()
		}
	}()

	return
}

func (hdl *Handler) HandleSleepState() {
	cstatus := hdl.Data.Charger.CarStatus
	status := hdl.SleepState
	if cstatus > 1 && status == true {
		log.Println("Resuming from Sleep Mode")
		hdl.SleepState = false
		hdl.PowerFox.PollAdjust(pollInterval)
	} else if cstatus == 1 && status == false {
		log.Println("Falling into Sleep Mode")
		hdl.SleepState = true
		hdl.PowerFox.PollAdjust(pollIntervalSleep)
	}
}

func (hdl *Handler) HandleCharging() {
	switch hdl.Data.Charger.CarStatus {
	case 4:
		fallthrough
	case 3:
		if hdl.SleepState {
			hdl.SleepState = false
		}
		if hdl.Data.Charger.Allowed == 0 && int(-hdl.Data.PowerFox.Watt) > hdl.LoadThreshold {
			hdl.Charger.Pub("alw", 1)
		}
		if hdl.Data.Charger.Amperage > 6 {
			hdl.Charger.Pub("amp", 6)
		}
	case 2:
		if hdl.SleepState {
			hdl.SleepState = false
		}
		amp := hdl.manageAmp(int(-hdl.Data.PowerFox.Watt) / hdl.getVoltage() / hdl.PhaseOverride)
		if amp < 6 {
			hdl.Charger.Pub("amp", 6)
			hdl.Charger.Pub("alw", 0)
		}
		hdl.Charger.Pub("amp", amp)
	}
}

func (hdl *Handler) getVoltage() int {
	volt := (hdl.Data.Charger.Stats.L1Volt + hdl.Data.Charger.Stats.L2Volt + hdl.Data.Charger.Stats.L3Volt) / 3
	if volt == 0 {
		return 230
	}
	return volt
}

// 0 - Leave currently value
// 1-5 - Turn off charging
// 6-16 - Adjust to value
func (hdl *Handler) manageAmp(diffAmp int) (newValue int) {
	if hdl.Data.ForceCharge {
		return hdl.Data.AmpOverride
	}

	if diffAmp == 0 {
		return 0
	}

	curAmp := hdl.Data.Charger.Amperage

	if curAmp+diffAmp > 16 {
		newValue = 16
	} else {
		newValue = curAmp + diffAmp
	}
	return newValue
}

func (hdl *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var buf bytes.Buffer

	if r.Method == "POST" {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		amp, err := strconv.Atoi(r.PostFormValue("ampOverride"))
		if err != nil {
			log.Println(err)
		} else {
			hdl.AmpOverride = amp
		}
		phases, err := strconv.Atoi(r.PostFormValue("phaseOverride"))
		if err != nil {
			log.Println(err)
		} else {
			hdl.PhaseOverride = phases
		}
		load, err := strconv.Atoi(r.PostFormValue("loadThreshold"))
		if err != nil {
			log.Println(err)
		} else {
			hdl.LoadThreshold = load
		}
		force, err := strconv.ParseBool(r.PostFormValue("force"))
		if err != nil {
			log.Println(err)
		} else {
			hdl.ForceCharge = force
		}
	}

	log.Println(hdl.UserSettings)

	data := struct {
		Charger   client.ChargerStats
		PowerFox  client.PowerFoxStats
		Overrides UserSettings
	}{
		Charger:   hdl.Data.Charger.ChargerStats,
		PowerFox:  hdl.Data.PowerFox.PowerFoxStats,
		Overrides: hdl.Data.UserSettings,
	}

	if err = tmpl.Index.Execute(&buf, data); err != nil {
		log.Println(err)
	}

	_, err = w.Write(buf.Bytes())
}
