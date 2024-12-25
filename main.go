package main

import (
	"fmt"
	"sync"
)

// Define a Patient struct to hold patient information
type Patient struct {
	ID        int
	Name      string
	Age       int
	Diagnosis string
}

// Define a PatientData struct to hold a slice of patients
type PatientData struct {
	Patients []Patient
}

// ProcessPatient is a function closure that performs some analysis on a single patient
func ProcessPatient(patient Patient, wg *sync.WaitGroup, minAge, maxAge int, diagnoses []string) {
	defer wg.Done()

	// Check if the patient matches the filtering criteria
	if patient.Age >= minAge && patient.Age <= maxAge {
		for _, diag := range diagnoses {
			if patient.Diagnosis == diag {
				// Example analysis: Check if the patient is over 65
				if patient.Age > 65 {
					fmt.Println("Patient", patient.Name, "is over 65 and has been flagged.")
				}
				return
			}
		}
	}
}

// AnalyzePatients processes a large dataset of patients in parallel with filtering
func AnalyzePatients(patientData PatientData, minAge, maxAge int, diagnoses []string) {
	var wg sync.WaitGroup

	// Iterate over the patients using a range loop
	for _, patient := range patientData.Patients {
		wg.Add(1)
		go ProcessPatient(patient, &wg, minAge, maxAge, diagnoses)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

func main() {
	// Example dataset of patients
	patients := []Patient{
		{ID: 1, Name: "Alice Johnson", Age: 68, Diagnosis: "Hypertension"},
		{ID: 2, Name: "Bob Smith", Age: 35, Diagnosis: "Asthma"},
		{ID: 3, Name: "Charlie Brown", Age: 72, Diagnosis: "Diabetes"},
		{ID: 4, Name: "Diana Prince", Age: 29, Diagnosis: "None"},
		{ID: 5, Name: "Edward Cullen", Age: 109, Diagnosis: "Diabetes"},
		{ID: 6, Name: "Franklin D. Roosevelt", Age: 95, Diagnosis: "Hypertension"},
	}

	// Create a PatientData struct to hold the dataset
	patientData := PatientData{Patients: patients}

	// Example 1: Analyze patients aged between 60 and 80 with "Hypertension" or "Diabetes" diagnosis
	diagnoses := []string{"Hypertension", "Diabetes"}
	AnalyzePatients(patientData, 60, 80, diagnoses)

	// Example 2: Analyze patients with age > 30 and "None" diagnosis
	AnalyzePatients(patientData, 30, 100, []string{"None"})

	// Example 3: Analyze all patients (no filtering)
	AnalyzePatients(patientData, 0, 100, []string{})
}
