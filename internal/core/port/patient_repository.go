package port

import "github.com/jairo/patient-api/internal/core/entity"

type PatientRepository interface {
	CreatePatient(*entity.Patient) error
	GetPatientByID(string) (*entity.Patient, error)
	UpdatePatient(*entity.Patient) error
	DeletePatient(string) error
	ListPatients() ([]*entity.Patient, error)
}
