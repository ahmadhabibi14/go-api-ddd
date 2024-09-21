package customer

import (
	"errors"
	"go-api-ddd/aggregate"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound = errors.New("the customer not found")
	ErrFailedToAddCustomer = errors.New("failed to add the customer")
	ErrUpdateCustomer = errors.New("failed to update the customer")
)

type CustomerRepository interface {
	Get(id uuid.UUID) (aggregate.Customer, error)
	Add(c aggregate.Customer) error
	Update(c aggregate.Customer) error
}