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

type DevicesStorage interface {
	io.Closer

	// CreateDevice creates a new model.Device.
	// Returns ErrAlreadyExists if device exists.
	CreateDevice(ctx context.Context, device model.Device) (model.Device, error)

	// CreateDevices bulk create a new [] model.Device.
	// Returns ErrAlreadyExists if devices exists.
	CreateDevices(ctx context.Context, devices []model.Device) ([]model.Device, error)
}
