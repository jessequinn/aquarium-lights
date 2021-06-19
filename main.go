package main

import (
	"aquarium-lights/internal/helpers"
	"aquarium-lights/internal/models"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/apex/log"
	"github.com/stianeikeland/go-rpio"
)

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func main() {
	var relay bool

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

	// Open and map memory to access gpio, check for errors.
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	// Set pin to output mode.
	data.SetModeOutput()

	// Clean up on ctrl-c and turn lights out.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		data.SetHigh()
		os.Exit(0)
	}()

	// Unmap gpio memory when done.
	defer rpio.Close()

	// Turn lights off to start.
	data.SetHigh()

	for {
		for _, v := range data.Schedules {
			for _, p := range v.Periods {
				relay = false
				if inTimeSpan(helpers.Bod(p.Start), helpers.Bod(p.End), time.Now()) {
					relay = true
				}
				if relay {
					v.Pin.Low()
					ctx := log.WithFields(log.Fields{
						"name":       v.Name,
						"pin":        v.Pin,
						"start_time": p.Start.String(),
						"end_time":   p.End.String(),
					})
					ctx.Info(fmt.Sprintf("%s relay turned on", v.Name))
				} else {
					v.Pin.High()
				}
			}
		}
	}
}
