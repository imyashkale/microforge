package dummy

import (
	"context"

	"github.com/imyashkale/microforge/pkg/log"
)

type Service interface {
	Get(ctx context.Context, id string) (Dummy, error)
	Create(ctx context.Context, cr CreateDummy) (Dummy, error)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new url service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the url with the specified the url ID.
func (svc service) Get(ctx context.Context, id string) (Dummy, error) {
	_, err := svc.repo.Get(ctx, id)
	if err != nil {
		return Dummy{}, err
	}
	return Dummy{}, nil
}

// Create
func (svc service) Create(ctx context.Context, cr CreateDummy) (Dummy, error) {
	return Dummy{}, nil
}
