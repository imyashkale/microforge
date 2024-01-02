package dummy

import (
	"context"

	"github.com/imyashkale/microforge/pkg/configs"
)

// ########## HADNLES DATABASE INTERACTION

// Repository encapsulates the logic to access url from the data source.
type Repository interface {
	// Get returns the url with the specified url ID.
	Get(ctx context.Context, id string) (Dummy, error)
}

type dummyRepository struct {
	database  configs.DynamoDB
	tableName string
}

// NewRepository
func NewRepository(db configs.DynamoDB, tableName string) Repository {
	return dummyRepository{
		database:  db,
		tableName: tableName,
	}
}

// Get
func (dr dummyRepository) Get(ctx context.Context, id string) (Dummy, error) {
	return Dummy{}, nil
}
