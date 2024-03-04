package handlers

import (
	"Assignment_1/util"
	"net/http"
	"reflect"
	"testing"
)

func TestBookCountHandler(t *testing.T) {
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
			BookCountHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestCountUniqueAuthors(t *testing.T) {
	type args struct {
		books []util.Book
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountUniqueAuthors(tt.args.books); got != tt.want {
				t.Errorf("CountUniqueAuthors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookCountRequest(t *testing.T) {
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
			bookCountRequest(tt.args.w, tt.args.r)
		})
	}
}

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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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

func Test_getTotalBookCount(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
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
