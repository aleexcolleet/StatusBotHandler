package models

//In models we create the structs that I will be using with domain and repo adaptations.

type URLs struct {
	urls     []string  `json:"urls"`
	urlsData []URLData `json:"urls_data"`
}
type URLData struct {
	url        string `json:"url"`
	status     bool   `json:"status"`
	comment    string `json:"comment"`
	statusCode int    `json:"status_code"`
}
