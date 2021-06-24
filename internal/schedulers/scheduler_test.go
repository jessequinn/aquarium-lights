package schedulers

import (
	"aquarium-lights/internal/helpers"

	"context"
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func TestNewScheduler(t *testing.T) {
	tests := []struct {
		name string
		want *Scheduler
	}{
		{
			name: "test1",
			want: &Scheduler{
				wg:            new(sync.WaitGroup),
				cancellations: make([]context.CancelFunc, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScheduler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScheduler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScheduler_Add(t *testing.T) {
	type args struct {
		ctx context.Context
		j   Job
		p   time.Duration
		o   time.Duration
	}
	tests := []struct {
		name string
		s    *Scheduler
		args args
	}{
		{
			name: "test1",
			s: &Scheduler{
				wg:            new(sync.WaitGroup),
				cancellations: make([]context.CancelFunc, 0),
			},
			args: args{
				ctx: context.WithValue(context.Background(), "values", helpers.ContextWithValue{Name: "test", Pin: rpio.Pin(22)}),
				j: func(ctx context.Context) {
					// Turn off
					value, ok := ctx.Value("values").(helpers.ContextWithValue)
					if ok {
						value.Pin.High()
						fmt.Printf("Device %s on pin %d turned off at %s\n", value.Name, value.Pin, time.Now().String())
					} else {
						fmt.Println("Could not retrieve values from context")
					}
				},
				p: time.Second * 5,
				o: time.Second * 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.args.ctx, tt.args.j, tt.args.p, tt.args.o)
		})
	}
}

func TestScheduler_Stop(t *testing.T) {
	tests := []struct {
		name string
		s    *Scheduler
	}{
		{
			name: "test1",
			s: &Scheduler{
				wg:            new(sync.WaitGroup),
				cancellations: make([]context.CancelFunc, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Stop()
		})
	}
}

//func TestScheduler_process(t *testing.T) {
//	ctxWithCancel, _ := context.WithCancel(context.WithValue(context.Background(), "values", helpers.ContextWithValue{Name: "test", Pin: rpio.Pin(22)}))
//	type args struct {
//		ctx context.Context
//		j   Job
//		p   time.Duration
//		o   time.Duration
//	}
//	tests := []struct {
//		name string
//		s    *Scheduler
//		args args
//	}{
//		{
//			name: "test1",
//			s: &Scheduler{
//				wg:            new(sync.WaitGroup),
//				cancellations: make([]context.CancelFunc, 0),
//			},
//			args: args{
//				ctx: ctxWithCancel,
//				j: func(ctx context.Context) {
//					// Turn off
//					value, ok := ctx.Value("values").(helpers.ContextWithValue)
//					if ok {
//						value.Pin.High()
//						fmt.Printf("Device %s on pin %d turned off at %s\n", value.Name, value.Pin, time.Now().String())
//					} else {
//						fmt.Println("Could not retrieve values from context")
//					}
//				},
//				p: time.Second * 5,
//				o: time.Second * 5,
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.s.process(tt.args.ctx, tt.args.j, tt.args.p, tt.args.o)
//		})
//	}
//}
