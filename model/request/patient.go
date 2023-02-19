package request

import (
	domainerrors "go-covid/domain/errors"

	"github.com/go-playground/validator"
)

type RegisterPatient struct {
	FirstName string  `json:"first_name" validate:"required"`
	LastName  string  `json:"last_name" validate:"required"`
	Gender    string  `json:"gender" validate:"required,oneof=Male Female"`
	Age       uint8   `json:"age" validate:"required,gt=0"`
	Lat       float64 `json:"lat" validate:"required,latitude"`
	Lng       float64 `json:"long" validate:"required,longitude"`
}

func (t *RegisterPatient) Validate() (err error) {
	errLocation := "[models/requests RegisterPatient Validate] "
	defer domainerrors.WrapErr(errLocation, &err)

	vInstance := validator.New()
	if err = vInstance.Struct(t); err != nil {
		ve, ok := err.(validator.ValidationErrors)
		if !ok {
			err = domainerrors.ErrInternal.Wrap(err)
			return
		}
		err = domainerrors.ValidationError{
			ValidatorErrors: ve,
		}.Wrap(err)
		return
	}
	return nil
}
