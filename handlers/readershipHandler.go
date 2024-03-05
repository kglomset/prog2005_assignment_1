package handlers

import (
	"Assignment_1/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ReadershipHandler handles HTTP requests for readership data.
func ReadershipHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET.
	if r.Method == http.MethodGet {
		ReadershipRequest(w, r)
	} else {
		http.Error(w, "HTTP method not supported", http.StatusNotImplemented)
	}
}

// ReadershipRequest handles requests for readership data for a specified language.
func ReadershipRequest(w http.ResponseWriter, r *http.Request) {

	// Extract language code from request path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}

	language := parts[4]

	// Parse and validate limit parameter from the query string
	limit := 0 // Default limit if not provided
	limitString := r.URL.Query().Get("limit")
	if limitString != "" {
		var err error
		limit, err = strconv.Atoi(limitString)
		if err != nil || limit <= 0 {
			http.Error(w, "Invalid limit parameter: must be a positive integer", http.StatusBadRequest)
			return
		}
	}

	// Retrieve readership data for the specified language and limit
	readershipData, err0 := getReadershipData(language, limit)
	if err0 != nil {
		http.Error(w, "Failed to retrieve readership data.", http.StatusInternalServerError)
	}

	//Convert readership data to JSON format with indentation
	data, err := json.MarshalIndent(readershipData, "", " ")
	if err != nil {
		http.Error(w, "Failed to generate book statistics. Please try again later.", http.StatusInternalServerError)
	}

	// Write JSON data to the response body
	_, err2 := w.Write(data)
	if err2 != nil {
		http.Error(w, "Error while writing response.", http.StatusInternalServerError)
	}
}

func retrievePopulationData(country string) (int, error) {
	//var populationData util.RestCountriesResponse
	url := fmt.Sprintf(util.RestCountriesEndPoint+"%s", country)
	restCountriesResp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to fetch data from RestCountries API for country %s: %v\n", country, err)
	}
	defer restCountriesResp.Body.Close()

	err = json.NewDecoder(restCountriesResp.Body).Decode(&util.RestCountriesResponse)
	if err != nil {
		log.Printf("Error decoding response from RestCountries API: %v\n", err)
	}

	return util.RestCountriesResponse[0].Population, nil
}

// retrieveLanguageData retrieves a list of countries associated with a specific language.
// It fetches data from the Gutendex API endpoint and returns the list as a slice of util.Countries structs.
func retrieveLanguageData(language string) ([]util.Countries, error) {

	// Make a GET request to the specified API
	lang2countResp, err1 := http.Get(util.L2CEndPoint + language)
	if err1 != nil {
		log.Printf("Error retrieving language data from Language2Countries API (language: %s): %v",
			language, err1)
		return nil, err1
	}
	defer lang2countResp.Body.Close()

	// Decode the JSON response into a slice of util.Countries structs.
	var countries []util.Countries
	err2 := json.NewDecoder(lang2countResp.Body).Decode(&countries)
	if err2 != nil {
		log.Printf("Error decoding response from Gutendex API (language: %s): %v", language, err2)
		return nil, err2
	}
	return countries, nil
}

// getReadershipData retrieves readership-related data for a given language.
// It allows limiting the results to a specified number of countries.
func getReadershipData(language string, limit int) ([]util.ReadershipData, error) {

	// Retrieve a list of countries for the specified language.
	countries, err1 := retrieveLanguageData(language)
	if err1 != nil {
		return nil, err1
	}

	// Initialize a slice to store the readership data
	var readershipData []util.ReadershipData

	// Iterate through the countries, potentially limited by the provided limit.
	for i, country := range countries {
		if limit > 0 && i >= limit {
			break
		}

		// Retrieve population data for the current country.
		population, err2 := retrievePopulationData(country.IsoCode)
		if err2 != nil {
			log.Println("Could not retrieve population data", err2)
		}

		// Get the total count of books and authors for the language within this country.
		totalCount, authorArray, _ := getAuthorsAndBooks(language)
		uniqueAuthors := CountUniqueAuthors(authorArray)

		// Create a ReadershipData struct and append it to the slice.
		readershipData = append(readershipData, util.ReadershipData{
			country.OfficialName,
			country.IsoCode,
			totalCount,
			uniqueAuthors,
			population,
		})
	}
	return readershipData, nil
}
