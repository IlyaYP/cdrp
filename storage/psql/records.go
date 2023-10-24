package psql

import (
	"context"
	"errors"
	"fmt"

	"github.com/IlyaYP/cdrp/model"
	"github.com/IlyaYP/cdrp/pkg"
	"github.com/IlyaYP/cdrp/storage/psql/schema"
	"github.com/uptrace/bun/driver/pgdriver"
)

// CreateRecord implements the storage.Record interface.
func (st Storage) CreateRecord(ctx context.Context, record model.Record) (model.Record, error) {
	dbObj := schema.NewRecordFromCanonical(record)

	_, err := st.db.NewInsert().
		Model(&dbObj).
		Returning("*").
		Exec(ctx)
	if err != nil {
		pgErr := &pgdriver.Error{}
		if errors.As(err, pgErr) {
			if pgErr.IntegrityViolation() {
				return model.Record{}, pkg.ErrAlreadyExists
			}
		}
		return model.Record{}, err
	}

	obj, err := dbObj.ToCanonical()
	if err != nil {
		return model.Record{}, fmt.Errorf("conveting to canonical model: %w", err)
	}

	return obj, nil
}
