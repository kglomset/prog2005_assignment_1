package handlers

import (
	"Assignment_1/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var startTime time.Time

func StartUptimeTracking() {
	startTime = time.Now()
}

func getUptime() float64 {
	upTime := time.Since(startTime).Seconds()
	return upTime
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	gutendexStatus := checkServiceStatus(util.GutendexEndPoint)
	languageAPIStatus := checkServiceStatus(util.L2CEndPoint)
	countriesAPIStatus := checkServiceStatus(util.RestCountriesEndPoint)

	//statusMessage := "Service is running"

	response := util.Status{
		GutendexAPI:  gutendexStatus,
		LanguageAPI:  languageAPIStatus,
		CountriesAPI: countriesAPIStatus,
		Version:      "v1",
		Uptime:       int64(getUptime()),
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
