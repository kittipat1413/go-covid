package services

import (
	"context"
	"go-covid/domain/entity"
	domainerrors "go-covid/domain/errors"
)

func (s *patientService) Summary(ctx context.Context, city string) (out []entity.PatientSummary, err error) {
	errLocation := "[services patient/get Summary] "
	defer domainerrors.WrapErr(errLocation, &err)

	out, err = s.patientrepo.GetCitiesSummary(ctx, nil, city)
	if err != nil {
		err = domainerrors.ErrInternal.Wrap(err)
		return
	}
	return
}
