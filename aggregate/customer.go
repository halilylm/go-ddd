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

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

func (c *Customer) GetName() string {
	return c.person.Name
}

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
