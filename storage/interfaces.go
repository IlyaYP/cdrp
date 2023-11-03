package storage

import (
	"context"
	"io"

	"github.com/IlyaYP/cdrp/model"
)

type RecordsStorage interface {
	io.Closer

	// CreateRecord creates a new model.Record.
	// Returns ErrAlreadyExists if record exists.
	CreateRecord(ctx context.Context, record model.Record) (model.Record, error)
}
