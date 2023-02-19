package repository

import (
	"context"
	"go-covid/data"
	"go-covid/domain/entity"
	domainerrors "go-covid/domain/errors"
)

type CreatePatientInput struct {
	Patient entity.Patient
}

func (p *patientRepository) Create(ctx context.Context, scope data.Scope, in CreatePatientInput) (out entity.Patient, err error) {
	errLocation := "[domain repository/patient Create] "
	defer domainerrors.WrapErr(errLocation, &err)

	if scope == nil {
		scope, err = data.NewScope(ctx, nil)
		if err != nil {
			err = domainerrors.ErrDatabase.Wrap(err)
			return
		} else {
			defer scope.End(&err)
		}
	}

	inputModel := new(patientModel).FromEntity(in.Patient)
	insertSql, insertArgs, err := data.CreateSqlBuilder().
		Insert(tableName).
		Columns(
			"first_name",
			"last_name",
			"gender",
			"age",
			"latitude",
			"longitude",
			"city",
			"district",
			"street",
		).
		Values(
			inputModel.FirstName,
			inputModel.LastName,
			inputModel.Gender,
			inputModel.Age,
			inputModel.Lat,
			inputModel.Lng,
			inputModel.City,
			inputModel.District,
			inputModel.Street,
		).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		err = domainerrors.ErrInternal.Wrap(err)
		return
	}

	dbOut := patientModel{}
	err = scope.Get(&dbOut, insertSql, insertArgs...)
	if err != nil {
		err = new(domainerrors.DatabaseError).Wrap(err)
		return
	}

	out, err = dbOut.ToEntity()
	return
}
