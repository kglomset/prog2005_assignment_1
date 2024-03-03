package handlers

import (
	"Assignment_1/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	// Innmaten til denne inneholder blant annet getReadershipData
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}

	language := parts[4]
	log.Println(language)
	population, err0 := retrievePopulationData(language)
	if err0 != nil {
		log.Println("Whats is wrong: ", err0)
	}
	data, err := json.Marshal(population)
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

func retrieveLanguageData() {

}

func getReadershipData() {

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
