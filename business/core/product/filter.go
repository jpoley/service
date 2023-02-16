package product

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/ardanlabs/service/business/sys/validate"
	"github.com/google/uuid"
)

// Used to check for sql injection problems.
var sqlInjection = regexp.MustCompile("^[A-Za-z0-9_]+$")

// =============================================================================

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	ID       *uuid.UUID `validate:"omitempty,uuid4"`
	Name     *string    `validate:"omitempty,min=3"`
	Cost     *int       `validate:"omitempty,numeric"`
	Quantity *int       `validate:"omitempty,numeric"`
}

// Validate checks the data in the model is considered clean.
func (qf *QueryFilter) Validate() error {
	if err := validate.Check(qf); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

// ByID sets the ID field of the QueryFilter value.
func (qf *QueryFilter) ByID(id uuid.UUID) {
	var zero uuid.UUID
	if id != zero {
		qf.ID = &id
	}
}

// ByName sets the Name field of the QueryFilter value.
func (qf *QueryFilter) ByName(name string) error {
	if name != "" {
		if !sqlInjection.MatchString(name) {
			return errors.New("invalid name format")
		}

		qf.Name = &name
	}

	return nil
}

// ByCost sets the Cost field of the QueryFilter value.
func (qf *QueryFilter) ByCost(cost int) {
	qf.Cost = &cost
}

// ByQuantity sets the Quantity field of the QueryFilter value.
func (qf *QueryFilter) ByQuantity(quantity int) {
	qf.Quantity = &quantity
}
