package models

import (
	"testing"

	"github.com/stianeikeland/go-rpio"
)

type SchedulesMock struct {
	Schedules []Schedule
}

func (sd *SchedulesMock) SetModeOutput() {}
func (sd *SchedulesMock) SetHigh()       {}

func TestSchedules_SetModeOutput(t *testing.T) {
	tests := []struct {
		name string
		s    *SchedulesMock
	}{
		{
			name: "test1",
			s: &SchedulesMock{Schedules: []Schedule{{
				Name:    "test1",
				Pin:     22,
				Periods: nil,
			}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.SetModeOutput()
		})
	}
}

func TestSchedules_SetHigh(t *testing.T) {
	tests := []struct {
		name string
		s    *SchedulesMock
	}{
		{
			name: "test1",
			s: &SchedulesMock{
				Schedules: []Schedule{
					{
						Name:    "test1",
						Pin:     rpio.Pin(22),
						Periods: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.SetHigh()
		})
	}
}
