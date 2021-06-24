package models

import (
	"github.com/stianeikeland/go-rpio"
)

type SchedulesInterface interface {
	SetModeOutput()
	SetHigh()
	GetSchedules() []Schedule
}

type Schedules struct {
	Schedules []Schedule `json:"schedules"`
}

type Schedule struct {
	Name    string   `json:"name"`
	Pin     rpio.Pin `json:"pin"`
	Periods []Period `json:"periods"`
}

type Period struct {
	Start CustomTime `json:"start_time"`
	End   CustomTime `json:"end_time"`
}

// GetSchedules returns all Schedules.
func (s *Schedules) GetSchedules() []Schedule {
	return s.Schedules
}

// SetModeOutput sets all pins to output mode.
func (s *Schedules) SetModeOutput() {
	for _, v := range s.Schedules {
		v.Pin.Output()
	}
}

// SetHigh sets all pins to high.
func (s *Schedules) SetHigh() {
	for _, v := range s.Schedules {
		v.Pin.High()
	}
}
