package services

import (
	"context"
	"go-covid/domain/entity"
	patientrepo "go-covid/domain/repository/patient"
	geocoderservice "go-covid/services/geocoder"
)

type PatientService interface {
	Register(ctx context.Context, in RegisterInput) (err error)
	Summary(ctx context.Context, city string) (out []entity.PatientSummary, err error)
}

type patientService struct {
	geocoderservice geocoderservice.GeocoderService
	patientrepo     patientrepo.PatientRepository
}

func NewPatientService(geocoderservice geocoderservice.GeocoderService, patientrepo patientrepo.PatientRepository) PatientService {
	return &patientService{
		geocoderservice: geocoderservice,
		patientrepo:     patientrepo,
	}
}
