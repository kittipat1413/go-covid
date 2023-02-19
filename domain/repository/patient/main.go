package repository

import (
	"context"
	"fmt"
	"go-covid/domain/entity"
	"strconv"
	"time"

	"go-covid/data"
)

const tableName = "patient"

type PatientRepository interface {
	Create(ctx context.Context, scope data.Scope, in CreatePatientInput) (out entity.Patient, err error)
	GetCitiesSummary(ctx context.Context, scope data.Scope, city string) (out []entity.PatientSummary, err error)
}

type patientRepository struct{}

func NewPatientRepository() PatientRepository {
	return &patientRepository{}
}

type patientsModel []patientModel
type patientModel struct {
	ID        int64     `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Gender    string    `db:"gender"`
	Age       uint8     `db:"age"`
	Lat       string    `db:"latitude"`
	Lng       string    `db:"longitude"`
	City      string    `db:"city"`
	District  string    `db:"district"`
	Street    string    `db:"street"`
	CreatedAt time.Time `db:"created_at"`
}

func (p *patientModel) FromEntity(e entity.Patient) (m *patientModel) {
	m = &patientModel{
		FirstName: e.FirstName,
		LastName:  e.LastName,
		Gender:    e.Gender.String(),
		Age:       e.Age,
		Lat:       fmt.Sprintf("%f", e.Lat),
		Lng:       fmt.Sprintf("%f", e.Lng),
		City:      e.City,
		District:  e.District,
		Street:    e.Street,
	}
	return
}

func (p *patientModel) ToEntity() (e entity.Patient, err error) {
	lat, err := strconv.ParseFloat(p.Lat, 64)
	if err == nil {
		return e, err
	}
	lng, err := strconv.ParseFloat(p.Lng, 64)
	if err == nil {
		return e, err
	}

	e = entity.Patient{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Gender:    new(entity.Gender).Parse(p.Gender),
		Age:       p.Age,
		City:      p.City,
		District:  p.District,
		Street:    p.Street,
		Location: entity.Location{
			Lat: lat,
			Lng: lng,
		},
	}
	return
}

func (p *patientsModel) ToEntities() (e entity.Patients, err error) {
	patients := make(entity.Patients, len(*p))
	for idx, patientModel := range *p {
		patients[idx], err = patientModel.ToEntity()
		if err == nil {
			return entity.Patients{}, err
		}
	}

	return patients, nil
}
