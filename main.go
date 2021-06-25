package main

import (
	"aquarium-lights/internal/helpers"
	"aquarium-lights/internal/models"
	"aquarium-lights/internal/schedulers"
	"github.com/apex/log"

	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	var relay bool
	var si models.SchedulesInterface

	// Read configuration from JSON.
	jsonFile, err := os.Open("configuration.json")
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	file, _ := ioutil.ReadAll(jsonFile)
	si = &models.Schedules{}
	err = json.Unmarshal(file, &si)
	if err != nil {
		panic(err)
	}

	// Open and map memory to access gpio, check for errors.
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	// Set pin to output mode.
	si.SetModeOutput()

	// Unmap gpio memory when done.
	defer rpio.Close()

	// Turn lights off to start.
	si.SetHigh()

	ctx := context.Background()
	worker := schedulers.NewScheduler()

	// UTC-3 12 = 9
	for _, v := range si.GetSchedules() {
		for _, p := range v.Periods {
			relay = false
			if helpers.InTimeSpan(helpers.Bod(p.Start), helpers.Bod(p.End), time.Now()) {
				relay = true
			}
			if relay {
				v.Pin.Low()
				logCtx := log.WithFields(log.Fields{
					"name":       v.Name,
					"pin":        v.Pin,
					"start_time": p.Start.String(),
					"end_time":   p.End.String(),
				})
				logCtx.Info(fmt.Sprintf("Device %s on pin %d turned on at %s\n", v.Name, v.Pin, time.Now().String()))
			} else {
				v.Pin.High()
			}
			worker.Add(context.WithValue(ctx, "values", helpers.ContextWithValue{
				Name: v.Name,
				Pin:  v.Pin,
			}), func(ctx context.Context) {
				// Turn on
				value, ok := ctx.Value("values").(helpers.ContextWithValue)
				if ok {
					value.Pin.Low()
					logCtx := log.WithFields(log.Fields{
						"name": value.Name,
						"pin":  value.Pin,
						"time": time.Now().String(),
					})
					logCtx.Info("Device turned on")
				} else {
					logCtx := log.WithFields(log.Fields{
						"time": time.Now().String(),
					})
					logCtx.Info("Could not retrieve values from context")
				}
			}, time.Hour*24, time.Hour*time.Duration(p.Start.Hour()+3)+time.Minute*time.Duration(p.Start.Minute()))
			worker.Add(context.WithValue(ctx, "values", helpers.ContextWithValue{
				Name: v.Name,
				Pin:  v.Pin,
			}), func(ctx context.Context) {
				// Turn off
				value, ok := ctx.Value("values").(helpers.ContextWithValue)
				if ok {
					value.Pin.High()
					fmt.Printf("Device %s on pin %d turned off at %s\n", value.Name, value.Pin, time.Now().String())
				} else {
					fmt.Println("Could not retrieve values from context")
				}
			}, time.Hour*24, time.Hour*time.Duration(p.End.Hour()+3)+time.Minute*time.Duration(p.End.Minute()))
		}
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	worker.Stop()
	si.SetHigh()
}
