package models

import (
	"reflect"
	"testing"
	"time"

	"github.com/stianeikeland/go-rpio"
)

// SchedulesMock struct.
type SchedulesMock struct {
	Schedules []Schedule
}

// Mocked functions. Cannot actually mock /dev/mem.
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

func TestSchedules_GetSchedules(t *testing.T) {
	timeNow := time.Now()
	tests := []struct {
		name string
		s    *Schedules
		want []Schedule
	}{
		{
			name: "test1",
			s: &Schedules{
				Schedules: []Schedule{
					{
						Name: "test1",
						Pin:  rpio.Pin(22),
						Periods: []Period{
							{
								Start: CustomTime{timeNow},
								End:   CustomTime{timeNow},
							},
						},
					},
				},
			},
			want: []Schedule{
				{
					Name: "test1",
					Pin:  rpio.Pin(22),
					Periods: []Period{
						{
							Start: CustomTime{timeNow},
							End:   CustomTime{timeNow},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetSchedules(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schedules.GetSchedules() = %v, want %v", got, tt.want)
			}
		})
	}
}
