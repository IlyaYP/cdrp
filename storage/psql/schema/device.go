package schema

import (
	"github.com/IlyaYP/cdrp/model"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Device struct {
	bun.BaseModel `bun:"table:devices,alias:d"`
	Pkid          uuid.UUID `bun:"type:uuid,pk"`
	Name          string    `bun:"type:varchar(129)"`
	Description   string    `bun:"type:varchar(512)"`
}

// NewDeviceFromCanonical creates a new Device DB object from canonical model.
func NewDeviceFromCanonical(obj model.Device) Device {
	return Device{
		Pkid:        obj.Pkid,
		Name:        obj.Name,
		Description: obj.Description,
	}
}

// ToCanonical converts a DB object to canonical model.
func (d Device) ToCanonical() (model.Device, error) {
	return model.Device{
		Pkid:        d.Pkid,
		Name:        d.Name,
		Description: d.Description,
	}, nil
}
