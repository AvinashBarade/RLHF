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

func TestCar_Upgrade(t *testing.T) {
	// Setup a valid car instance to be used across subtests
	validCar := Car{Make: "Toyota", Model: "Corolla", Year: 2020, Mileage: 15000}

	t.Run("Valid Upgrade", func(t *testing.T) {
		if err := validCar.Upgrade(); err != nil {
			t.Fatalf("expected upgrade to succeed, got error %v", err)
		}
		expectedCar := Car{Make: "Toyota", Model: "Corolla", Year: 2021, Mileage: 16000}
		if validCar != expectedCar {
			t.Fatalf("expected car to be %v, got %v", expectedCar, validCar)
		}
	})

	t.Run("Upgrade Should Fail With Invalid Mileage", func(t *testing.T) {
		errCar := Car{Make: "Honda", Model: "Civic", Year: 2019, Mileage: -100}
		if err := errCar.Upgrade(); err == nil {
			t.Fatal("expected upgrade to fail")
		}
		expectedErr := fmt.Errorf("invalid car details")
		if err.Error() != expectedErr.Error() {
			t.Fatalf("expected error message '%v', got '%v'", expectedErr, err)
		}
	})

	t.Run("Upgrade Should Fail With Missing Make", func(t *testing.T) {
		errCar := Car{Model: "BMW", Year: 2021, Mileage: 5000}
		if err := errCar.Upgrade(); err == nil {
			t.Fatal("expected upgrade to fail")
		}
		expectedErr := fmt.Errorf("invalid car details")
		if err.Error() != expectedErr.Error() {
			t.Fatalf("expected error message '%v', got '%v'", expectedErr, err)
		}
	})
}
