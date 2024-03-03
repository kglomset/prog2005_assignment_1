package handlers

import "net/http"

func ReadershipHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ReadershipRequest(w, r)
	default:
		http.Error(w, "Place some fitting error message here", http.StatusNotImplemented)
		return
	}
}

func ReadershipRequest(w http.ResponseWriter, r *http.Request) {
	// Innmaten til denne inneholder blant annet getReadershipData
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
