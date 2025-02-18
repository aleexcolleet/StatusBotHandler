package usecases

import (
	"MST_FV/internal/repositories"
	"context"
	"fmt"
)

// Services is the dependencies injection. Has the Port implementations to
// be used in the core part of the program
type Services struct {
	urlRepo      repositories.UrlRepo
	checkUrlRepo repositories.CheckUrlRepo
	messageRepo  repositories.Message
}

// ConsultAndSend Domain interface to interact with the adapters
type ConsultAndSend interface {
	consultAndSend(ctx context.Context) error
}

// NewServices is a constructor for the Services
func NewServices(urlRepo repositories.UrlRepo, messageRepo repositories.Message, checkUrlRepo repositories.CheckUrlRepo) *Services {
	return &Services{urlRepo: urlRepo, messageRepo: messageRepo, checkUrlRepo: checkUrlRepo}
}

func (s *Services) ConsultAndSend(ctx context.Context) error {

	urls, err := s.urlRepo.GetUrls(context.Background())
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	err = s.urlRepo.LoadUrls(context.Background(), urls)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	urlsDataResp, err := s.checkUrlRepo.GetCheckResp(context.Background(), urls)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	err = s.urlRepo.LoadStatusResponse(context.Background(), urlsDataResp)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	/*
		urlResp, err := s.urlRepo.GetStatusResponse(context.Background())
		if err != nil {
			return fmt.Errorf("%v", err)
		}
	*/
	msgs, err := s.messageRepo.GetMessages(context.Background(), urlsDataResp)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	err = s.messageRepo.SendMessages(context.Background(), msgs)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}
