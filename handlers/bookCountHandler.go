package handlers

import (
	"Assignment_1/structs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

/*
const GutendexAPIURL = "http://129.241.150.113:8000/books/"

func GetBooksByLanguage(w http.ResponseWriter, r *http.Request) {
	// Extract language parameter from the request
	language := r.URL.Query().Get("language")

	// Make the request to the Gutendex API
	resp, err := http.Get(GutendexAPIURL + "?language=" + language)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error making Gutendex API request: %s", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check if the request was successful (status code 200)
	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Gutendex API request failed with status code: %d", resp.StatusCode), http.StatusInternalServerError)
		return
	}

	// Decode the response body into a slice of Book structs
	var books []structs.BookCount
	err = json.NewDecoder(resp.Body).Decode(&books)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding Gutendex API response: %s", err), http.StatusInternalServerError)
		return
	}

	// Return the books in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

*/

type GutendexResponse struct {
	Count   int `json:"count"`
	Results []struct {
		Authors []struct {
			Name      string `json:"name"`
			BirthYear int    `json:"birth_year"`
			DeathYear int    `json:"death_year"`
		} `json:"authors"`
	} `json:"results"`
}

type BookCountResponse struct {
	Language string  `json:"language"`
	Books    int     `json:"books"`
	Authors  int     `json:"authors"`
	Fraction float64 `json:"fraction"`
}

func GetBookCount(w http.ResponseWriter, r *http.Request) {
	languages := r.URL.Query().Get("language")
	languageCodes := strings.Split(languages, ",")

	var response []structs.BookCount

	for _, lang := range languageCodes {
		// Make requests to Gutendex API for each language
		gutendexURL := fmt.Sprintf("http://129.241.150.113:8000/books/?languages=%s", lang)
		resp, err := http.Get(gutendexURL)
		if err != nil {
			log.Println("Error fetching data from Gutendex:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		var gutendexData GutendexResponse
		err = json.NewDecoder(resp.Body).Decode(&gutendexData)
		if err != nil {
			log.Println("Error decoding Gutendex response:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Extract relevant data
		booksCount := gutendexData.Count
		authorsCount := len(gutendexData.Results[0].Authors)
		totalBooksFromGutendex := 100000 // Replace with actual total books count (if available)

		resp1 := BookCountResponse{
			Language: lang,
			Books:    booksCount,
			Authors:  authorsCount,
			Fraction: float64(booksCount) / float64(totalBooksFromGutendex),
		}

		response = append(response, resp1)
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
