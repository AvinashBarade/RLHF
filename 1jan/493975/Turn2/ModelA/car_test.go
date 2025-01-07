package car

import (
	"fmt"
	"testing"
)

type Car struct {
	Make    string
	Model   string
	Year    int
	Mileage int
}

func (c Car) GetYear() int {
	return c.Year
}

func (c Car) Validate() error {
	if c.Make == "" || c.Model == "" || c.Year < 1900 || c.Mileage < 0 {
		return fmt.Errorf("invalid car details")
	}
	return nil
}

func (c Car) Description() string {
	return fmt.Sprintf("%s %s (%d) - %d miles", c.Make, c.Model, c.Year, c.Mileage)
}

func (c Car) Upgrade() error {
	if err := c.Validate(); err != nil {
		return err
	}
	c.Year += 1
	c.Mileage += 1000
	return nil
}

func TestTable_Car_Upgrade(t *testing.T) {
	cases := []struct {
		name        string
		car         Car
		expectedErr error
		expectedCar Car
	}{
		{
			name:        "valid upgrade",
			car:         Car{Make: "Toyota", Model: "Corolla", Year: 2020, Mileage: 15000},
			expectedErr: nil,
			expectedCar: Car{Make: "Toyota", Model: "Corolla", Year: 2021, Mileage: 16000},
		},
		{
			name:        "invalid mileage",
			car:         Car{Make: "Toyota", Model: "Corolla", Year: 2020, Mileage: -100},
			expectedErr: fmt.Errorf("invalid car details"),
			expectedCar: Car{Make: "Toyota", Model: "Corolla", Year: 2020, Mileage: -100},
		},
		{
			name:        "empty make",
			car:         Car{Model: "Corolla", Year: 2020, Mileage: 15000},
			expectedErr: fmt.Errorf("invalid car details"),
			expectedCar: Car{Model: "Corolla", Year: 2020, Mileage: 15000},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.car.Upgrade()
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
				return
			}
			if tc.car != tc.expectedCar {
				t.Errorf("expected car %v, got %v", tc.expectedCar, tc.car)
			}
		})
	}
}
