package handlers

import (
	"Assignment_1/structs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var startTime = time.Now()

func GetStatus(w http.ResponseWriter, r *http.Request) {
	// Your implementation for getting status information goes here
	// ...

	// For the sake of this example, let's assume a status message

	gutendexStatus := checkServiceStatus("http://129.241.150.113:8000/books/")
	languageAPIStatus := checkServiceStatus("http://129.241.150.113:3000/language2countries/")
	countriesAPIStatus := checkServiceStatus("https://restcountries.com/v3.1/all")

	//statusMessage := "Service is running"
	uptimeSeconds := time.Since(startTime).Seconds()

	response := structs.Status{
		GutendexAPI:  gutendexStatus,
		LanguageAPI:  languageAPIStatus,
		CountriesAPI: countriesAPIStatus,
		Version:      "v1",
		Uptime:       int64(uptimeSeconds),
	}
	// Return the status message in the response
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	// Handle error in encoding
	if err != nil {
		log.Println("Error encoding response: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func checkServiceStatus(serviceURL string) string {
	resp, err := http.Get(serviceURL)
	if err != nil {
		log.Fatal(err)
	}
	statusText := fmt.Sprintf("%d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	return statusText
}
