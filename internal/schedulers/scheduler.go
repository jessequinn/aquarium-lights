package schedulers

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job func(ctx context.Context)

type Scheduler struct {
	wg            *sync.WaitGroup
	cancellations []context.CancelFunc
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		wg:            new(sync.WaitGroup),
		cancellations: make([]context.CancelFunc, 0),
	}
}

// Add starts goroutine which constantly calls provided job with interval delay
func (s *Scheduler) Add(ctx context.Context, j Job, p, o time.Duration) {
	ctx, cancel := context.WithCancel(ctx)
	s.cancellations = append(s.cancellations, cancel)

	s.wg.Add(1)
	go s.process(ctx, j, p, o)
}

// Stop cancels all running jobs
func (s *Scheduler) Stop() {
	for _, cancel := range s.cancellations {
		cancel()
	}
	s.wg.Wait()
}

func (s *Scheduler) process(ctx context.Context, j Job, p, o time.Duration) {
	first := time.Now().Truncate(p).Add(o)
	if first.Before(time.Now()) {
		first = first.Add(p)
	}
	firstC := time.After(first.Sub(time.Now()))
	fmt.Printf("time now: %v, first: %v\n", time.Now(), first)

	// Receiving from a nil channel blocks forever
	ticker := &time.Ticker{C: nil}

	for {
		select {
		case <-firstC:
			fmt.Println("firstC")
			// The ticker has to be started before j as it can take some time to finish
			ticker = time.NewTicker(p)
			j(ctx)
		case <-ticker.C:
			fmt.Println("ticker")
			j(ctx)
		case <-ctx.Done():
			s.wg.Done()
			ticker.Stop()
			return
		}
	}
}
