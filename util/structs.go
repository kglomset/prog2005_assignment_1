package util

type BookCountData struct {
	Language string  `json:"language"`
	Books    int     `json:"books"`
	Authors  int     `json:"authors"`
	Fraction float64 `json:"fractions"`
}

type ReadershipData struct {
	Country    string `json:"country"`
	Isocode    string `json:"isocode"`
	Books      int    `json:"books"`
	Authors    int    `json:"authors"`
	Readership int    `json:"readership"`
}

type Status struct {
	GutendexAPI  string `json:"gutendexAPI"`
	LanguageAPI  string `json:"languageAPI"`
	CountriesAPI string `json:"countriesAPI"`
	Version      string `json:"version"`
	Uptime       int64  `json:"uptime"`
}
type Person struct {
	Name string `json:"name"`
}

type BookCountResponse struct {
	Languages []BookCountData `json:"languages"`
}

type Book struct {
	Id      int      `json:"id"`
	Title   string   `json:"title"`
	Authors []Person `json:"authors"`
}

type BookCountResult struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Book `json:"results"`
}

var RestCountriesResponse []struct {
	//Country    string `json:"common"`
	Population int `json:"population"`
}

type Countries struct {
	IsoCode      string `json:"ISO3166_1_Alpha_2"`
	OfficialName string `json:"Official_Name"`
	Language     string `json:"Language"`
}
