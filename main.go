package main

import (
	"Assignment_1/handlers"
	"log"
	"net/http"
)

func main() {
	/*
		countriesURL := "http://129.241.150.113:8080/v3.1"
		gutendexURL := "http://129.241.150.113:8000/books/"
		lang2countryURL := "http://129.241.150.113:3000/language2countries/"

		r, err := http.NewRequest(http.MethodGet, countriesURL, nil)
		if err != nil {
			fmt.Errorf("Error in creating request:", err.Error())
		}
	*/

	bookCount := handlers.GetBookCount
	http.HandleFunc("/librarystats/v1/bookcount/", bookCount)

	readership := handlers.GetReaderShip
	http.HandleFunc("/librarystats/v1/readership/", readership)

	status := handlers.GetStatus
	http.HandleFunc("/librarystats/v1/status/", status)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
