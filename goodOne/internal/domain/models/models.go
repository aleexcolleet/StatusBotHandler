package models

type Urls struct {
	Urls    []string  `json:"Urls"`
	UrlData []UrlData `json:"UrlData"`
}
type UrlData struct {
	Url        string `json:"Url"`
	Status     bool   `json:"Status"`
	Comment    string `json:"Comment"`
	StatusCode int    `json:"StatusCode"`
}
