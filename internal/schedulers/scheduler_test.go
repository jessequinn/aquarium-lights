package schedulers

import (
	"context"
	"os"
	"reflect"
	"sync"
	"testing"
	"time"
)

var S Scheduler

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	S = Scheduler{
		wg:            new(sync.WaitGroup),
		cancellations: make([]context.CancelFunc, 0),
	}
}

func shutdown() {
	S.Stop()
}

func TestNewScheduler(t *testing.T) {
	tests := []struct {
		name string
		want *Scheduler
	}{
		{
			name: "test1",
			want: &S,
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
			s:    &S,
			args: args{
				ctx: context.Background(),
				j: func(ctx context.Context) {
					t.Log("TestScheduler_Add")
				},
				p: time.Second * 5,
				o: 0,
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
			s:    &S,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Stop()
		})
	}
}

func TestScheduler_process(t *testing.T) {
	ctxWithCancel, cancel := context.WithCancel(context.Background())
	S.cancellations = append(S.cancellations, cancel)
	S.wg.Add(1)

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
			s:    &S,
			args: args{
				ctx: ctxWithCancel,
				j: func(ctx context.Context) {
					t.Log("TestScheduler_process")
					cancel()
				},
				p: time.Second * 1,
				o: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.process(tt.args.ctx, tt.args.j, tt.args.p, tt.args.o)
		})
	}
}
