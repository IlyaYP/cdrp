package common

import (
	"fmt"

	"github.com/IlyaYP/cdrp/pkg/axl"
	"github.com/IlyaYP/cdrp/storage/psql"
)

// Config combines sub-configs for all services, storages and providers.
type Config struct {
	CdrPath     string      `mapstructure:"path"`
	PSQLStorage psql.Config `mapstructure:"psql_storage"`
	AXL         axl.Config  `mapstructure:"cucm_axl"`
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

// BuildAXLClient builds AXL dependency.
func (c Config) BuildAXLClient() *axl.AXLClient {

	client := axl.NewClient(c.AXL.CUCM).
		SetAuthentication(c.AXL.User, c.AXL.Pass).
		SetSchemaVersion(c.AXL.Schema).
		SetInsecureSkipVerify(c.AXL.Insec).
		SetRequestResponseDump(c.AXL.Dump)
	return client
}
