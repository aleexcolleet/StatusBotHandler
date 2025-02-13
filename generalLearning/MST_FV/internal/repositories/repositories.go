package repositories

import (
	"MicroServ2/internal/repositories/models"
	"context"
)

// URLRepo is an interface for the repository itself. Doesn't have dependencies
type URLRepo interface {
	LoadURL(ctx context.Context, urls models.URLs) error
	GetURL(ctx context.Context) (models.URLs, error)
	LoadResponse(ctx context.Context, urlsResponse []models.URLData) error
}

//Checker Port Definition.
// Is an interface that needs a previous repository to function. It's a dependency.

type CheckerRepo interface {
	//TODO
	CheckURLStatus(ctx context.Context) error
}

//Message Port definition
