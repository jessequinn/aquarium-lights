package helpers

import (
	"aquarium-lights/internal/models"
	"reflect"
	"testing"
	"time"
)

func TestBod(t *testing.T) {
	timeNow := time.Now()
	type args struct {
		t models.CustomTime
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "test1",
			args: args{t: models.CustomTime{Time: timeNow}},
			want: time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second(), 0, timeNow.Location()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bod(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bod() = %v, want %v", got, tt.want)
			}
		})
	}
}
