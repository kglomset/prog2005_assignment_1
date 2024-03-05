package handlers

import (
	"Assignment_1/util"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// Tests the ReadershipHandler method for correct REST method
func TestReadershipHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Successful request with valid language and limit",
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/readership?lang=en&limit=10", nil)
					return req
				}(),
			},
		},
		{
			name: "Request with invalid HTTP method",
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					req, _ := http.NewRequest(http.MethodPost, "/readership", nil)
					return req
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadershipHandler(tt.args.w, tt.args.r)
		})
	}
}

// Tests ReadershipRequest for successful requests.
func TestReadershipRequest(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Successful request with valid language and limit",
			args: args{
				w: httptest.NewRecorder(), // Mock response writer
				r: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/readership?lang=en&limit=10", nil)
					return req
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadershipRequest(tt.args.w, tt.args.r)
		})
	}
}

// Tests the getReadershipData method for successfully retrieving data corresponding to language code
// and the specified limit
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
		{"Successful response 'no' limit 1", args{"no", 1}, []util.ReadershipData{{
			"Iceland", "IS", 21, 16, 366425}}, false},
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

// Tests retrieveLanguageData for the correct data from corresponding data structures given the correct language input
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
		{"Successful data from Norway}", args{"no"}, []util.Countries{
			{"IS", "Iceland", "no"},
			{"NO", "Norway", "no"},
			{"SJ", "Svalbard and Jan Mayen Islands", "no"}},
			false},
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

// Tests retrievePopulationdata for the correct population from RESTCountries API
func Test_retrievePopulationData(t *testing.T) {
	tests := []struct {
		name    string
		country string
		want    int
		wantErr bool
	}{
		{"Population is found with iso code no ", "no", 5379475, false},
		{"Population is found with iso code se ", "se", 10353442, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := retrievePopulationData(tt.country)
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
