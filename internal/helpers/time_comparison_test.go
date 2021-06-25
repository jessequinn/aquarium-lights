package helpers

import (
	"testing"
	"time"
)

// ctLayout custom layout.
const ctLayout = "2006-01-02T15:04:05.000Z0700"

func TestInTimeSpan(t *testing.T) {
	timeStart := time.Now().Local().Add(-1 * time.Hour * time.Duration(1))
	timeEnd := time.Now().Local().Add(time.Hour * time.Duration(5))
	timeNow := time.Now()
	type args struct {
		start time.Time
		end   time.Time
		check time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				start: timeNow,
				end:   timeNow,
				check: timeNow,
			},
			want: false,
		},
		{
			name: "test2",
			args: args{
				start: timeStart,
				end:   timeEnd,
				check: timeNow,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InTimeSpan(tt.args.start, tt.args.end, tt.args.check); got != tt.want {
				t.Errorf("InTimeSpan() = %v, want %v", got, tt.want)
			}
		})
	}
}
