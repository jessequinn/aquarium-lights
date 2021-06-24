package models

//import (
//	"github.com/stianeikeland/go-rpio"
//	"testing"
//)
//
//func TestSchedules_SetModeOutput(t *testing.T) {
//	tests := []struct {
//		name string
//		s    *Schedules
//	}{
//		{
//			name: "test1",
//			s: &Schedules{
//				Schedules: []Schedule{
//					{
//						Name:    "test1",
//						Pin:     rpio.Pin(22),
//						Periods: nil,
//					},
//				},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.s.SetModeOutput()
//		})
//	}
//}
//
//func TestSchedules_SetHigh(t *testing.T) {
//	tests := []struct {
//		name string
//		s    *Schedules
//	}{
//		{
//			name: "test1",
//			s: &Schedules{
//				Schedules: []Schedule{
//					{
//						Name:    "test1",
//						Pin:     rpio.Pin(22),
//						Periods: nil,
//					},
//				},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.s.SetHigh()
//		})
//	}
//}
