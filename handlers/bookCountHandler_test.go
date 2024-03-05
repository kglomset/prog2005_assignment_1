package handlers

import (
	"Assignment_1/util"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// Tests the BookCountHandler method for correct REST method
func TestBookCountHandler(t *testing.T) {
	tests := []struct {
		name   string
		method string
		path   string
		code   int
	}{
		{"Method = GET (Status OK)", http.MethodGet, util.BookCountEndPoint,
			http.StatusOK},
		{"Method = Post (Status not implemented when " +
			"using POST method)", http.MethodPost, util.BookCountEndPoint + "something wrong here",
			http.StatusNotImplemented},
		{"Method = GET (Status bad request" +
			" when not specifying language query)", http.MethodGet, util.BookCountEndPoint + "/",
			http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()
			BookCountHandler(w, r)
		})
	}
}

// Tests countUniqueAuthors for successfully counting unique authors
func TestCountUniqueAuthors(t *testing.T) {
	type args struct {
		books []util.Book
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Empty slice", args{books: []util.Book{}}, 0},
		{"Single book with unique author", args{books: []util.Book{{Authors: []util.Person{{
			"Khabib Nurmagomedov"}}}}}, 1},
		{"Single book with duplicate author",
			args{books: []util.Book{{Authors: []util.Person{{
				"Khabib Nurmagomedov"}, {
				"Daniel Cormier"}, {
				"Daniel Cormier"}}}}},
			2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountUniqueAuthors(tt.args.books); got != tt.want {
				t.Errorf("CountUniqueAuthors() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Tests bookCountRequest for successful language code requests.
func Test_bookCountRequest(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Successful request with valid language code",
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/bookcount?lang=no", nil)
					return req
				}(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bookCountRequest(tt.args.w, tt.args.r)
		})
	}
}

// Tests getAuthorsAndBooks for successful request and correct response
func Test_getAuthorsAndBooks(t *testing.T) {
	type args struct {
		language string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   []util.Book
		wantErr bool
	}{
		{"Success with non-empty data", args{"br"}, 1, []util.Book{{
			23685, "Sarmoniou an Aotrou Quere", []util.Person{{Name: "Quéré, Jean"}}}},
			false},
		{"Error scenario ", args{language: "gjøvik"}, 0, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := getAuthorsAndBooks(tt.args.language)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAuthorsAndBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getAuthorsAndBooks() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getAuthorsAndBooks() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

// Tests getBookStats for successfully displaying correct data
func Test_getBookStats(t *testing.T) {
	type args struct {
		language string
	}
	tests := []struct {
		name    string
		args    args
		want    util.BookCountData
		wantErr bool
	}{
		{"Success with non-empty data", args{"no"},
			util.BookCountData{"no", 21, 16, 0.000288}, false},
		{"Success with non-empty data", args{"fi"},
			util.BookCountData{"fi", 2834, 887, 0.038839}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getBookStats(tt.args.language)
			if (err != nil) != tt.wantErr {
				t.Errorf("getBookStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBookStats() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// Tests getTotalBookCount for the correct amount of books on Gutendex API
func Test_getTotalBookCount(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		{"Total bookcount is correct", 72968, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getTotalBookCount()
			if (err != nil) != tt.wantErr {
				t.Errorf("getTotalBookCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getTotalBookCount() got = %v, want %v", got, tt.want)
			}
		})
	}
}
