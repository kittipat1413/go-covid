package services

import (
	"context"
	"go-covid/domain/entity"
	domainerrors "go-covid/domain/errors"
	"go-covid/external/heregeocoder"
)

type GeocoderService interface {
	GetReverseGeocode(ctx context.Context, location entity.Location) (out *entity.Address, err error)
}

func NewGeocoderService() GeocoderService {
	return &geocoderService{}
}

type geocoderService struct{}

func (s *geocoderService) GetReverseGeocode(ctx context.Context, location entity.Location) (out *entity.Address, err error) {
	client := heregeocoder.NewClient(ctx)

	res, err := client.GetReverseGeocode(location.Lat, location.Lng)
	if err != nil {
		err = domainerrors.ErrInternal.Wrap(err)
		return
	}
	out = res.ToAddress()
	return
}
