package aggregate_test

import (
	"errors"
	"go-api-ddd/aggregate"
	"testing"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test string
		name string 
		expectedErr error
	}

	testCases := []testCase{
		{
			test: "Empty name validation",
			name: "",
			expectedErr: aggregate.ErrInvalidPerson,
		}, {
			test: "Valid name",
			name: "Ahmad Habibi",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases{
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}