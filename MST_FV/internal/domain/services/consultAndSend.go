package services

import (
	"MST_FV/internal/repositories"
	"context"
	"fmt"
)

// Services is the dependencies injection. Has the Port implementations to
// be used in the core part of the program
type Services struct {
	urlRepo      repositories.UrlRepo
	messageRepo  repositories.Message
	checkUrlRepo repositories.CheckUrlRepo
}

// ConsultAndSend Domain interface to interact with the adapters
type ConsultAndSend interface {
	consultAndSend(ctx context.Context)
}

// NewServices is a constructor for the Services
func NewServices(urlRepo repositories.UrlRepo, messageRepo repositories.Message, checkUrlRepo repositories.CheckUrlRepo) *Services {
	return &Services{urlRepo: urlRepo, messageRepo: messageRepo, checkUrlRepo: checkUrlRepo}
}

func (s *Services) consultAndSend(ctx context.Context) {

	//Fetch Urls from wherever they're
	urls, err := s.urlRepo.GetUrls(context.Background())
	if err != nil {
		fmt.Errorf("%v", err)
	}

	//Call checker and fetch return struct (response Data)
	urlsDataResp, err := s.checkUrlRepo.GetCheckResp(context.Background(), urls)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	//Load the checkerResp into the repository
	err = s.urlRepo.LoadStatusResponse(context.Background(), urlsDataResp)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	//GetStatusResponse which I won't need because I already have it, but I'll implement it
	//urlsDataResp, err := s.urlRepo.GetStatusResponse(context.Background())
	
	//Message interface call
}
