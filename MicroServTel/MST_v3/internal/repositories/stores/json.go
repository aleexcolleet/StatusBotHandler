package stores

import (
	"cmd/main.go/internal/repositories"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

/*
This is an implementation of the port. It takes a JSON file as a source for URLs and
then it stores the Resp on another JSON file (works as a repository)
*/

// JsonURL is a go struct that contains info about each URL converted from the Jsonfile.
// We need to map (relate dates that are in different formats) variables. Golang knows now that
// url is associated to URL and will take the same value converted.
type JsonURL struct {
	Url string
	Id  int
}

type JsonStores struct {
	URLs []repositories.URL
}

// NewJsonStores is a constructor that returns a JsonStores struct
func NewJsonStores() *JsonStores {
	return &JsonStores{}
}

/*
LoadURLs is a function that loads URLs from Json file and stores them in a struct JsonStores
 1. Open the file JsonURLs.json with os.Open("JsonURLs.json") and respective defer file.Close()
 2. Read the file written on bytes with the func os.Readfile(file)
 3. Convert bytes into a Go Struct
*/
func (S *JsonStores) LoadURLs(ctx context.Context) error {
	filePath, _ := filepath.Abs("JsonURLs.json")
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening JsonURLs.json: %v", err)
	}
	defer file.Close()

	bytesRode, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading JsonURLs.json: %v", err)
	}

	// In order to add the id num, I need to create a tmp variables (not struct)
	//to take values, and then the actual struct will add the URL value and
	//id according to iterator + 1
	var data struct {
		URLs []string `json:"urls"`
	}
	err = json.Unmarshal(bytesRode, &data)
	if err != nil {
		return fmt.Errorf("error unmarshalling JsonURLs.json: %v", err)
	}
	//Make creates a slice of type []JsonURL with data.URLs length
	S.URLs = make([]repositories.URL, len(data.URLs))
	for i, url := range data.URLs {
		S.URLs[i] = repositories.URL{
			Url: url,
			Id:  i + 1,
		}
	}
	return nil
}

func (S *JsonStores) GetURLs(ctx context.Context) (repositories.URLs, error) {
	return repositories.URLs{
		URLs: S.URLs,
	}, nil

}
