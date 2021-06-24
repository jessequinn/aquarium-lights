package models

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestCustomTime_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		ct      *CustomTime
		args    args
		wantErr bool
	}{
		{
			name:    "test1",
			ct:      &CustomTime{},
			args:    args{[]byte("2021-01-01T10:00:00.000-0300")},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ct.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("CustomTime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCustomTime_MarshalJSON(t *testing.T) {
	time, _ := time.Parse(ctLayout, "2021-01-01T10:00:00.000-0300")
	tests := []struct {
		name    string
		ct      *CustomTime
		want    []byte
		wantErr bool
	}{
		{
			name:    "test1",
			ct:      &CustomTime{time},
			want:    []byte(fmt.Sprintf("%q", "2021-01-01T10:00:00.000-0300")),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ct.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomTime.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomTime.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomTime_IsSet(t *testing.T) {
	time, _ := time.Parse(ctLayout, "2021-01-01T10:00:00.000-0300")
	tests := []struct {
		name string
		ct   *CustomTime
		want bool
	}{
		{
			name: "test1",
			ct:   &CustomTime{time},
			want: true,
		},
		{
			name: "test2",
			ct:   &CustomTime{},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ct.IsSet(); got != tt.want {
				t.Errorf("CustomTime.IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
