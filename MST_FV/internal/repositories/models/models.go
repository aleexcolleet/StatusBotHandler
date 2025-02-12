package models

// Models struct definition
type URLs struct {
	URLs []string
}

type URLData struct {
	URL        string
	Status     bool
	Comment    string
	StatusCode int
}
