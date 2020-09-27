package model

import "github.com/artkescha/CRUD_for_car/model/status"

type Car struct {
	ID      string        `json:"-"`
	Model   *string       `json:"model" valid:"required, runelength(1|300)" minLength:"1"`
	Vendor  *string       `json:"vendor" valid:"required, runelength(1|300)" minLength:"1"`
	Price   uint64        `json:"price"`
	Status  status.Status `json:"status"`
	Mileage uint64        `json:"mileage"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Car) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Car) SetID(id string) error {
	c.ID = id
	return nil
}
