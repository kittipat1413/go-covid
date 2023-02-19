package repository

import (
	"context"
	"go-covid/data"
	"go-covid/domain/entity"
	domainerrors "go-covid/domain/errors"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
)

func (p *patientRepository) GetCitiesSummary(ctx context.Context, scope data.Scope, city string) (out []entity.PatientSummary, err error) {
	errLocation := "[domain repository/patient GetCitiesSummary] "
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

	var sqlBuilder squirrel.SelectBuilder
	if len(city) != 0 {
		sqlBuilder = data.CreateSqlBuilder().Select(
			"city",
			"district",
			"gender",
			"AVG(age) AS avg_age",
			"count(id) AS count",
		).
			From(tableName).
			GroupBy("city", "district", "gender").
			Where(
				squirrel.Eq{"city": city},
			)
	} else {
		sqlBuilder = data.CreateSqlBuilder().Select(
			"city",
			"gender",
			"AVG(age) AS avg_age",
			"count(id) AS count",
		).
			From(tableName).
			GroupBy("city", "gender")
	}

	getSql, getArgs, err := sqlBuilder.ToSql()
	if err != nil {
		err = domainerrors.ErrInternal.Wrap(err)
		return
	}

	dbOut := []struct {
		City         string  `db:"city"`
		District     string  `db:"district"`
		Gender       string  `db:"gender"`
		AvgAge       float64 `db:"avg_age"`
		PatientCount uint64  `db:"count"`
	}{}

	if err = scope.Select(&dbOut, getSql, getArgs...); err != nil {
		err = domainerrors.DatabaseError{}.Wrap(err)
		return
	}

	err = copier.Copy(&out, dbOut)
	if err != nil {
		return []entity.PatientSummary{}, err
	}

	return
}
