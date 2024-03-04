package handlers

import (
	"net/http"
	"testing"
)

func TestGetStatus(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetStatusFromEndPoints(tt.args.w, tt.args.r)
		})
	}
}

func TestStartUptimeTracking(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StartUptimeTracking()
		})
	}
}

func Test_checkServiceStatus(t *testing.T) {
	type args struct {
		serviceURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkServiceStatus(tt.args.serviceURL); got != tt.want {
				t.Errorf("checkServiceStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getUptime(t *testing.T) {
	tests := []struct {
		name string
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUptime(); got != tt.want {
				t.Errorf("getUptime() = %v, want %v", got, tt.want)
			}
		})
	}
}
