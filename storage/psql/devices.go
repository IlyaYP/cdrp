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

// CreateDevice implements the storage.Device interface.
func (st Storage) CreateDevice(ctx context.Context, device model.Device) (model.Device, error) {
	dbObj := schema.NewDeviceFromCanonical(device)

	_, err := st.db.NewInsert().
		Model(&dbObj).
		Returning("*").
		Exec(ctx)
	if err != nil {
		pgErr := &pgdriver.Error{}
		if errors.As(err, pgErr) {
			if pgErr.IntegrityViolation() {
				return model.Device{}, pkg.ErrAlreadyExists
			}
		}
		return model.Device{}, err
	}

	obj, err := dbObj.ToCanonical()
	if err != nil {
		return model.Device{}, fmt.Errorf("conveting to canonical model: %w", err)
	}

	return obj, nil
}

// CreateDevices implements the storage.Device interface.
func (st Storage) CreateDevices(ctx context.Context, devices []model.Device) ([]model.Device, error) {

	// _, err := db.NewInsert().
	// 	Model(&book).
	// 	On("CONFLICT (id) DO UPDATE").
	// 	Set("title = EXCLUDED.title").
	// 	Exec(ctx)
	dbObjs := []schema.Device{}
	for _, device := range devices {
		dbObj := schema.NewDeviceFromCanonical(device)
		dbObjs = append(dbObjs, dbObj)
	}

	_, err := st.db.NewInsert().
		Model(&dbObjs).
		On("CONFLICT (pkid) DO UPDATE").
		Set("name = EXCLUDED.name, description = EXCLUDED.description").
		Returning("*").
		Exec(ctx)
	if err != nil {
		pgErr := &pgdriver.Error{}
		if errors.As(err, pgErr) {
			if pgErr.IntegrityViolation() {
				return nil, pkg.ErrAlreadyExists
			}
		}
		return nil, err
	}

	objs := []model.Device{}
	for _, dbObj := range dbObjs {
		obj, err := dbObj.ToCanonical()
		if err != nil {
			return nil, fmt.Errorf("conveting to canonical model: %w", err)
		}
		objs = append(objs, obj)
	}

	return objs, nil
}

// CreateTableDevices creates table devices
func (st Storage) CreateTableDevices(ctx context.Context) error {
	_, err := st.db.NewCreateTable().Model((*schema.Device)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// DropTableDevices drops table devices
func (st Storage) DropTableDevices(ctx context.Context) error {
	_, err := st.db.NewDropTable().Model((*schema.Device)(nil)).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// ResetTableDevices drops table devices
func (st Storage) ResetTableDevices(ctx context.Context) error {
	err := st.db.ResetModel(ctx, (*schema.Device)(nil))
	if err != nil {
		return err
	}

	return nil
}
