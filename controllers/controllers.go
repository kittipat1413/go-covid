package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	domainerrors "go-covid/domain/errors"

	"github.com/julienschmidt/httprouter"
)

type ControllerInterface interface {
	Mount(router *httprouter.Router) error
}

func MountAll(router *httprouter.Router, controllers []ControllerInterface) error {
	for _, controller := range controllers {
		if err := controller.Mount(router); err != nil {
			return err
		}
	}
	return nil
}

func ReadJSON(b io.ReadCloser, obj interface{}) error {
	err := json.NewDecoder(b).Decode(obj)
	if err != nil {
		return domainerrors.RequiredParametersMissing{}
	}
	return nil
}

func GetLimitOffsetFromRequest(r *http.Request) (limit *uint64, offset *uint64, err error) {
	var (
		limitArg  uint64
		offsetArg uint64
	)
	limitParam := r.URL.Query().Get("limit")
	if len(limitParam) != 0 {
		limitArg, err = strconv.ParseUint(limitParam, 10, 0)
		if err != nil {
			err = domainerrors.ErrBadRequest.Wrap(err)
			return
		}
		limit = &limitArg
	}

	offsetParam := r.URL.Query().Get("offset")
	if len(offsetParam) != 0 {
		offsetArg, err = strconv.ParseUint(offsetParam, 10, 0)
		if err != nil {
			err = domainerrors.ErrBadRequest.Wrap(err)
			return
		}
		offset = &offsetArg
	}

	return
}
