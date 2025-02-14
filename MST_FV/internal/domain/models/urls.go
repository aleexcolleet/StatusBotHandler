package models

//In models we create the structs that I will be using with domain and repo adaptations.

type URLs struct {
	Urls     []string  `json:"urls"`
	UrlsData []URLData `json:"urls_data"`
}
type URLData struct {
	Url        string `json:"url"`
	Status     bool   `json:"status"`
	Comment    string `json:"comment"`
	StatusCode int    `json:"status_code"`
}
