package structs

type BookCount struct {
	Language  string  `json:"language"`
	Books     string  `json:"books"`
	Authors   string  `json:"authors"`
	Fractions float32 `json:"fractions"`
}
type Readership struct {
	Country    string `json:"country"`
	Isocode    string `json:"isocode"`
	Books      int    `json:"books"`
	Authors    int    `json:"authors"`
	Readership int    `json:"readership"`
}

type ServiceStatus struct {
	Name  string `json:"name"`
	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`
}

type Status struct {
	GutendexAPI  string `json:"gutendexAPI"`
	LanguageAPI  string `json:"languageAPI"`
	CountriesAPI string `json:"countriesAPI"`
	Version      string `json:"version"`
	Uptime       int64  `json:"uptime"`
}
