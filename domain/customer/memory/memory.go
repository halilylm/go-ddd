package memory

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/halilylm/go-ddd/aggregate"
	"github.com/halilylm/go-ddd/domain/customer"
	"sync"
)

type MemoryRepository struct {
	customers map[uuid.UUID]*aggregate.Customer
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]*aggregate.Customer),
	}
}

// Get finds a customer by ID
func (mr *MemoryRepository) Get(uid uuid.UUID) (*aggregate.Customer, error) {
	v, ok := mr.customers[uid]
	if !ok {
		return nil, customer.ErrCustomerNotFound
	}
	return v, nil
}

// Add will add a new customer to the repository
func (mr *MemoryRepository) Add(c *aggregate.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]*aggregate.Customer)
		mr.Unlock()
	}
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

// Update will replace an existing customer information with the new customer information
func (mr *MemoryRepository) Update(c *aggregate.Customer) error {
	if _, ok := mr.customers[c.GetID()]; !ok {
		return customer.ErrCustomerNotFound
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}
