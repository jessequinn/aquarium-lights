package main

import (
	"aquarium_lights/internal/helpers"
	"aquarium_lights/internal/models"
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	"io/ioutil"
	"os"
	"time"
)

//var (
//	display = rpio.Pin(22)
//	sump    = rpio.Pin(23)
//)



func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func main() {
	var relay bool

	//// Open and map memory to access gpio, check for errors.
	//if err := rpio.Open(); err != nil {
	//	panic(err)
	//}
	//
	//// Set pin to output mode.
	//display.Output()
	//sump.Output()
	//
	//// Clean up on ctrl-c and turn lights out.
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	//go func() {
	//	<-c
	//	display.High()
	//	sump.High()
	//	os.Exit(0)
	//}()
	//
	//// Unmap gpio memory when done.
	//defer rpio.Close()
	//
	//// Turn lights off to start.
	//display.High()
	//sump.High()
	//
	//// configure schedules
	//schs := Schedules{}
	//
	//if err := schs.configure("sump", sump, []periodString{{"2021-01-01T18:00:00.000-0300", "2021-12-31T23:59:00.000-0300"}, {"2021-01-01T00:00:00.000-0300", "2021-12-31T10:00:00.000-0300"}}); err != nil {
	//	panic(err)
	//}
	//
	//if err := schs.configure("display", display, []periodString{{"2021-01-01T10:00:00.000-0300", "2021-12-31T18:00:00.000-0300"}}); err != nil {
	//	panic(err)
	//}
	//

	jsonFile, err := os.Open("configuration.json")
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	file, _ := ioutil.ReadAll(jsonFile)
	data := models.Schedules{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		panic(err)
	}

	for {
		for _, v := range data.Schedules {
			for _, p := range v.Periods {
				relay = false
				if inTimeSpan(helpers.Bod(p.Start), helpers.Bod(p.End), time.Now()) {
					relay = true
				}
				if relay {
					//v.Pin.Low()
					ctx := log.WithFields(log.Fields{
						"name":       v.Name,
						"pin":        v.Pin,
						"start_time": p.Start.String(),
						"end_time":   p.End.String(),
					})
					ctx.Info(fmt.Sprintf("%s relay turned on", v.Name))
				} else {
					//v.Pin.High()
				}
			}
		}
	}
}
