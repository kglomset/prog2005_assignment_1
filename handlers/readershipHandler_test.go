package handlers

import (
	"Assignment_1/util"
	"net/http"
	"reflect"
	"testing"
)

func TestReadershipHandler(t *testing.T) {
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
			ReadershipHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestReadershipRequest(t *testing.T) {
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
			ReadershipRequest(tt.args.w, tt.args.r)
		})
	}
}

func Test_getReadershipData(t *testing.T) {
	type args struct {
		language string
		limit    int
	}
	tests := []struct {
		name    string
		args    args
		want    []util.ReadershipData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getReadershipData(tt.args.language, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("getReadershipData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getReadershipData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_retrieveLanguageData(t *testing.T) {
	type args struct {
		language string
	}
	tests := []struct {
		name    string
		args    args
		want    []util.Countries
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := retrieveLanguageData(tt.args.language)
			if (err != nil) != tt.wantErr {
				t.Errorf("retrieveLanguageData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("retrieveLanguageData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_retrievePopulationData(t *testing.T) {
	type args struct {
		country string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := retrievePopulationData(tt.args.country)
			if (err != nil) != tt.wantErr {
				t.Errorf("retrievePopulationData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("retrievePopulationData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
