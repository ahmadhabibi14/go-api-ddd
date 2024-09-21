package memory

import (
	"errors"
	"go-api-ddd/aggregate"
	"go-api-ddd/domain/customer"
	"testing"

	"github.com/google/uuid"
)

func TestMemory_GetCustom(t *testing.T) {
	type testCase struct {
		name string
		id uuid.UUID
		expectedErr error
	}

	cust, err := aggregate.NewCustomer("habi")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()
	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name: "no customer by id",
			id: uuid.MustParse("de63cda0-25d3-4de3-b7e6-6c0f737fa91c"),
			expectedErr: customer.ErrCustomerNotFound,
		}, {
			name: "customer by id",
			id: id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}