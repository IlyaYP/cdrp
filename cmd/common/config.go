package common

import (
	"fmt"

	"github.com/IlyaYP/cdrp/storage/psql"
)

// Config combines sub-configs for all services, storages and providers.
type Config struct {
	CdrPath     string
	PSQLStorage psql.Config `mapstructure:"psql_storage"`
}

// BuildPsqlStorage builds psql.Storage dependency.
func (c Config) BuildPsqlStorage() (*psql.Storage, error) {
	st, err := psql.New(
		psql.WithConfig(c.PSQLStorage),
	)
	if err != nil {
		return nil, fmt.Errorf("building psql storage: %w", err)
	}

	return st, nil
}
