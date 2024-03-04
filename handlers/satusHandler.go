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

// StartUptimeTracking starts a timer to track how long the server have been active
func StartUptimeTracking() {
	startTime = time.Now()
}

// getUptime returns the current timestamp for uptime in seconds
func getUptime() float64 {
	upTime := time.Since(startTime).Seconds()
	return upTime
}

// GetStatusFromEndPoints checks the status of Gutendex API, Language2Countries API and RESTcountries API,
// and creates a status struct with the corresponding servicestatus of each endpoint.
func GetStatusFromEndPoints(w http.ResponseWriter, r *http.Request) {
	gutendexStatus := checkServiceStatus(util.GutendexEndPoint)
	languageAPIStatus := checkServiceStatus(util.L2CEndPoint)
	countriesAPIStatus := checkServiceStatus(util.RestCountriesEndPoint)

	response := util.Status{
		GutendexAPI:  gutendexStatus,
		LanguageAPI:  languageAPIStatus,
		CountriesAPI: countriesAPIStatus,
		Version:      "v1",
		Uptime:       int64(getUptime()),
	}
	// Return the status message in the response
	w.Header().Set("Content-Type", "application/json")

	statusData, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		log.Println("Error encoding response: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	_, err2 := w.Write(statusData)
	if err2 != nil {
		http.Error(w, "Error while writing response", http.StatusInternalServerError)
		return
	}
}

// checkServiceStatus takes a service URL as parameter and checks the status on the endpoint and returns the
// status text.
func checkServiceStatus(serviceURL string) string {
	resp, err := http.Get(serviceURL)
	if err != nil {
		log.Fatal(err)
	}
	statusText := fmt.Sprintf("%d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	return statusText
}
