package handlers

import (
	"Assignment_1/util"
	"testing"
)

// Tests the service status for the three endpoints used in the project
func Test_checkServiceStatus(t *testing.T) {
	type args struct {
		serviceURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Gutendex status OK", args{serviceURL: util.GutendexEndPoint}, "200: OK"},
		{"Language2Countries status OK", args{serviceURL: util.L2CEndPoint + "/no"}, "200: OK"},
		{"REST Countries status OK", args{serviceURL: util.RestCountriesEndPoint + "no"}, "200: OK"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkServiceStatus(tt.args.serviceURL); got != tt.want {
				t.Errorf("checkServiceStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
