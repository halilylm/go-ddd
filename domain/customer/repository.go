package customer

import (
	"errors"
	"github.com/google/uuid"
	"github.com/halilylm/go-ddd/aggregate"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found")
	ErrFailedToAddCustomer = errors.New("error when adding a customer")
	ErrUpdateCustomer      = errors.New("error when updating the customer")
)

type CustomerRepository interface {
	Get(uuid uuid.UUID) (*aggregate.Customer, error)
	Add(customer *aggregate.Customer) error
	Update(customer *aggregate.Customer) error
}
