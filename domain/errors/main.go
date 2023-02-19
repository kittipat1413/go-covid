package errors

import (
	"fmt"
	"go-covid/internal/constants"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/stoewer/go-strcase"
)

var (
	ErrInternal   = InternalError{ErrCode: constants.StatusCodeGenericInternalError, Message: "internal server error"}
	ErrNotFound   = NotFoundError{ErrCode: constants.StatusCodeGenericNotFoundError, Message: "not found"}
	ErrBadRequest = BadRequestError{ErrCode: constants.StatusCodeGenericBadRequestError, Message: "bad request"}
	ErrDatabase   = BadRequestError{ErrCode: constants.StatusCodeDatabaseError, Message: "database error(s)"}
)

type Interface interface {
	Code() string
	GetMessage() string
	GetHttpCode() int
}

type InternalError struct {
	ErrCode string `json:"code"`
	Message string `json:"message"`
}

func (e InternalError) Code() string {
	return e.ErrCode
}
func (e InternalError) Error() string {
	return e.Message
}
func (e InternalError) GetMessage() string {
	return e.Message
}
func (e InternalError) GetHttpCode() int {
	return http.StatusInternalServerError
}
func (e InternalError) Wrap(err error) error {
	return fmt.Errorf("%w: %v", e, err)
}

type DatabaseError struct {
	ErrCode string `json:"code"`
	Message string `json:"message"`
}

func (e DatabaseError) Code() string {
	return e.ErrCode
}
func (e DatabaseError) Error() string {
	return e.Message
}
func (e DatabaseError) GetMessage() string {
	return e.Message
}
func (e DatabaseError) GetHttpCode() int {
	return http.StatusInternalServerError
}
func (e DatabaseError) Wrap(err error) error {
	return fmt.Errorf("%w: %v", e, err)
}

type NotFoundError struct {
	ErrCode string `json:"code"`
	Message string `json:"message"`
}

func (e NotFoundError) Code() string {
	return e.ErrCode
}
func (e NotFoundError) Error() string {
	return e.Message
}
func (e NotFoundError) GetMessage() string {
	return e.Message
}
func (e NotFoundError) GetHttpCode() int {
	return http.StatusNotFound
}
func (e NotFoundError) Wrap(err error) error {
	return fmt.Errorf("%w: %v", e, err)
}

type UnprocessableEntityError struct {
	Message string `json:"message"`
}

func (e UnprocessableEntityError) Code() string {
	return constants.StatusCodeUnprocessableEntity
}
func (e UnprocessableEntityError) Error() string {
	return e.GetMessage()
}
func (e UnprocessableEntityError) GetMessage() string {
	return e.Message
}
func (e UnprocessableEntityError) GetHttpCode() int {
	return http.StatusUnprocessableEntity
}
func (e UnprocessableEntityError) Wrap(err error) error {
	return fmt.Errorf("%w: %v", e, err)
}

type BadRequestError struct {
	ErrCode string `json:"code"`
	Message string `json:"message"`
}

func (e BadRequestError) Code() string {
	return e.ErrCode
}
func (e BadRequestError) Error() string {
	return e.Message
}
func (e BadRequestError) GetMessage() string {
	return e.Message
}
func (e BadRequestError) GetHttpCode() int {
	return http.StatusBadRequest
}
func (e BadRequestError) Wrap(err error) error {
	return fmt.Errorf("%w: %v", e, err)
}

type RequiredParametersMissing struct {
	BadRequestError
}

func (e RequiredParametersMissing) Code() string {
	return constants.StatusCodeMissingRequiredParameters
}
func (e RequiredParametersMissing) Error() string {
	return "required parameter(s) are missing"
}
func (e RequiredParametersMissing) GetMessage() string {
	return "required parameter(s) are missing"
}

type ValidationError struct {
	ValidatorErrors validator.ValidationErrors
}

func (e ValidationError) Error() string {
	return e.GetMessage()
}
func (e ValidationError) Code() string {
	return constants.StatusCodeInvalidParameters
}
func (e ValidationError) GetMessage() string {
	var errFields []string
	for _, fe := range e.ValidatorErrors {
		errFields = append(errFields, strcase.LowerCamelCase(fe.Field()))
	}
	return fmt.Sprintf("%s has invalid format", strings.Join(errFields, ", "))
}
func (e ValidationError) GetHttpCode() int {
	return http.StatusBadRequest
}
func (e ValidationError) Wrap(err error) error {
	return fmt.Errorf("%w: %v", e, err)
}

func WrapErr(name string, errptr *error) {
	if *errptr != nil {
		*errptr = fmt.Errorf(name+": %w", *errptr)
	}
}
