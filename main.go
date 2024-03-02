package main

import (
	"Assignment_1/handlers"
	"Assignment_1/util"
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

	http.HandleFunc(util.BookCountEndPoint, handlers.BookCountHandler)
	http.HandleFunc(util.ReadershipEndPoint, handlers.readershipRequest)
	http.HandleFunc(util.StatusEndPoint, handlers.GetStatus)
	log.Fatal(http.ListenAndServe(util.DefaultPort, nil))
}
