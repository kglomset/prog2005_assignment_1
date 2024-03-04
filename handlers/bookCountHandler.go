package handlers

import (
	"Assignment_1/util"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// BookCountHandler handles HTTP requests for book count data.
// It accepts GET requests and delegates the processing to the bookCountRequest function.
// Any other HTTP method results in a "405 Method Not Allowed" error with a suggestion to use GET.
func BookCountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		bookCountRequest(w, r)
	} else {
		http.Error(w, "REST method '"+r.Method+"' is not supported. Please try again with method "+
			" '"+http.MethodGet+"' instead. ", http.StatusMethodNotAllowed)
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
	countResponse, error_1 := http.Get(util.GutendexEndPoint)
	if error_1 != nil {
		log.Printf("Failed to retrieve the total book count from Gutendex API: %v\n", error_1)
		return 0, error_1
	}
	defer countResponse.Body.Close()

	var totalBookCount util.BookCountResult
	error_2 := json.NewDecoder(countResponse.Body).Decode(&totalBookCount)
	if error_2 != nil {
		log.Printf("Failed to decode response from Gutendex API: %v\n", error_2)
		return 0, error_2
	}
	return totalBookCount.Count, nil
}

// Retrieves the number of authors and books for a specified language from the Gutendex API
func getAuthorsAndBooks(language string) (int, []util.Book, error) {
	var totalCount int
	var authorArray []util.Book

	for i := 1; ; i++ {
		gutendexResponse, err1 := http.Get(util.GutendexEndPoint + "?languages=" +
			language + "&page=" + strconv.Itoa(i))
		if err1 != nil {
			log.Println(" EDIT THIS STANDARD_ERROR", err1.Error())
		}

		var result util.BookCountResult
		err2 := json.NewDecoder(gutendexResponse.Body).Decode(&result)
		if err2 != nil {
			log.Println("EDIT THIS DECODING_ERROR"+"of the Gutendex API's respnse.", err2.Error())
		}

		totalCount = result.Count
		authorArray = append(authorArray, result.Results...)

		if result.Next == "" {
			break
		}
	}
	return totalCount, authorArray, nil
}

// Takes the data that is to compose a struct of type BookCountData, creates the struct and returns it
func getBookStats(language string) (util.BookCountData, error) {
	totalCount, authorArray, _ := getAuthorsAndBooks(language)
	uniqueAuthors := CountUniqueAuthors(authorArray)

	totalBookCount, _ := getTotalBookCount()
	totalBooks := totalBookCount
	fraction := float64(totalBooks) / float64(totalCount)

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
	isoCodes := r.URL.Query().Get("languages")
	languages := strings.Split(isoCodes, ",")

	if isoCodes == "" || len(languages) == 0 {
		http.Error(w, "You need to specify one or more two-letter language codes to use this query.", http.StatusBadRequest)
		return
	}

	for _, language := range languages {
		if len(language) != 2 {
			http.Error(w, "Every language code must be 2 letters long.", http.StatusBadRequest)
			return
		}
	}

	var bookStats []util.BookCountData
	for _, language := range languages {
		bookStat, err5 := getBookStats(language)
		if err5 != nil {
			http.Error(w, "Failed to retrieve book statistics.", http.StatusInternalServerError)
			return
		}
		bookStats = append(bookStats, bookStat)
	}

	jsonData, err6 := json.Marshal(bookStats)
	if err6 != nil {
		http.Error(w, "Do something about this error"+"of Book Count endpoint's statistics.", http.StatusInternalServerError)
		return
	}

	_, err7 := w.Write(jsonData)
	if err7 != nil {
		http.Error(w, "Error while writing response.", http.StatusInternalServerError)
		return
	}
}
