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

func ReadershipHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ReadershipRequest(w, r)
	} else {
		http.Error(w, "Place some fitting error message here", http.StatusNotImplemented)
	}
}

func ReadershipRequest(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}

	language := parts[4]

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
	readershipData, err0 := getReadershipData(language, limit)
	if err0 != nil {
		log.Println("Whats is wrong: ", err0)
	}
	data, err := json.MarshalIndent(readershipData, "", " ")
	if err != nil {
		log.Println("test")
	}
	_, err2 := w.Write(data)
	if err2 != nil {
		log.Println("test2")
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

func retrieveLanguageData(language string) ([]util.Countries, error) {
	lang2countResp, err1 := http.Get(util.L2CEndPoint + language)
	if err1 != nil {
		log.Println("Error getting response", err1)
	}
	defer lang2countResp.Body.Close()

	var countries []util.Countries
	err2 := json.NewDecoder(lang2countResp.Body).Decode(&countries)
	if err2 != nil {
		log.Println("Problem decoding response men legg til mer her", err2)
	}
	return countries, nil
}

func getReadershipData(language string, limit int) ([]util.ReadershipData, error) {
	countries, err1 := retrieveLanguageData(language)
	if err1 != nil {
		return nil, err1
	}

	var readershipData []util.ReadershipData
	for i, country := range countries {
		if limit > 0 && i >= limit {
			break
		}
		population, err2 := retrievePopulationData(country.IsoCode)
		if err2 != nil {
			log.Println("Could not retrieve population data", err2)
		}

		totalCount, authorArray, _ := getAuthorsAndBooks(language)
		uniqueAuthors := CountUniqueAuthors(authorArray)

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

/*
  {
     "country": "Norway", Hentes fra Language2Countries
     "isocode": "NO",
     "books": 21, Hentes fra gutendex - Har allerede metode for denne og
     "authors": 14, Hentes fra gutendex
     "readership": 5379475 - Hentes fra restcountries
  },
Her er det viktig å få med at det skal hentes land som det snakkes f.eks norsk i hvis no er imput. Det finnes flere
norsktalende land, og alle disse listes opp i l2c responsen. Dersom det er flere land skal hvert land inn i hver sin entry,
dette gjelder både country og readership

*/
