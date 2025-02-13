package domain

import (
	"MST_FV2/internal/repositories"
	"context"
	"fmt"
)

type Services struct {
	UrlRepo     repositories.UrlRepo
	MessageRepo repositories.MessageRepo
}

func NewServices(urlRepo repositories.UrlRepo, messageRepo repositories.MessageRepo) *Services {
	return &Services{UrlRepo: urlRepo, MessageRepo: messageRepo}
}

// Domain interface
type ServiceInterface interface {
	AnalyzeAndSend()
}

func (s *Services) AnalyzeAndSend() {

	//URLRepo call
	urls, err := s.UrlRepo.GetUrls(context.Background())
	if err != nil {
		fmt.Printf("Error getting urls: %s\n", err.Error())
	}

	// Checker
	urlData, err := checker.checkerUrls(urls)

	//MessageRepo
	UrlMessage, err := s.GetMessage(context.Background(), urlData)
	if err != nil {
		fmt.Printf("Error getting message: %s\n", err.Error())
	}
	err = s.MessageRepo.SendMessage(context.Background(), UrlMessage)

}
