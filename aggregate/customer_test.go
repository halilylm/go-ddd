package aggregate_test

import (
	"github.com/halilylm/go-ddd/aggregate"
	"testing"
)

func TestNewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}
	testCases := []testCase{
		{
			test:        "empty name validation",
			name:        "",
			expectedErr: aggregate.ErrEmptyCustomerName,
		},
		{
			test:        "valid name",
			name:        "halil",
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)
			if err != tc.expectedErr {
				t.Errorf("expected %v got %v", tc.expectedErr, err)
			}
		})
	}
}
