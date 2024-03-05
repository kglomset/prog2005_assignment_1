package handlers

import (
	"Assignment_1/util"
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

// BookCountHandler handles HTTP requests for book count data.
// It accepts GET requests and delegates the processing to the bookCountRequest function.
func BookCountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		bookCountRequest(w, r)
	} else {
		http.Error(w, "HTTP method not supported", http.StatusNotImplemented)
	}
}

// CountUniqueAuthors counts the number of unique authors present across a slice of books.
// It excludes authors with empty names and those named "unknown" (case-insensitive).
func CountUniqueAuthors(books []util.Book) int {
	uniqueAuthors := make(map[string]bool)
	for _, book := range books {
		for _, author := range book.Authors {
			if author.Name != "" && author.Name != "unknown" {
				uniqueAuthors[author.Name] = true
			}
		}
	}
	return len(uniqueAuthors)
}

// Gets the total number of books from the Gutendex API and returns it as an integer
func getTotalBookCount() (int, error) {
	countResponse, err1 := http.Get(util.GutendexEndPoint)
	if err1 != nil {
		log.Printf("Failed to retrieve the total book count from Gutendex API: %v\n", err1)
		return 0, err1
	}
	defer countResponse.Body.Close()

	var totalBookCount util.BookCountResult
	err2 := json.NewDecoder(countResponse.Body).Decode(&totalBookCount)
	if err2 != nil {
		log.Printf("Failed to decode response from Gutendex API: %v\n", err2)
		return 0, err2
	}
	return totalBookCount.Count, nil
}

// Retrieves the number of authors and books for a specified language from the Gutendex API.
// This function iterates through paginated results until there are no more pages available.
func getAuthorsAndBooks(language string) (int, []util.Book, error) {
	var totalCount int
	var authorArray []util.Book

	for page := 1; ; page++ { // Loop until there are no more pages
		gutendexURL := util.GutendexEndPoint + "?languages=" +
			language + "&page=" + strconv.Itoa(page)

		// Fetch data from the Gutendex API
		gutendexResponse, err1 := http.Get(gutendexURL)
		if err1 != nil {
			log.Printf("Error fetching data from Gutendex API (page %d): %v", page, err1)
			return 0, nil, err1
		}
		defer gutendexResponse.Body.Close()

		// Decode the JSON response from Gutendex
		var result util.BookCountResult
		err2 := json.NewDecoder(gutendexResponse.Body).Decode(&result)
		if err2 != nil {
			log.Printf("Error decoding Gutendex API response (page %d): %v", page, err2)
			return 0, nil, err2
		}

		// Update total count and book list
		totalCount = result.Count
		authorArray = append(authorArray, result.Results...)

		// Check if there are more pages available
		if result.Next == "" {
			break
		}
	}
	return totalCount, authorArray, nil
}

// getBookStats retrieves and calculates book-related statistics for a given language.
func getBookStats(language string) (util.BookCountData, error) {
	totalCount, authorArray, err := getAuthorsAndBooks(language)
	if err != nil {
		return util.BookCountData{}, err
	}

	// Count the number of unique authors from the author array.
	uniqueAuthors := CountUniqueAuthors(authorArray)

	// Get the total number of books across all languages.
	totalBookCount, err := getTotalBookCount()
	if err != nil {
		return util.BookCountData{}, err
	}

	// Calculate the fraction representing the proportion of books in the specified language with six decimals.
	fraction := math.Round(float64(totalCount)/float64(totalBookCount)*1000000) / 1000000

	// Create and return a BookCountData struct containing the calculated statistics.
	return util.BookCountData{
		Language: language,
		Books:    totalCount,
		Authors:  uniqueAuthors,
		Fraction: fraction,
	}, nil
}

// Handles the HTTP request for book count data
func bookCountRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	// Extract language codes from the request query parameters
	isoCodes := r.URL.Query().Get("languages")
	languages := strings.Split(isoCodes, ",")

	// Validate the presence and format of language codes
	if isoCodes == "" || len(languages) == 0 {
		http.Error(w, "You need to specify one or more two-letter language codes to use this query.", http.StatusBadRequest)
		return
	}

	for _, language := range languages {
		if len(language) != 2 {
			http.Error(w, "A language code can only be two letters long.", http.StatusBadRequest)
			return
		}
	}

	// Collect book statistics for each requested language
	var bookStats []util.BookCountData
	for _, language := range languages {
		bookStat, err := getBookStats(language)
		if err != nil {
			http.Error(w, "Failed to retrieve book statistics.", http.StatusInternalServerError)
			return
		}
		bookStats = append(bookStats, bookStat)
	}

	// Marshal the book statistics data to JSON format with indentation
	jsonData, err := json.MarshalIndent(bookStats, "", " ")
	if err != nil {
		http.Error(w, "Failed to generate book statistics. Please try again later.", http.StatusInternalServerError)
		return
	}

	// Write the JSON response to the client
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Error while writing response.", http.StatusInternalServerError)
		return
	}
}
