package main

import (
	"aquarium-lights/internal/models"
	"aquarium-lights/internal/schedulers"
	"context"
	"github.com/stianeikeland/go-rpio"
	"syscall"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Read configuration from JSON.
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

	//Open and map memory to access gpio, check for errors.
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

	ctx := context.Background()
	worker := schedulers.NewScheduler()

	// UTC-3 12 = 9
	for _, v := range data.Schedules {
		for _, p := range v.Periods {
			worker.Add(ctx, func(ctx context.Context) {
				// Turn on
				v.Pin.Low()
				fmt.Printf("%s %d %v\n", v.Name, v.Pin, p.Start)
			}, time.Hour*24, time.Hour*time.Duration(p.Start.Hour()+3)+time.Minute*time.Duration(p.Start.Minute()))
			worker.Add(ctx, func(ctx context.Context) {
				// Turn off
				v.Pin.High()
				fmt.Printf("%s %d %v\n", v.Name, v.Pin, p.End)
			}, time.Hour*24, time.Hour*time.Duration(p.End.Hour()+3)+time.Minute*time.Duration(p.End.Minute()))
		}
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit
	worker.Stop()
}
