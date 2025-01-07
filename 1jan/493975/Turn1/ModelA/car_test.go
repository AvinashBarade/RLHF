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

func TestCar_NewCar_ShouldCreateValidCarInstance(t *testing.T) {
	c := Car{Make: "Toyota", Model: "Corolla", Year: 2020, Mileage: 15000}
	if c.GetYear() != 2020 {
		t.Errorf("expected year to be 2020, got %d", c.GetYear())
	}
}

func TestCar_Validate_ShouldPassWithCompleteDetails(t *testing.T) {
	c := Car{Make: "Toyota", Model: "Corolla", Year: 2020, Mileage: 15000}
	if err := c.Validate(); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestCar_Validate_ShouldFailWithEmptyMake(t *testing.T) {
	c := Car{Model: "Corolla", Year: 2020, Mileage: 15000}
	if err := c.Validate(); err == nil {
		t.Error("expected error with empty Make")
	}
}

func TestCar_Description_ShouldReturnStringifiedDetails(t *testing.T) {
	c := Car{Make: "Toyota", Model: "Corolla", Year: 2020, Mileage: 15000}
	expected := "Toyota Corolla (2020) - 15000 miles"
	actual := c.Description()
	if actual != expected {
		t.Errorf("expected description to be '%s', got '%s'", expected, actual)
	}
}
