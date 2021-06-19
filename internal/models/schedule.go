package models

import (
	"github.com/stianeikeland/go-rpio"
)

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

type periodString struct {
	Start string
	End   string
}

//func (s *Schedules) configure(name string, pin rpio.Pin, ps []periodString) error {
//	var p []Period
//	layout := "2006-01-02T15:04:05.000Z0700"
//	for _, v := range ps {
//		st, err := time.Parse(layout, v.Start)
//		if err != nil {
//			return err
//		}
//		en, err := time.Parse(layout, v.End)
//		if err != nil {
//			return err
//		}
//		p = append(p, Period{
//			Start: st,
//			End:   en,
//		})
//	}
//	s.Schedules = append(s.Schedules, Schedule{
//		Name:    name,
//		Pin:     pin,
//		Periods: p,
//	})
//	return nil
//}

//func (s *Schedules) ListSchedules() []Schedule {
//	return s.Schedules
//}