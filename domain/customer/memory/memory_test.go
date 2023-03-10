package memory

import (
	"github.com/google/uuid"
	"github.com/halilylm/go-ddd/aggregate"
	"github.com/halilylm/go-ddd/domain/customer"
	"testing"
)

func TestCustomerRepository_Get(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}
	c, err := aggregate.NewCustomer("halil")
	if err != nil {
		t.Fatal(err)
	}
	id := c.GetID()
	repo := CustomerRepository{
		customers: map[uuid.UUID]*aggregate.Customer{
			id: &c,
		},
	}
	testCases := []testCase{
		{
			name:        "No Customer By ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: customer.ErrCustomerNotFound,
		}, {
			name:        "Customer By ID",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := repo.Get(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemory_AddCustomer(t *testing.T) {
	type testCase struct {
		name        string
		cust        string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add Customer",
			cust:        "Percy",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := CustomerRepository{
				customers: map[uuid.UUID]*aggregate.Customer{},
			}

			cust, err := aggregate.NewCustomer(tc.cust)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(&cust)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(cust.GetID())
			if err != nil {
				t.Fatal(err)
			}
			if found.GetID() != cust.GetID() {
				t.Errorf("Expected %v, got %v", cust.GetID(), found.GetID())
			}
		})
	}
}
