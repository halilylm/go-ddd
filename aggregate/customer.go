package aggregate

import (
	"errors"
	"github.com/google/uuid"
	"github.com/halilylm/go-ddd/entity"
	"github.com/halilylm/go-ddd/valueobject"
)

var (
	ErrEmptyCustomerName = errors.New("a customer should have a valid name")
)

// Customer is an aggregate including
// all entities to represent a Customer
type Customer struct {
	// person is the root entity,
	// person.ID is the main identifier
	person *entity.Person
	// products that this customer owns
	products []*entity.Item
	// transactions performed by the customer
	transactions []valueobject.Transaction
}

// GetID retrieves customer's id
func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

// SetID sets customer's id
func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

// SetName sets customer's name
func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

// GetName retrieves customer's name
func (c *Customer) GetName() string {
	return c.person.Name
}

// NewCustomer is a factory function to
// create customer aggregates
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrEmptyCustomerName
	}
	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}
	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}
