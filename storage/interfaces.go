package storage

import (
	"context"
	"io"

	"github.com/IlyaYP/cdrp/model"
)

type RecordsStorage interface {
	io.Closer

	// CreateRecord creates a new model.Record.
	// Returns ErrAlreadyExists if user exists.
	CreateRecord(ctx context.Context, record model.Record) error
}
