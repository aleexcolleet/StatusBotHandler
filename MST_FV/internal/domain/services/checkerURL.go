package services

import (
	"MicroServ2/internal/repositories"
	"context"
)

/*
	checkerURL.go is a service that I use to request httpUrls from a client
	and returns it to store the response and other data into the repo.
	HOW?

	- Create the constructor for an instance of the port (repo interface)
	- Fetch URLs struct from storage in repo.
	- Check URls with CheckURL one by one so that I can treat and store errors.
		This func returns a tmp struct repositories.URLData
	-
*/

type CheckerURL struct {
	repo repositories.CheckerRepo
}

func NewCheckerURL(ctx context.Context, repo repositories.URLRepo) *CheckerURL {
	return &CheckerURL{
		repo: repo,
	}
}

// CheckAndStoreURLStatus todo-> I'm not sure if this should store the repos. or only check 'em
func (C *CheckerURL) CheckAndStoreURLStatus(ctx context.Context) error {
	urlResp, err := C.repo.CheckURLStatus(context.Background())
	//	C.urlRepo.LoadResponse(context.Background(), urlData)
	return nil
}

/*
	*DNS: Translation between domain name and its IP.
		Could mean an error if translation fails
*/
