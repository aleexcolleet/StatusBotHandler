package models

//In models we create the structs that I will be using with domain and repo adaptations.

type URLs struct {
	urls     []string
	urlsData []URLData
}
type URLData struct {
	url        string
	status     bool
	comment    string
	statusCode int
}
