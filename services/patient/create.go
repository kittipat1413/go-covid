package services

import (
	"context"
	"go-covid/domain/entity"
	domainerrors "go-covid/domain/errors"
	patientrepo "go-covid/domain/repository/patient"
)

type RegisterInput struct {
	FirstName string
	LastName  string
	Gender    string
	Age       uint8
	Lat       float64
	Lng       float64
}

func (s *patientService) Register(ctx context.Context, in RegisterInput) (err error) {
	errLocation := "[services patient/create Register] "
	defer domainerrors.WrapErr(errLocation, &err)

	patientLocation := entity.Location{Lat: in.Lat, Lng: in.Lng}
	var patientAddr *entity.Address
	patientAddr, err = s.geocoderservice.GetReverseGeocode(ctx, patientLocation)
	if err != nil {
		return domainerrors.ErrInternal.Wrap(err)
	}

	if !patientAddr.IsCountryName(entity.Thailand) {
		return domainerrors.UnprocessableEntityError{Message: "you are not allowed to access from your location"}
	}

	patient := entity.Patient{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Gender:    new(entity.Gender).Parse(in.Gender),
		Age:       in.Age,
		City:      patientAddr.City,
		District:  patientAddr.District,
		Street:    patientAddr.Street,
		Location:  patientLocation,
	}

	_, err = s.patientrepo.Create(ctx, nil, patientrepo.CreatePatientInput{Patient: patient})
	if err != nil {
		return domainerrors.ErrInternal.Wrap(err)
	}

	return
}
