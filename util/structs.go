package util

// BookCountData represents the data that are displayed to the user of the BookCount endpoint
type BookCountData struct {
	Language string  `json:"language"`
	Books    int     `json:"books"`
	Authors  int     `json:"authors"`
	Fraction float64 `json:"fractions"`
}

// ReadershipData represents the data that are displayed to the user of the Readership endpoint
type ReadershipData struct {
	Country    string `json:"country"`
	Isocode    string `json:"isocode"`
	Books      int    `json:"books"`
	Authors    int    `json:"authors"`
	Readership int    `json:"readership"`
}

// Status represents the data that are displayed to the user of the Status endpoint
type Status struct {
	GutendexAPI  string `json:"gutendexAPI"`
	LanguageAPI  string `json:"languageAPI"`
	CountriesAPI string `json:"countriesAPI"`
	Version      string `json:"version"`
	Uptime       int64  `json:"uptime"`
}

// Person represents the name of an author in the gutendex response
type Person struct {
	Name string `json:"name"`
}

// BookCountResponse represents the collection of languages from BookCountData
type BookCountResponse struct {
	Languages []BookCountData `json:"languages"`
}

// Book represents a book from the Gutendex API
type Book struct {
	Id      int      `json:"id"`
	Title   string   `json:"title"`
	Authors []Person `json:"authors"`
}

// BookCountResult represents the result from a request to Gutendex API
type BookCountResult struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Book `json:"results"`
}

// Represents a response from the Restcountries API
var RestCountriesResponse []struct {
	//Country    string `json:"common"`
	Population int `json:"population"`
}

// Countries represents the information about a country
type Countries struct {
	IsoCode      string `json:"ISO3166_1_Alpha_2"`
	OfficialName string `json:"Official_Name"`
	Language     string `json:"Language"`
}
