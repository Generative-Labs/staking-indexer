package services

import (
	"context"

	"github.com/babylonchain/staking-indexer/config"
)

// Service layer contains the business logic and is used to interact with
// the database and other external clients (if any).
type Services struct {
	cfg *config.Config
}

func New(
	ctx context.Context,
	cfg *config.Config,
) (*Services, error) {
	return &Services{
		cfg: cfg,
	}, nil
}

// DoHealthCheck checks the health of the services by ping the database.
func (s *Services) DoHealthCheck(ctx context.Context) error {
	return nil
}
